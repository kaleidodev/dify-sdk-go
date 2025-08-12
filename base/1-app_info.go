package base

import (
	"context"
	"net/http"

	"github.com/kaleidodev/dify-sdk-go/types"
)

// AppInfo 获取应用基本信息
func (c *AppClient) AppInfo() (resp types.AppInfo, err error) {

	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodGet, "/info", nil)
	if err != nil {
		return
	}

	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)

	return
}
