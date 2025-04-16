package base

import (
	"context"
	"fmt"
	"net/http"

	"github.com/safejob/dify-sdk-go/types"
)

// ConversationRename 会话重命名
func (c *AppClient) ConversationRename(rename *types.ConversationRenameReq) (resp types.ConversationRenameResp, err error) {
	type RenameParm struct {
		Name         string `json:"name,omitempty"`
		AutoGenerate bool   `json:"auto_generate,omitempty"`
		User         string `json:"user"`
	}
	req := RenameParm{
		Name:         rename.Name,
		AutoGenerate: rename.AutoGenerate,
		User:         rename.User,
	}

	httpReq, err := (*Client)(c).HttpClient().CreateBaseRequest(context.Background(), http.MethodPost, fmt.Sprintf("/conversations/%s/name", rename.ConversationId), req)
	if err != nil {
		return
	}
	err = (*Client)(c).HttpClient().SendJSONRequest(httpReq, &resp)
	if err != nil {
		return
	}

	return
}
