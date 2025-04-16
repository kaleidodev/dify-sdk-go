package base

import (
	"context"
	"net/http"

	"github.com/safejob/dify-sdk-go/types"
)

// AppMeta 获取应用Meta信息(获取工具icon)
func (c *AppClient) AppMeta() (resp types.AppMeta, err error) {
	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodGet, "/meta", nil)
	if err != nil {
		return
	}

	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}

	return
}
