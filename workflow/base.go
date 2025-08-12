package workflow

import "github.com/kaleidodev/dify-sdk-go/interfaces"

type App struct {
	client interfaces.ClientInterface
	interfaces.Workflow
}

func NewWorkflow(client interfaces.ClientInterface, app interfaces.Workflow) *App {
	return &App{
		client:   client,
		Workflow: app,
	}
}
