package base

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kaleidodev/dify-sdk-go/types"
)

// AnnotationReplySettingJobStatus 查询标注回复初始设置任务状态
func (c *AppClient) AnnotationReplySettingJobStatus(action types.AnnotationAction, jobId string) (resp types.AnnotationSettingJobStatusResp, err error) {
	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodGet, fmt.Sprintf("/apps/annotation-reply/%s/status/%s", action, jobId), nil)
	if err != nil {
		return
	}
	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}

	return
}
