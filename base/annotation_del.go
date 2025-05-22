package base

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// AnnotationDel 删除标注
func (c *AppClient) AnnotationDel(annotationId string) error {
	if annotationId == "" {
		return fmt.Errorf("annotationId can not be empty")
	}

	type Resp struct {
		Result string `json:"result"`
	}

	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodDelete, fmt.Sprintf("/apps/annotations/%s", annotationId), nil)
	if err != nil {
		return err
	}
	var resp Resp
	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		if err == io.EOF { // HTTP/1.1 204 NO CONTENT
			return nil
		}
		return err
	}
	if resp.Result != "success" {
		return fmt.Errorf("err resp=%v", resp)
	}
	return nil
}
