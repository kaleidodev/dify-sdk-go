package base

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kaleidodev/dify-sdk-go/types"
)

// MsgFeedback 消息反馈（点赞）
func (c *AppClient) MsgFeedback(feedback types.FeedbackReq) error {
	if feedback.User == "" {
		feedback.User = c.GetUser()
	}

	type Resp struct {
		Result string `json:"result"`
	}
	type Feedback struct {
		Rating  types.Feedback `json:"rating,omitempty"`
		User    string         `json:"user"`
		Content string         `json:"content"`
	}
	if feedback.Rating == "null" {
		feedback.Rating = ""
	}
	req := Feedback{
		Rating:  feedback.Rating,
		User:    feedback.User,
		Content: feedback.Content,
	}

	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodPost, fmt.Sprintf("/messages/%s/feedbacks", feedback.MessageId), req)
	if err != nil {
		return err
	}
	var resp Resp
	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return err
	}
	if resp.Result != "success" {
		return fmt.Errorf("err resp=%v", resp)
	}
	return nil
}
