package workflow

import (
	"context"
	"fmt"
	"net/http"

	"github.com/safejob/dify-sdk-go/types"
)

// Logs 获取 workflow 日志
func (c *App) Logs(keyword string, status types.Status, page, limit int) (resp types.WorkflowLogs, err error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}

	httpReq, err := c.client.CreateBaseRequest(context.Background(), http.MethodGet, fmt.Sprintf("/workflows/logs?keyword=%s&status=%s&page=%d&limit=%d",
		keyword, status, page, limit), nil)
	if err != nil {
		return
	}

	err = c.client.SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}

	return
}
