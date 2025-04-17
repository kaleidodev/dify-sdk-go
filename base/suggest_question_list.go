package base

import (
	"context"
	"fmt"
	"net/http"
)

// SuggestQuestionList 获取下一轮建议问题列表(应用需要开启下一步问题建议)
func (c *AppClient) SuggestQuestionList(messageId, user string) (suggest []string, err error) {
	if user == "" {
		user = c.GetUser()
	}
	
	type Resp struct {
		Result string   `json:"result"`
		Data   []string `json:"data"`
	}
	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodGet, fmt.Sprintf("/messages/%s/suggested?user=%s", messageId, user), nil)
	if err != nil {
		return
	}
	var resp Resp
	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}
	if resp.Result != "success" {
		err = fmt.Errorf("err resp=%v", resp)
		return
	}
	return resp.Data, nil
}
