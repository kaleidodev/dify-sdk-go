package base

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kaleidodev/dify-sdk-go/types"
)

// ConversationVars 获取对话变量 (注意这里是会话变量，不是流程开始设置的变量)
func (c *AppClient) ConversationVars(conversationId, user, lastId string, limit int64) (resp types.ConversationVarsResp, err error) {
	if conversationId == "" {
		return resp, fmt.Errorf("conversationId is empty")
	}
	if user == "" {
		user = c.user
	}
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodGet, fmt.Sprintf("/conversations/%s/variables?user=%s&last_id=%s&limit=%d", conversationId, user, lastId, limit), nil)
	if err != nil {
		return
	}
	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}

	return
}
