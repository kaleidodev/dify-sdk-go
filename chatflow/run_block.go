package chatflow

import (
	"context"
	"net/http"

	"github.com/safejob/dify-sdk-go/types"
)

// RunBlock 发送对话消息(阻塞式)
func (c *App) RunBlock(ctx context.Context, req types.ChatRequest) (resp types.ChatCompletionResponse, err error) {
	req.ResponseMode = "blocking"

	if req.User == "" {
		req.User = c.GetUser()
	}

	if req.Inputs == nil {
		req.Inputs = make(map[string]interface{})
	}

	httpReq, err := c.client.CreateBaseRequest(ctx, http.MethodPost, "/chat-messages", req)
	if err != nil {
		return
	}
	err = c.client.SendJSONRequest(httpReq, &resp)
	return
}
