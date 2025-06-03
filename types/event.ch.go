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

// ParseToStructCh 将事件解析为一个大的统一结构体(字段有冗余)
func (c *EventCh) ParseToStructCh() <-chan ChunkChatCompletionResponse {
	streamChannel := make(chan ChunkChatCompletionResponse, 500)

	go func() {
		defer func() {
			close(streamChannel)
		}()

		for {
			select {
			case <-c.ctx.Done():

			case data, ok := <-c.ch:
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

// SimplePrint 仅输出最终文字结果
func (c *EventCh) SimplePrint() (ch <-chan string, conversationId, taskId *string) {
	streamChannel := make(chan string, 500)
	id := ""
	taskid := ""
	taskId = &taskid
	conversationId = &id

	go func() {
		defer func() {
			close(streamChannel)
		}()

		for {
			select {
			case <-c.ctx.Done():
				return

			case data, ok := <-c.ch:
				if !ok {
					return
				}

				event, err := c.processEvent(data)
				if err != nil {
					log.Printf("Error processing event: %v data: %s", err, string(data))
					continue
				}

				var str string
				switch event.Type {
				case EVENT_ERROR:
					eventData := event.Data.(*EventError)
					str = fmt.Sprintf("糟糕,请求出错了! Status: %d Code: %s Message: %s", eventData.Status, eventData.Code, eventData.Message)
				case EVENT_MESSAGE:
					eventData := event.Data.(*EventMessage)
					str = eventData.Answer

					if id == "" && eventData.ConversationId != "" {
						id = eventData.ConversationId
					}
					if taskid == "" && eventData.TaskId != "" {
						taskid = eventData.TaskId
					}
				case EVENT_MESSAGE_END:
				case EVENT_TTS_MESSAGE:
					eventData := event.Data.(*EventTtsMessage)
					str = eventData.Audio

					if id == "" && eventData.ConversationId != "" {
						id = eventData.ConversationId
					}
					if taskid == "" && eventData.TaskId != "" {
						taskid = eventData.TaskId
					}
				case EVENT_TTS_MESSAGE_END:
				case EVENT_MESSAGE_FILE:
				case EVENT_MESSAGE_REPLACE:
					eventData := event.Data.(*EventMessageReplace)
					str = fmt.Sprintf("\n%s\n", eventData.Answer)
				case EVENT_AGENT_THOUGHT:
					eventData := event.Data.(*EventAgentThought)
					if eventData.Observation != "" {
						str = fmt.Sprintf("  \n> **调用工具: %s** \n```json\n// 请求：\n%s\n\n// 响应:\n%s\n```  \n", eventData.Tool, eventData.ToolInput, eventData.Observation)
					}

					if id == "" && eventData.ConversationId != "" {
						id = eventData.ConversationId
					}
					if taskid == "" && eventData.TaskId != "" {
						taskid = eventData.TaskId
					}
				case EVENT_AGENT_MESSAGE:
					eventData := event.Data.(*EventAgentMessage)
					str = eventData.Answer

					if id == "" && eventData.ConversationId != "" {
						id = eventData.ConversationId
					}
					if taskid == "" && eventData.TaskId != "" {
						taskid = eventData.TaskId
					}
				case EVENT_WORKFLOW_STARTED:
				case EVENT_WORKFLOW_FINISHED:
					eventData := event.Data.(*EventWorkflowFinished)

					// 这部分内容实际上是前面内容的重复，先不输出
					//for _, v := range eventData.Data.Outputs {
					//	str = fmt.Sprintf("%s%s,", str, v)
					//}
					//str = strings.TrimSuffix(str, ",")

					if id == "" && eventData.WorkflowRunId != "" {
						id = eventData.WorkflowRunId
					}
					if taskid == "" && eventData.TaskId != "" {
						taskid = eventData.TaskId
					}
				case EVENT_NODE_STARTED:
				case EVENT_NODE_FINISHED:
				case EVENT_NODE_RETRY:
				case EVENT_PARALLEL_BRANCH_STARTED:
				case EVENT_PARALLEL_BRANCH_FINISHED:
				case EVENT_ITERATION_STARTED:
				case EVENT_ITERATION_NEXT:
				case EVENT_ITERATION_COMPLETED:
				case EVENT_LOOP_STARTED:
				case EVENT_LOOP_NEXT:
				case EVENT_LOOP_COMPLETED:
				case EVENT_TEXT_CHUNK:
				case EVENT_TEXT_REPLACE:
				case EVENT_AGENT_LOG:
				default:
					log.Printf("Unknown event: %v", event.Type)
				}

				if str == " " {
					continue
				}

				if str != "" {
					streamChannel <- str
				}
			}
		}
	}()

	return streamChannel, conversationId, taskId
}

// ParseToEventCh 将事件按事件类型解析为不同的结构体(字段更准确、冗余少)
func (c *EventCh) ParseToEventCh() <-chan Event {
	eventChannel := make(chan Event, 500)

	go func() {
		defer func() {
			close(eventChannel)
		}()

		for {
			select {
			case <-c.ctx.Done():
				return

			case data, ok := <-c.ch:
				if !ok {
					return
				}

				event, err := c.processEvent(data)
				if err != nil {
					log.Printf("Error processing event: %v data: %s", err, string(data))
					continue
				}

				eventChannel <- event
			}
		}
	}()

	return eventChannel
}

type Event struct {
	Type string
	Data interface{}
}

type eventType struct {
	Event string `json:"event"` // 事件类型
}

type eventCreator func() interface{}

var eventTypeMap = map[string]eventCreator{
	EVENT_ERROR:                    func() interface{} { return &EventError{} },
	EVENT_MESSAGE:                  func() interface{} { return &EventMessage{} },
	EVENT_MESSAGE_END:              func() interface{} { return &EventMessageEnd{} },
	EVENT_TTS_MESSAGE:              func() interface{} { return &EventTtsMessage{} },
	EVENT_TTS_MESSAGE_END:          func() interface{} { return &EventTtsMessageEnd{} },
	EVENT_MESSAGE_FILE:             func() interface{} { return &EventMessageFile{} },
	EVENT_MESSAGE_REPLACE:          func() interface{} { return &EventMessageReplace{} },
	EVENT_AGENT_THOUGHT:            func() interface{} { return &EventAgentThought{} },
	EVENT_AGENT_MESSAGE:            func() interface{} { return &EventAgentMessage{} },
	EVENT_WORKFLOW_STARTED:         func() interface{} { return &EventWorkflowStarted{} },
	EVENT_WORKFLOW_FINISHED:        func() interface{} { return &EventWorkflowFinished{} },
	EVENT_NODE_STARTED:             func() interface{} { return &EventNodeStarted{} },
	EVENT_NODE_FINISHED:            func() interface{} { return &EventNodeFinished{} },
	EVENT_NODE_RETRY:               func() interface{} { return &EventNodeRetry{} },
	EVENT_PARALLEL_BRANCH_STARTED:  func() interface{} { return &EventParallelBranchStarted{} },
	EVENT_PARALLEL_BRANCH_FINISHED: func() interface{} { return &EventParallelBranchFinished{} },
	EVENT_ITERATION_STARTED:        func() interface{} { return &EventIterationStarted{} },
	EVENT_ITERATION_NEXT:           func() interface{} { return &EventIterationNext{} },
	EVENT_ITERATION_COMPLETED:      func() interface{} { return &EventIterationCompleted{} },
	EVENT_LOOP_STARTED:             func() interface{} { return &EventLoopStarted{} },
	EVENT_LOOP_NEXT:                func() interface{} { return &EventLoopNext{} },
	EVENT_LOOP_COMPLETED:           func() interface{} { return &EventLoopCompleted{} },
	EVENT_TEXT_CHUNK:               func() interface{} { return &EventTextChunk{} },
	EVENT_TEXT_REPLACE:             func() interface{} { return &EventTextReplace{} },
	EVENT_AGENT_LOG:                func() interface{} { return &EventAgentLog{} },
}

func (c *EventCh) processEvent(data []byte) (Event, error) {
	var event eventType
	if err := json.Unmarshal(data, &event); err != nil {
		return Event{}, fmt.Errorf("failed to unmarshal event type: %w", err)
	}

	creator, exists := eventTypeMap[event.Event]
	if !exists {
		return Event{}, fmt.Errorf("unknown event type: %s", event.Event)
	}

	targetType := creator()

	if err := json.Unmarshal(data, targetType); err != nil {
		return Event{}, fmt.Errorf("failed to unmarshal event data: %w", err)
	}

	return Event{
		Type: event.Event,
		Data: targetType,
	}, nil
}
