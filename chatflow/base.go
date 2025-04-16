package chatflow

import (
	"github.com/safejob/dify-sdk-go/interfaces"
)

type App struct {
	client interfaces.ClientInterface
	interfaces.Chatflow
}

func NewChatflow(client interfaces.ClientInterface, app interfaces.Chatflow) *App {
	return &App{
		client:   client,
		Chatflow: app,
	}
}
