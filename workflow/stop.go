package workflow

import (
	"context"
	"fmt"
	"net/http"
)

// Stop 停止响应
func (c *App) Stop(taskId, user string) error {
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

	httpReq, err := c.client.CreateBaseRequest(context.Background(), http.MethodPost, fmt.Sprintf("/workflows/tasks/%s/stop", taskId), req)
	if err != nil {
		return err
	}
	var resp Resp
	err = c.client.SendJSONRequest(httpReq, &resp)
	if err != nil {
		return err
	}
	if resp.Result != "success" {
		return fmt.Errorf("err resp=%v", resp)
	}
	return nil
}
