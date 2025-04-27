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

const ErrEventStr = `{"event":"error","status":%d,"code":"%s","message":"%s"}`

const (
	EVENT_PING                     = "ping"
	EVENT_ERROR                    = "error"
	EVENT_MESSAGE                  = "message"
	EVENT_MESSAGE_END              = "message_end"
	EVENT_TTS_MESSAGE              = "tts_message"
	EVENT_TTS_MESSAGE_END          = "tts_message_end"
	EVENT_MESSAGE_FILE             = "message_file"
	EVENT_MESSAGE_REPLACE          = "message_replace"
	EVENT_AGENT_THOUGHT            = "agent_thought"
	EVENT_AGENT_MESSAGE            = "agent_message"
	EVENT_WORKFLOW_STARTED         = "workflow_started"
	EVENT_WORKFLOW_FINISHED        = "workflow_finished"
	EVENT_NODE_STARTED             = "node_started"
	EVENT_NODE_FINISHED            = "node_finished"
	EVENT_NODE_RETRY               = "node_retry"
	EVENT_PARALLEL_BRANCH_STARTED  = "parallel_branch_started"
	EVENT_PARALLEL_BRANCH_FINISHED = "parallel_branch_finished"
	EVENT_ITERATION_STARTED        = "iteration_started"
	EVENT_ITERATION_NEXT           = "iteration_next"
	EVENT_ITERATION_COMPLETED      = "iteration_completed"
	EVENT_LOOP_STARTED             = "loop_started"
	EVENT_LOOP_NEXT                = "loop_next"
	EVENT_LOOP_COMPLETED           = "loop_completed"
	EVENT_TEXT_CHUNK               = "text_chunk"
	EVENT_TEXT_REPLACE             = "text_replace"
	EVENT_AGENT_LOG                = "agent_log"
)

type AnnotationAction string

const (
	AnnotationEnable  AnnotationAction = "enable"
	AnnotationDisable AnnotationAction = "disable"
)
