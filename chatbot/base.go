package chatbot

import (
	"github.com/kaleidodev/dify-sdk-go/interfaces"
)

type App struct {
	client interfaces.ClientInterface
	interfaces.Chatbot
	appType string // 应用类型 App/Agent
}

const (
	AppTypeChatbot = "Chatbot"
	AppTypeAgent   = "Agent"
)

func NewChatbot(client interfaces.ClientInterface, app interfaces.Chatbot) *App {

	return &App{
		client:  client,
		Chatbot: app,
		appType: AppTypeChatbot,
	}
}

func NewAgent(client interfaces.ClientInterface, app interfaces.Chatbot) *App {
	return &App{
		client:  client,
		Chatbot: app,
		appType: AppTypeAgent,
	}
}
