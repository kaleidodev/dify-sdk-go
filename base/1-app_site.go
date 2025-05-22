package base

import (
	"context"
	"net/http"

	"github.com/safejob/dify-sdk-go/types"
)

// AppSite 获取应用 WebApp 设置
func (c *AppClient) AppSite() (resp types.AppSite, err error) {

	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodGet, "/site", nil)
	if err != nil {
		return
	}

	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)

	return
}
