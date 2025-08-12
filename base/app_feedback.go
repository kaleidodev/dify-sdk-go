package base

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kaleidodev/dify-sdk-go/types"
)

// AppFeedback 获取APP的消息点赞和反馈
func (c *AppClient) AppFeedback(page, limit int64) (resp types.AppFeedbackResp, err error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}

	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodGet, fmt.Sprintf("/app/feedbacks?page=%d&limit=%d", page, limit), nil)
	if err != nil {
		return
	}

	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)

	return
}
