package base

import (
	"context"
	"net/http"

	"github.com/kaleidodev/dify-sdk-go/types"
)

// TextToAudio 文字转语音
func (c *AppClient) TextToAudio(info types.Text2Audio) error {
	if info.User == "" {
		info.User = c.GetUser()
	}
	type Resp struct {
		Result string `json:"result"`
	}
	req := make(map[string]string)
	req["message_id"] = info.MessageId
	req["text"] = info.Text
	req["user"] = info.User

	httpReq, err := (*Client)(c).HttpClient().CreateFormRequest(context.Background(), http.MethodPost, "/text-to-audio", req)
	if err != nil {
		return err
	}
	var resp Resp
	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return err
	}
	return nil
}

/* 服务端似乎不支持这个接口
状态码: 415
响应内容: {
    "code": "unsupported_media_type",
    "message": "Did not attempt to load JSON data because the request Content-Type was not 'application/json'.",
    "status": 415
}
*/
