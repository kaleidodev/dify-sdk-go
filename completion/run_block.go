package completion

import (
	"context"
	"net/http"

	"github.com/safejob/dify-sdk-go/types"
)

// RunBlock 发送对话消息(阻塞式)
func (c *App) RunBlock(ctx context.Context, req types.CompletionRequest) (resp types.ChatbotCompletionBlockingResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	req.ResponseMode = "blocking"

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

	httpReq, err := c.client.CreateBaseRequest(ctx, http.MethodPost, "/completion-messages", req)
	if err != nil {
		return
	}
	err = c.client.SendJSONRequest(httpReq, &resp)
	return
}
