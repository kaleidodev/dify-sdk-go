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

func TestAgentApp(t *testing.T) {
	client, err := dify.NewClient(dify.ClientConfig{
		ApiServer: os.Getenv("APIServer"),
		ApiKey:    os.Getenv("AgentApiKey"),
		User:      "agent-demo",
		Debug:     true,
		Timeout:   time.Second * 180,
		Transport: nil,
	})
	if err != nil {
		t.Fatal("初始化客户端失败：", err)
	}

	t.Run("Agent_RunBlock", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		resp, err := client.AgentApp().RunBlock(ctx, types.ChatRequest{
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

	t.Run("Agent_Run_ParseToStructCh-1", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh := client.AgentApp().Run(ctx, types.ChatRequest{
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
					t.Logf("|%s|", msg.Answer)
				}
			}
		}
	})

	t.Run("Agent_Run_ParseToStructCh-2", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh := client.AgentApp().Run(ctx, types.ChatRequest{
			Query:            "帮我构思一个国庆五天的出游计划，尽可能详细一点",
			Inputs:           input,
			ResponseMode:     "",
			User:             "",
			ConversationId:   "",
			Files:            nil,
			AutoGenerateName: nil,
		}).ParseToStructCh()
		t.Logf("err=%v", err)
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

	t.Run("Agent_Run_SimplePrint", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh, conversationId := client.DebugOff().AgentApp().Run(ctx, types.ChatRequest{
			Query:            "你知道现在的时间以及星期么？",
			Inputs:           input,
			ResponseMode:     "",
			User:             "",
			ConversationId:   "",
			Files:            nil,
			AutoGenerateName: nil,
		}).SimplePrint()

		// 方式一
		for msg := range eventCh {
			fmt.Printf("%s", msg)
		}
		fmt.Printf("\n本次会话conversationId=%s\n", *conversationId)

		// 方式二
		//for {
		//	select {
		//	case msg, ok := <-eventCh:
		//		if !ok {
		//			fmt.Printf("\n本次会话conversationId=%s\n", *conversationId)
		//			return
		//		}
		//		fmt.Printf("%s", msg)
		//	}
		//}
	})

	t.Run("Agent_Run_ParseToEventCh", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh := client.AgentApp().Run(ctx, types.ChatRequest{
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

	t.Run("Agent_Run_Stop", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh := client.AgentApp().Run(ctx, types.ChatRequest{
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
					err := client.AgentApp().Stop(msg.TaskId, "")
					t.Logf("err=%v", err)
				}
			}
		}
	})

	t.Run("Agent_UploadFile", func(t *testing.T) {
		f, err := os.Open("/Users/alsc/Downloads/abcd")
		if err != nil {
			t.Logf("Error opening file: %v\n", err)
			return
		}
		defer f.Close()

		resp, err := client.AgentApp().UploadFile(
			"/Users/alsc/Downloads/会议室分布.png",
			nil,
			"",
		)
		t.Logf("resp=%v err=%v\n", resp, err)

		resp, err = client.AgentApp().UploadFile(
			"",
			f,
			"",
		)
		t.Logf("resp=%v err=%v\n", resp, err)
	})

	t.Run("Agent_AppInfo", func(t *testing.T) {
		resp, err := client.AgentApp().AppInfo()
		t.Logf("resp=%+v err=%v", resp, err)
	})

	t.Run("Agent_AppParameter", func(t *testing.T) {
		resp, err := client.AgentApp().AppParameter()
		t.Logf("resp=%+v err=%v", resp, err)
	})

	t.Run("Agent_MsgFeedback", func(t *testing.T) {
		err := client.AgentApp().MsgFeedback(types.FeedbackReq{
			MessageId: "a89094dd-8dac-4b51-aa77-920099ae4ef9",
			Rating:    types.MsgFeedbackNull,
			User:      "",
			Content:   "非常不错",
		})
		t.Logf("err=%v", err)
	})

	t.Run("Agent_SuggestQuestionList", func(t *testing.T) {
		resp, err := client.AgentApp().SuggestQuestionList("c71918e4-bb23-4ff9-bb63-e5fa5aaf6afa", "")
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Agent_History", func(t *testing.T) {
		resp, err := client.AgentApp().HistoryPro("0a9a0917-0c36-4121-8934-17367bb803c0", "", "", 20)
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Agent_ConversationList", func(t *testing.T) {
		resp, err := client.AgentApp().ConversationList("")
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Agent_ConversationDel", func(t *testing.T) {
		err := client.AgentApp().ConversationDel("adc2ad24-fa4e-4dbb-8c16-ead1eaaa6c38", "")
		t.Logf("err=%v", err)
	})

	t.Run("Agent_ConversationRename", func(t *testing.T) {
		resp, err := client.AgentApp().ConversationRename(types.ConversationRenameReq{
			ConversationId: "f6da1bba-6341-42ed-9021-4a88b2f0dd0a",
			Name:           "修改后的新名称",
			AutoGenerate:   false,
			User:           "",
		})
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Agent_ConversationVars", func(t *testing.T) {
		resp, err := client.AgentApp().ConversationVars("936221e9-1779-467a-bbfb-319ecae9864f", "", "", 0)
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Agent_AudioToText", func(t *testing.T) {
		client.AgentApp().AudioToText("", nil, "")
	})

	t.Run("Agent_TextToAudio", func(t *testing.T) {
		err := client.AgentApp().TextToAudio(types.Text2Audio{
			MessageId: "",
			Text:      "你是谁？今天是几号",
			User:      "",
		})
		t.Logf("resp=%v err=%v", "", err)
	})

	t.Run("Agent_AppMeta", func(t *testing.T) {
		resp, err := client.AgentApp().AppMeta()
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Agent_AppSite", func(t *testing.T) {
		resp, err := client.AgentApp().AppSite()
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Agent_AppFeedback", func(t *testing.T) {
		resp, err := client.AgentApp().AppFeedback(1, 20)
		t.Logf("resp=%v err=%v", resp, err)
	})
}
