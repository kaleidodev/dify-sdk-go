package workflow

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kaleidodev/dify-sdk-go/types"
)

// Status 获取workflow执行情况
func (c *App) Status(workflowRunId string) (resp types.WorkflowStatus, err error) {
	httpReq, err := c.client.CreateBaseRequest(context.Background(), http.MethodGet, fmt.Sprintf("/workflows/run/%s", workflowRunId), nil)
	if err != nil {
		return
	}

	err = c.client.SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}

	return
}
