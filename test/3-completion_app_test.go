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

func TestCompletionApp(t *testing.T) {
	client, err := dify.NewClient(dify.ClientConfig{
		ApiServer: os.Getenv("APIServer"),
		ApiKey:    os.Getenv("CompletionApiKey"),
		User:      "completion-demo",
		Debug:     true,
		Timeout:   time.Second * 180,
		Transport: nil,
	})
	if err != nil {
		t.Fatal("初始化客户端失败：", err)
	}

	t.Run("Completion_RunBlock", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		resp, err := client.CompletionApp().RunBlock(ctx, types.CompletionRequest{
			Query:        "你好!你知道我是谁么？",
			Inputs:       input,
			ResponseMode: "",
			User:         "golang-test-completion",
			Files:        nil,
		})
		t.Logf("resp=%+v err=%v", resp, err)
	})

	t.Run("Completion_Run_ParseToStructCh", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh := client.CompletionApp().Run(ctx, types.CompletionRequest{
			Query:        "帮我构思一个国庆五天的出游计划，尽可能详细一点",
			Inputs:       input,
			ResponseMode: "",
			User:         "",
			Files:        nil,
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
				t.Log(msg.Answer)
			}
		}
	})

	t.Run("Completion_Run_SimplePrint", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh, meta := client.DebugOff().CompletionApp().Run(ctx, types.CompletionRequest{
			Query:        "你知道现在的时间以及星期么？",
			Inputs:       input,
			ResponseMode: "",
			User:         "",
			Files:        nil,
		}).SimplePrint()

		// 方式一
		for msg := range eventCh {
			fmt.Printf("%s", msg)
		}
		fmt.Printf("\n本次会话conversationId=%s taskId=%s\n", meta.ConversationId, meta.TaskId)

		// 方式二
		//for {
		//	select {
		//	case msg, ok := <-eventCh:
		//		if !ok {
		//			fmt.Printf("\n本次会话conversationId=%s taskId=%s\n", meta.ConversationId, meta.TaskId)
		//			return
		//		}
		//		fmt.Printf("%s", msg)
		//	}
		//}
	})

	t.Run("Completion_Run_ParseToEventCh", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh := client.CompletionApp().Run(ctx, types.CompletionRequest{
			Query:        "你知道现在的时间以及星期么？",
			Inputs:       input,
			ResponseMode: "",
			User:         "",
			Files:        nil,
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

	t.Run("Completion_Run_Stop", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["name"] = "张三"

		eventCh := client.CompletionApp().Run(ctx, types.CompletionRequest{
			Query:        "帮我构思一个国庆五天的出游计划，尽可能详细一点",
			Inputs:       input,
			ResponseMode: "",
			User:         "",
			Files:        nil,
		}).ParseToStructCh()
		cnt := 0
		for {
			select {
			case msg, ok := <-eventCh:
				if !ok {
					return
				}
				t.Log(msg.Answer)
				cnt++
				if cnt == 4 {
					err := client.CompletionApp().Stop(msg.TaskId, "")
					t.Logf("err=%v", err)
				}
			}
		}
	})

	t.Run("Completion_UploadFile", func(t *testing.T) {
		f, err := os.Open("/Users/alsc/Downloads/abcd")
		if err != nil {
			t.Logf("Error opening file: %v\n", err)
			return
		}
		defer f.Close()

		resp, err := client.CompletionApp().UploadFile(
			"/Users/alsc/Downloads/会议室分布.png",
			nil,
			"",
		)
		t.Logf("resp=%v err=%v\n", resp, err)

		resp, err = client.CompletionApp().UploadFile(
			"",
			f,
			"",
		)
		t.Logf("resp=%v err=%v\n", resp, err)
	})

	t.Run("Completion_AppInfo", func(t *testing.T) {
		resp, err := client.CompletionApp().AppInfo()
		t.Logf("resp=%+v err=%v", resp, err)
	})

	t.Run("Completion_AppParameter", func(t *testing.T) {
		resp, err := client.CompletionApp().AppParameter()
		t.Logf("resp=%+v err=%v", resp, err)
	})

	t.Run("Completion_MsgFeedback", func(t *testing.T) {
		err := client.CompletionApp().MsgFeedback(types.FeedbackReq{
			MessageId: "a935054f-e2e9-4ca2-adef-f2c11af117a8",
			Rating:    types.MsgFeedbackNull,
			User:      "",
			Content:   "非常不错",
		})
		t.Logf("err=%v", err)
	})

	t.Run("Completion_TextToAudio", func(t *testing.T) {
		err := client.CompletionApp().TextToAudio(types.Text2Audio{
			MessageId: "",
			Text:      "你是谁？今天是几号",
			User:      "",
		})
		t.Logf("resp=%v err=%v", "", err)
	})

	t.Run("Completion_AnnotationList", func(t *testing.T) {
		resp, err := client.CompletionApp().AnnotationList(0, 0)
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Completion_AnnotationCreate", func(t *testing.T) {
		resp, err := client.CompletionApp().AnnotationCreate("我的问题", "我的答案")
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Completion_AnnotationUpdate", func(t *testing.T) {
		resp, err := client.CompletionApp().AnnotationUpdate("我的问题222", "aaa", "044d9e78-128e-4dd2-80c0-6fd5b7053b3f")
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Completion_AnnotationDel", func(t *testing.T) {
		err := client.CompletionApp().AnnotationDel("044d9e78-128e-4dd2-80c0-6fd5b7053b3f")
		t.Logf("err=%v", err)
	})

	t.Run("Completion_AnnotationReplySetting", func(t *testing.T) {
		resp, err := client.CompletionApp().AnnotationReplySetting(types.AnnotationEnable, types.AnnotationSetting{
			EmbeddingProviderName: "langgenius/tongyi/tongyi",
			EmbeddingModelName:    "text-embedding-v1",
			ScoreThreshold:        0.8,
		})
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Completion_AnnotationReplySettingJobStatus", func(t *testing.T) {
		resp, err := client.CompletionApp().AnnotationReplySettingJobStatus(types.AnnotationEnable, "0af37662-561b-48ec-977c-c87d4d99b228")
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Completion_AppSite", func(t *testing.T) {
		resp, err := client.CompletionApp().AppSite()
		t.Logf("resp=%v err=%v", resp, err)
	})

	t.Run("Completion_AppFeedback", func(t *testing.T) {
		resp, err := client.CompletionApp().AppFeedback(1, 20)
		t.Logf("resp=%v err=%v", resp, err)
	})
}
