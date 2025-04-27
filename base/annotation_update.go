package base

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/safejob/dify-sdk-go/types"
)

// AnnotationUpdate 更新标注
func (c *AppClient) AnnotationUpdate(question, answer, annotationId string) (resp types.Annotation, err error) {
	if question == "" || answer == "" {
		return resp, errors.New("empty question or answer")
	}

	if annotationId == "" {
		err = errors.New("annotation id can not be empty")
	}

	type ACParm struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}
	req := ACParm{
		Question: question,
		Answer:   answer,
	}

	// todo 文档中是PUT方法 实际需要使用POST方法
	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodPost, fmt.Sprintf("/apps/annotations/%s", annotationId), req)
	if err != nil {
		return
	}
	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}

	return
}
