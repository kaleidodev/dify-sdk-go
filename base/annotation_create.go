package base

import (
	"context"
	"errors"
	"net/http"

	"github.com/kaleidodev/dify-sdk-go/types"
)

// AnnotationCreate 创建标注
func (c *AppClient) AnnotationCreate(question, answer string) (resp types.Annotation, err error) {
	if question == "" || answer == "" {
		return resp, errors.New("empty question or answer")
	}

	type ACParm struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}
	req := ACParm{
		Question: question,
		Answer:   answer,
	}

	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodPost, "/apps/annotations", req)
	if err != nil {
		return
	}
	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}

	return
}
