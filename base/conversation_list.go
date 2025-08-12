package base

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kaleidodev/dify-sdk-go/types"
)

// ConversationList 获取会话列表
func (c *AppClient) ConversationList(user string) (resp types.ConversationListResp, err error) {
	if user == "" {
		user = c.GetUser()
	}

	return c.ConversationListPro(user, "", "", 20)
}

// ConversationListPro 获取会话列表
// sortBt 默认 -updated_at(按更新时间倒序排列) 可选值：created_at, -created_at, updated_at, -updated_at  -代表倒序
func (c *AppClient) ConversationListPro(user, lastId, sortBy string, limit int64) (resp types.ConversationListResp, err error) {
	if user == "" {
		user = c.GetUser()
	}

	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	if sortBy == "" {
		sortBy = "-updated_at"
	}
	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodGet, fmt.Sprintf("/conversations?&user=%s&last_id=%s&limit=%d&sort_by=%s",
		user, lastId, limit, sortBy), nil)
	if err != nil {
		return
	}

	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}

	return
}
