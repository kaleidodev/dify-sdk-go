package chatflow

import (
	"context"
	"fmt"
	"net/http"

	"github.com/safejob/dify-sdk-go/types"
)

// Run 发送对话消息(流式)
func (c *App) Run(ctx context.Context, req types.ChatRequest) (event *types.EventCh) {
	if ctx == nil {
		ctx = context.Background()
	}

	req.ResponseMode = "streaming"

	if req.User == "" {
		req.User = c.GetUser()
	}

	httpResp, err := c.client.SendRawRequest(ctx, http.MethodPost, "/chat-messages", req)
	if err != nil {
		ch := make(chan []byte, 500)
		ch <- []byte(fmt.Sprintf(types.ErrEventStr, 500, "request err", err.Error()))
		close(ch)

		return types.NewEventCh(ch, ctx)
	}

	return types.NewEventCh(c.client.SSEEventHandle(ctx, httpResp), ctx)
}
