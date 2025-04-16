package base

import (
	"context"
	"github.com/safejob/dify-sdk-go/types"
	"net/http"
)

// AppParameter 获取应用参数
func (c *AppClient) AppParameter() (resp *types.AppParameter, err error) {

	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodGet, "/parameters", nil)
	if err != nil {
		return
	}

	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)

	return
}
