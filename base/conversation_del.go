package base

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// ConversationDel 删除会话
func (c *AppClient) ConversationDel(conversationId, user string) error {
	if user == "" {
		user = c.GetUser()
	}

	type Data struct {
		User string `json:"user"`
	}
	type Resp struct {
		Result string `json:"result"`
	}
	req := Data{
		User: user,
	}

	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodDelete, fmt.Sprintf("/conversations/%s", conversationId), req)
	if err != nil {
		return err
	}
	var resp Resp
	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		if err == io.EOF { // HTTP/1.1 204 NO CONTENT
			return nil
		}
		return err
	}
	if resp.Result != "success" {
		return fmt.Errorf("err resp=%v", resp)
	}
	return nil
}
