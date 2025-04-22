package main

import (
	"context"
	"log"

	"github.com/safejob/dify-sdk-go"
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

	// 获取应用基本信息
	appInfo, err := client.ChatbotApp().AppInfo()
	if err != nil {
		log.Fatalf("Error getting app info: %v", err)
	}
	log.Printf("应用名称：%s 应用描述：%s \n", appInfo.Name, appInfo.Description)

	// 阻塞式调用示例
	ctx := context.Background()
	resp, err := client.ChatbotApp().RunBlock(ctx, types.ChatRequest{
		Query: "请帮我生成五一假期的出行计划",
	})
	if err != nil {
		log.Fatalf("Error running client: %v", err)
	}

	log.Printf("响应结果:\n %s", resp.Answer)
}
