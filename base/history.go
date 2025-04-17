package base

import (
	"context"
	"fmt"
	"net/http"

	"github.com/safejob/dify-sdk-go/types"
)

// History 获取会话历史消息(倒序返回20条)
func (c *AppClient) History(conversationId, user string) (resp types.MessageHistory, err error) {
	if user == "" {
		user = c.GetUser()
	}

	return c.HistoryPro(conversationId, user, "", 20)
}

// HistoryPro 获取会话历史消息(倒序返回limit条)
func (c *AppClient) HistoryPro(conversationId, user, firstId string, limit int64) (resp types.MessageHistory, err error) {
	if user == "" {
		user = c.GetUser()
	}
	
	if limit <= 0 {
		limit = 20
	}
	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodGet, fmt.Sprintf("/messages?conversation_id=%s&user=%s&first_id=%s&limit=%d",
		conversationId, user, firstId, limit), nil)
	if err != nil {
		return
	}

	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}

	return
}
