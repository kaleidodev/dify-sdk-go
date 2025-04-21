package completion

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/safejob/dify-sdk-go/types"
)

// Run 发送对话消息(流式)
func (c *App) Run(ctx context.Context, req types.CompletionRequest) (chan types.ChunkChatCompletionResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	req.ResponseMode = "streaming"

	if req.User == "" {
		req.User = c.GetUser()
	}

	if req.Inputs == nil {
		req.Inputs = make(map[string]interface{})
	}
	_, ok := req.Inputs["query"]
	if !ok {
		req.Inputs["query"] = req.Query
	}

	httpResp, err := c.client.SendRawRequest(ctx, http.MethodPost, "/completion-messages", req)
	if err != nil {
		return nil, err
	}

	dataChan := c.client.SSEEventHandle(ctx, httpResp)

	streamChannel := make(chan types.ChunkChatCompletionResponse)
	go c.chatMessagesStreamHandle(ctx, dataChan, streamChannel)

	return streamChannel, nil
}

func (c *App) chatMessagesStreamHandle(ctx context.Context, dataChan chan []byte, streamChannel chan types.ChunkChatCompletionResponse) {
	if ctx == nil {
		ctx = context.Background()
	}

	defer func() {
		close(streamChannel)
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			data, ok := <-dataChan
			if !ok {
				return
			}

			var resp types.ChunkChatCompletionResponse
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
}
