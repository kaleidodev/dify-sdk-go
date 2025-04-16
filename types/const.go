package types

type Feedback string

const (
	MsgFeedbackLike    Feedback = "like"    // 点赞
	MsgFeedbackDislike Feedback = "dislike" // 点踩
	MsgFeedbackNull    Feedback = "null"    // 撤销点赞
)

type Status string

const (
	StatusSucceeded Status = "succeeded"
	StatusFailed    Status = "failed"
	StatusStopped   Status = "stopped"
)
