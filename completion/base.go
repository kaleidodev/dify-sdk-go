package completion

import "github.com/safejob/dify-sdk-go/interfaces"

type App struct {
	client interfaces.ClientInterface
	interfaces.Completion
}

func NewCompletion(client interfaces.ClientInterface, app interfaces.Completion) *App {
	return &App{
		client:     client,
		Completion: app,
	}
}
