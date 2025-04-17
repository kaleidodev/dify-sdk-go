package interfaces

import (
	"context"
	"net/http"
	"os"

	"github.com/safejob/dify-sdk-go/types"
)

// ClientInterface http工具函数
type ClientInterface interface {
	CreateBaseRequest(ctx context.Context, method, apiUrl string, body interface{}) (*http.Request, error)
	CreateFormRequest(ctx context.Context, method, apiUrl string, data map[string]string) (*http.Request, error)
	SendJSONRequest(req *http.Request, res interface{}) error
	SendRawRequest(ctx context.Context, method, apiUrl string, req interface{}) (*http.Response, error)
	SSEEventHandle(ctx context.Context, resp *http.Response) (ch chan []byte)
	//SendRequest(req *http.Request) (*http.Response, error)
}

// AppCommon 应用通用函数
type AppCommon interface {
	GetUser() string

	UploadFile(filePath string, f *os.File, user string) (info types.FileInfo, err error)
	AppInfo() (resp types.AppInfo, err error)
	AppParameter() (resp types.AppParameter, err error)
}

// Chatbot Chatbot和Agent类型应用
type Chatbot interface {
	AppCommon
	MsgFeedback(feedback types.FeedbackReq) error
	SuggestQuestionList(messageId, user string) (suggest []string, err error)
	History(conversationId, user string) (resp types.MessageHistory, err error)
	HistoryPro(conversationId, user, firstId string, limit int64) (resp types.MessageHistory, err error)
	ConversationList(user string) (resp types.ConversationListResp, err error)
	ConversationListPro(user, lastId, sortBy string, limit int64) (resp types.ConversationListResp, err error)
	ConversationDel(conversationId, user string) error
	ConversationRename(rename types.ConversationRenameReq) (resp types.ConversationRenameResp, err error)
	AudioToText(filePath string, f *os.File, user string) (text string, err error)
	TextToAudio(info types.Text2Audio) error
	AppMeta() (resp types.AppMeta, err error)
}

// Completion 类型应用
type Completion interface {
	AppCommon
	MsgFeedback(feedback types.FeedbackReq) error
	TextToAudio(info types.Text2Audio) error
}

// Chatflow  类型应用
type Chatflow interface {
	AppCommon
	MsgFeedback(feedback types.FeedbackReq) error
	SuggestQuestionList(messageId, user string) (suggest []string, err error)
	History(conversationId, user string) (resp types.MessageHistory, err error)
	HistoryPro(conversationId, user, firstId string, limit int64) (resp types.MessageHistory, err error)
	ConversationList(user string) (resp types.ConversationListResp, err error)
	ConversationListPro(user, lastId, sortBy string, limit int64) (resp types.ConversationListResp, err error)
	ConversationDel(conversationId, user string) error
	ConversationRename(rename types.ConversationRenameReq) (resp types.ConversationRenameResp, err error)
	AudioToText(filePath string, f *os.File, user string) (text string, err error)
	TextToAudio(info types.Text2Audio) error
	AppMeta() (resp types.AppMeta, err error)
}

// Workflow 类型应用
type Workflow interface {
	AppCommon
}
