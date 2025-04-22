package base

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/safejob/dify-sdk-go/chatbot"
	"github.com/safejob/dify-sdk-go/chatflow"
	"github.com/safejob/dify-sdk-go/completion"
	"github.com/safejob/dify-sdk-go/workflow"
)

type Client struct {
	apiServer  string
	apiKey     string
	user       string
	debug      bool
	timeout    time.Duration
	httpClient *http.Client
}

type HttpClient Client
type AppClient Client

func NewClient(apiServer, apiKey, user string, debug bool, timeout time.Duration, httpClient *http.Client) (*Client, error) {
	if apiServer == "" {
		return nil, errors.New("apiServer is required")
	}

	if apiKey == "" {
		return nil, errors.New("apiKey is required")
	}

	_, err := url.Parse(apiServer)
	if err != nil {
		return nil, fmt.Errorf("invalid API server URL: %w", err)
	}

	return &Client{
		apiServer:  strings.TrimRight(apiServer, "/"),
		apiKey:     apiKey,
		user:       user,
		debug:      debug,
		timeout:    timeout,
		httpClient: httpClient,
	}, nil
}

func (c *Client) HttpClient() *HttpClient {
	return (*HttpClient)(c)
}

func (c *Client) DebugOn() *Client {
	c.debug = true
	return c
}

func (c *Client) DebugOff() *Client {
	c.debug = false
	return c
}

func (c *Client) ChatbotApp() *chatbot.App {
	return chatbot.NewChatbot((*HttpClient)(c), (*AppClient)(c))
}

func (c *Client) AgentApp() *chatbot.App {
	return chatbot.NewAgent((*HttpClient)(c), (*AppClient)(c))
}

func (c *Client) ChatflowApp() *chatflow.App {
	return chatflow.NewChatflow((*HttpClient)(c), (*AppClient)(c))
}

func (c *Client) CompletionApp() *completion.App {
	return completion.NewCompletion((*HttpClient)(c), (*AppClient)(c))
}

func (c *Client) WorkflowApp() *workflow.App {
	return workflow.NewWorkflow((*HttpClient)(c), (*AppClient)(c))
}
