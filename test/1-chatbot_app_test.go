package test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/safejob/dify-sdk-go"
	"github.com/safejob/dify-sdk-go/types"
)

func TestChatbotApp(t *testing.T) {
	client, err := dify.NewClient(dify.ClientConfig{
		ApiServer: os.Getenv("APIServer"),
		ApiKey:    os.Getenv("ChatbotApiKey"),
		User:      "chatbot-demo",
		Debug:     true,
		Timeout:   time.Second * 180,
		Transport: nil,
	})
	if err != nil {
		t.Fatal("初始化客户端失败：", err)
	}

	t.Run("Chatbot_RunBlock", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		resp, err := client.ChatbotApp().RunBlock(ctx, types.ChatRequest{
			Query:            "你好!你知道我是谁么？",
			Inputs:           input,
			ResponseMode:     "",
			User:             "",
			ConversationId:   "",
			Files:            nil,
			AutoGenerateName: nil,
		})
		t.Logf("resp=%+v err=%v", resp, err)
	})

	t.Run("Chatbot_Run_ParseToStructCh-1", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh := client.ChatbotApp().Run(ctx, types.ChatRequest{
			Query:            "你知道现在的时间么？",
			Inputs:           input,
			ResponseMode:     "",
			User:             "",
			ConversationId:   "",
			Files:            nil,
			AutoGenerateName: nil,
		}).ParseToStructCh()
		for {
			select {
			case msg, ok := <-eventCh:
				if !ok {
					return
				}
				if msg.Event == "error" {
					t.Logf("status=%d code=%s message=%s", msg.Status, msg.Code, msg.Message)
				}
				if msg.Answer != "" {
					t.Log(msg.Answer)
				}
			}
		}
	})

	t.Run("Chatbot_Run_ParseToStructCh-2", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh := client.ChatbotApp().Run(ctx, types.ChatRequest{
			Query:            "帮我构思一个国庆五天的出游计划，尽可能详细一点",
			Inputs:           input,
			ResponseMode:     "",
			User:             "",
			ConversationId:   "",
			Files:            nil,
			AutoGenerateName: nil,
		}).ParseToStructCh()
		for {
			select {
			case msg, ok := <-eventCh:
				if !ok {
					return
				}
				if msg.Event == "error" {
					t.Logf("status=%d code=%s message=%s", msg.Status, msg.Code, msg.Message)
				}
				if msg.Answer != "" {
					t.Log(msg.Answer)
				}
			}
		}
	})

	t.Run("Chatbot_Run_SimplePrint", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh := client.DebugOff().ChatbotApp().Run(ctx, types.ChatRequest{
			Query:            "你知道现在的时间以及星期么？",
			Inputs:           input,
			ResponseMode:     "",
			User:             "",
			ConversationId:   "",
			Files:            nil,
			AutoGenerateName: nil,
		}).SimplePrint()
		for {
			select {
			case msg, ok := <-eventCh:
				if !ok {
					return
				}
				fmt.Printf("%s", msg)
			}
		}
	})

	t.Run("Chatbot_Run_ParseToEventCh", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh := client.ChatbotApp().Run(ctx, types.ChatRequest{
			Query:            "你知道现在的时间以及星期么？",
			Inputs:           input,
			ResponseMode:     "",
			User:             "",
			ConversationId:   "",
			Files:            nil,
			AutoGenerateName: nil,
		}).ParseToEventCh()
		for {
			select {
			case msg, ok := <-eventCh:
				if !ok {
					return
				}
				t.Logf("====>event: %s %+v\n", msg.Type, msg.Data)
			}
		}
	})

	t.Run("Chatbot_Run_Stop", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh := client.ChatbotApp().Run(ctx, types.ChatRequest{
			Query:            "帮我构思一个国庆五天的出游计划，尽可能详细一点",
			Inputs:           input,
			ResponseMode:     "",
			User:             "",
			ConversationId:   "",
			Files:            nil,
			AutoGenerateName: nil,
		}).ParseToStructCh()
		cnt := 0
		for {
			select {
			case msg, ok := <-eventCh:
				if !ok {
					return
				}
				if msg.Answer != "" {
					t.Log(msg.Answer)
				}
				cnt++
				if cnt == 4 {
					err := client.ChatbotApp().Stop(msg.TaskId, "")
					t.Logf("err=%v", err)
				}
			}
		}
	})

	t.Run("Chatbot_UploadFile", func(t *testing.T) {
		f, err := os.Open("/Users/alsc/Downloads/abcd")
		if err != nil {
			t.Logf("Error opening file: %v\n", err)
			return
		}
		defer f.Close()

		resp, err := client.ChatbotApp().UploadFile(
			"/Users/alsc/Downloads/会议室分布.png",
			nil,
			"",
		)
		t.Logf("resp=%v err=%v\n", resp, err)

		resp, err = client.ChatbotApp().UploadFile(
			"",
			f,
			"",
		)
		t.Logf("resp=%v err=%v\n", resp, err)
	})

	t.Run("Chatbot_AppInfo", func(t *testing.T) {
		resp, err := client.ChatbotApp().AppInfo()
		t.Logf("resp=%+v err=%v", resp, err)
	})

	t.Run("Chatbot_AppParameter", func(t *testing.T) {
		resp, err := client.ChatbotApp().AppParameter()
		t.Logf("resp=%+v err=%v", resp, err)
	})

	t.Run("Chatbot_MsgFeedback", func(t *testing.T) {
		err := client.ChatbotApp().MsgFeedback(types.FeedbackReq{
			MessageId: "a89094dd-8dac-4b51-aa77-920099ae4ef9",
			Rating:    types.MsgFeedbackNull,
			User:      "",
			Content:   "非常不错",
		})
		t.Logf("err=%v", err)
	})

	t.Run("Chatbot_SuggestQuestionList", func(t *testing.T) {
		resp, err := client.ChatbotApp().SuggestQuestionList("c71918e4-bb23-4ff9-bb63-e5fa5aaf6afa", "")
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Chatbot_History", func(t *testing.T) {
		resp, err := client.ChatbotApp().HistoryPro("0a9a0917-0c36-4121-8934-17367bb803c0", "", "", 20)
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Chatbot_ConversationList", func(t *testing.T) {
		resp, err := client.ChatbotApp().ConversationList("")
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Chatbot_ConversationDel", func(t *testing.T) {
		err := client.ChatbotApp().ConversationDel("adc2ad24-fa4e-4dbb-8c16-ead1eaaa6c38", "")
		t.Logf("err=%v", err)
	})

	t.Run("Chatbot_ConversationRename", func(t *testing.T) {
		resp, err := client.ChatbotApp().ConversationRename(types.ConversationRenameReq{
			ConversationId: "f6da1bba-6341-42ed-9021-4a88b2f0dd0a",
			Name:           "修改后的新名称",
			AutoGenerate:   false,
			User:           "",
		})
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Chatbot_ConversationVars", func(t *testing.T) {
		resp, err := client.ChatbotApp().ConversationVars("0dc15d8a-d218-44df-9ea1-c807837bea90", "", "", 0)
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Chatbot_AudioToText", func(t *testing.T) {
		client.ChatbotApp().AudioToText("", nil, "")
	})

	t.Run("Chatbot_TextToAudio", func(t *testing.T) {
		err := client.ChatbotApp().TextToAudio(types.Text2Audio{
			MessageId: "",
			Text:      "你是谁？今天是几号",
			User:      "",
		})
		t.Logf("resp=%v err=%v", "", err)
	})

	t.Run("Chatbot_AppMeta", func(t *testing.T) {
		resp, err := client.ChatbotApp().AppMeta()
		t.Logf("resp=%v err=%v", resp, err)
	})
}
