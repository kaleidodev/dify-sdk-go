package chatbot

import (
	"context"
	"fmt"
	"net/http"
	"strings"

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
		ch := make(chan []byte, 2)
		// err.Error()="Post "http://op-dify-gld.keruyun.com/v1/chat-messages": http: ContentLength=100 with Body length 0" 替换双引号,以解决json.Unmarshal报错的问题
		ch <- []byte(fmt.Sprintf(types.ErrEventStr, 500, "request err", strings.Replace(err.Error(), `"`, `\"`, -1)))
		close(ch)

		return types.NewEventCh(ch, ctx)
	}

	return types.NewEventCh(c.client.SSEEventHandle(ctx, httpResp), ctx)
}
