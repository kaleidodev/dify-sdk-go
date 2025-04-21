package types

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type EventCh struct {
	ch  chan []byte
	ctx context.Context
}

func NewEventCh(ch chan []byte, ctx context.Context) *EventCh {
	if ctx == nil {
		ctx = context.Background()
	}

	return &EventCh{
		ch:  ch,
		ctx: ctx,
	}
}

func (c *EventCh) ParseToStructCh() <-chan ChunkChatCompletionResponse {
	streamChannel := make(chan ChunkChatCompletionResponse, 10)

	go func() {
		defer func() {
			close(streamChannel)
		}()

		for {
			select {
			case <-c.ctx.Done():

			default:
				data, ok := <-c.ch
				if !ok {
					return
				}

				var resp ChunkChatCompletionResponse
				err := json.Unmarshal(data, &resp)
				if err != nil {
					log.Printf("Error unmarshalling chunk completion response: %v", err)
					resp.Event = "error"
					resp.Status = 500
					resp.Code = "json unmarshal error"
					resp.Message = err.Error()
				}
				streamChannel <- resp
			}
		}
	}()

	return streamChannel
}

type Event struct {
	Type string
	Data interface{}
}

type eventType struct {
	Event string `json:"event"` // 事件类型
}

func (c *EventCh) ParseToEventCh() <-chan Event {
	eventChannel := make(chan Event, 10)

	go func() {
		defer func() {
			close(eventChannel)
		}()

		for {
			select {
			case <-c.ctx.Done():
				return

			default:
				data, ok := <-c.ch
				if !ok {
					return
				}

				event, err := c.processEvent(data)
				if err != nil {
					log.Printf("Error processing event: %v", err)
					continue
				}

				eventChannel <- event
			}
		}
	}()

	return eventChannel
}

var eventTypeMap = map[string]interface{}{
	EVENT_ERROR:                    &EventError{},
	EVENT_MESSAGE:                  &EventMessage{},
	EVENT_MESSAGE_END:              &EventMessageEnd{},
	EVENT_TTS_MESSAGE:              &EventTtsMessage{},
	EVENT_TTS_MESSAGE_END:          &EventTtsMessageEnd{},
	EVENT_MESSAGE_FILE:             &EventMessageFile{},
	EVENT_MESSAGE_REPLACE:          &EventMessageReplace{},
	EVENT_AGENT_THOUGHT:            &EventAgentThought{},
	EVENT_AGENT_MESSAGE:            &EventAgentMessage{},
	EVENT_WORKFLOW_STARTED:         &EventWorkflowStarted{},
	EVENT_WORKFLOW_FINISHED:        &EventWorkflowFinished{},
	EVENT_NODE_STARTED:             &EventNodeStarted{},
	EVENT_NODE_FINISHED:            &EventNodeFinished{},
	EVENT_NODE_RETRY:               &EventNodeRetry{},
	EVENT_PARALLEL_BRANCH_STARTED:  &EventParallelBranchStarted{},
	EVENT_PARALLEL_BRANCH_FINISHED: &EventParallelBranchFinished{},
	EVENT_ITERATION_STARTED:        &EventIterationStarted{},
	EVENT_ITERATION_NEXT:           &EventIterationNext{},
	EVENT_ITERATION_COMPLETED:      &EventIterationCompleted{},
	EVENT_LOOP_STARTED:             &EventLoopStarted{},
	EVENT_LOOP_NEXT:                &EventLoopNext{},
	EVENT_LOOP_COMPLETED:           &EventLoopCompleted{},
	EVENT_TEXT_CHUNK:               &EventTextChunk{},
	EVENT_TEXT_REPLACE:             &EventTextReplace{},
	EVENT_AGENT_LOG:                &EventAgentLog{},
}

func (c *EventCh) processEvent(data []byte) (Event, error) {
	var event eventType
	if err := json.Unmarshal(data, &event); err != nil {
		return Event{}, fmt.Errorf("failed to unmarshal event type: %w", err)
	}

	var result Event
	result.Type = event.Event

	targetType, exists := eventTypeMap[event.Event]
	if !exists {
		return Event{}, fmt.Errorf("unknown event type: %s", event.Event)
	}

	if err := json.Unmarshal(data, &targetType); err != nil {
		return Event{}, fmt.Errorf("failed to unmarshal event data: %w", err)
	}

	result.Data = targetType
	return result, nil
}
