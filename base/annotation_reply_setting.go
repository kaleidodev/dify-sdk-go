package base

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kaleidodev/dify-sdk-go/types"
)

// AnnotationReplySetting 标注回复初始设置
func (c *AppClient) AnnotationReplySetting(action types.AnnotationAction, setting types.AnnotationSetting) (resp types.AnnotationSettingJobResp, err error) {
	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodPost, fmt.Sprintf("/apps/annotation-reply/%s", action), setting)
	if err != nil {
		return
	}
	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}

	return
}
