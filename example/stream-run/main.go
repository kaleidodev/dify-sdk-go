package main

import (
	"context"
	"fmt"
	"log"

	"github.com/safejob/dify-sdk-go"
	"github.com/safejob/dify-sdk-go/base"
	"github.com/safejob/dify-sdk-go/types"
)

func main() {
	// 构建客户端
	client, err := dify.NewClient(dify.ClientConfig{
		ApiServer: "http://your.domain/v1",
		ApiKey:    "your-api-key",
		User:      "demo-client",
	})
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	ctx := context.Background()
	input := make(map[string]interface{})
	input["name"] = "张三" // 这里根据实际dify配置，填写需要使用的变量
	request := types.ChatRequest{
		Query:            "帮我构思一个国庆五天的出游计划，尽可能详细一点",
		Inputs:           input,
		ResponseMode:     "",
		User:             "",
		ConversationId:   "",
		Files:            nil,
		AutoGenerateName: nil,
	}

	choise := 1
	switch choise {
	case 1:
		// ParseToStructCh调用示例
		ParseToStructChDemo(ctx, client, request)
	case 2:
		// SimplePrint调用示例
		SimplePrintDemo(ctx, client, request)
	case 3:
		// ParseToEventCh调用示例
		ParseToEventChDemo(ctx, client, request)
	}
}

// ParseToStructCh调用示例
func ParseToStructChDemo(ctx context.Context, client *base.Client, request types.ChatRequest) {
	eventCh := client.AgentApp().Run(ctx, request).ParseToStructCh()
	for {
		select {
		case msg, ok := <-eventCh:
			// 这里的msg是一个大结构体 字段非常多
			if !ok {
				return
			}
			if msg.Event == "error" {
				log.Printf("status=%d code=%s message=%s", msg.Status, msg.Code, msg.Message)
			}
			if msg.Answer != "" {
				fmt.Printf("%s", msg.Answer)
			}
		}
	}
}

// SimplePrint调用示例
func SimplePrintDemo(ctx context.Context, client *base.Client, request types.ChatRequest) {
	eventCh, conversationId := client.AgentApp().Run(ctx, request).SimplePrint()
	for {
		select {
		case msg, ok := <-eventCh:
			// 这里的msg是字符串
			if !ok {
				fmt.Printf("本次会话conversationId=%s", *conversationId)
				return
			}
			fmt.Printf("%s", msg)
		}
	}
}

// ParseToEventCh调用示例
func ParseToEventChDemo(ctx context.Context, client *base.Client, request types.ChatRequest) {
	eventCh := client.AgentApp().Run(ctx, request).ParseToEventCh()
	for {
		select {
		case msg, ok := <-eventCh:
			// 这里的msg是具体的结构体，需要根据msg.Type的值进行断言
			if !ok {
				return
			}
			switch msg.Type {
			case types.EVENT_AGENT_THOUGHT:
				event := msg.Data.(*types.EventAgentThought)
				fmt.Printf("agent thought: %s", event.Thought)
			case types.EVENT_AGENT_MESSAGE:
				event := msg.Data.(*types.EventAgentMessage)
				fmt.Printf("%s", event.Answer)
			}
		}
	}
}
