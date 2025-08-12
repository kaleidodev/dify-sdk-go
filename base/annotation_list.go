package base

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kaleidodev/dify-sdk-go/types"
)

// AnnotationList 获取标注列表
func (c *AppClient) AnnotationList(page, limit int) (resp types.AnnotationListResp, err error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}

	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodGet, fmt.Sprintf("/apps/annotations?page=%d&limit=%d", page, limit), nil)
	if err != nil {
		return
	}

	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}

	return
}
