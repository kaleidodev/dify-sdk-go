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

func TestWorkflowApp(t *testing.T) {
	client, err := dify.NewClient(dify.ClientConfig{
		ApiServer: os.Getenv("APIServer"),
		ApiKey:    os.Getenv("WorkflowApiKey"),
		User:      "workflow-demo",
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
		input["arg2"] = "test2"
		input["arg1"] = "test1"

		resp, err := client.WorkflowApp().RunBlock(ctx, types.WorkflowRequest{
			Inputs:       input,
			ResponseMode: "",
			User:         "",
		})
		t.Logf("resp=%+v err=%v", resp, err)
	})

	t.Run("Workflow_Run_ParseToStructCh", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["arg2"] = "test2"
		input["arg1"] = "test1"

		eventCh := client.WorkflowApp().Run(ctx, types.WorkflowRequest{
			Inputs:       input,
			ResponseMode: "",
			User:         "",
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

	t.Run("Workflow_Run_SimplePrint", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["arg2"] = "test2"
		input["arg1"] = "test1"

		eventCh := client.DebugOff().WorkflowApp().Run(ctx, types.WorkflowRequest{
			Inputs:       input,
			ResponseMode: "",
			User:         "",
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

	t.Run("Workflow_Run_ParseToEventCh", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["arg2"] = "test2"
		input["arg1"] = "test1"

		eventCh := client.WorkflowApp().Run(ctx, types.WorkflowRequest{
			Inputs:       input,
			ResponseMode: "",
			User:         "",
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

	t.Run("Workflow_Run_Stop", func(t *testing.T) {
		ctx := context.Background()

		input := make(map[string]interface{})
		input["arg2"] = "test2"
		input["arg1"] = "test1"

		eventCh := client.WorkflowApp().Run(ctx, types.WorkflowRequest{
			Inputs:       input,
			ResponseMode: "",
			User:         "",
		}).ParseToStructCh()
		cnt := 0
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
				cnt++
				if cnt == 1 {
					err := client.WorkflowApp().Stop(msg.TaskId, "")
					t.Logf("err=%v", err)
				}
			}
		}
	})

	t.Run("Workflow_UploadFile", func(t *testing.T) {
		f, err := os.Open("/Users/alsc/Downloads/abcd")
		if err != nil {
			t.Logf("Error opening file: %v\n", err)
			return
		}
		defer f.Close()

		resp, err := client.WorkflowApp().UploadFile(
			"/Users/alsc/Downloads/会议室分布.png",
			nil,
			"",
		)
		t.Logf("resp=%v err=%v\n", resp, err)

		resp, err = client.WorkflowApp().UploadFile(
			"",
			f,
			"",
		)
		t.Logf("resp=%v err=%v\n", resp, err)
	})

	t.Run("Workflow_AppInfo", func(t *testing.T) {
		resp, err := client.WorkflowApp().AppInfo()
		t.Logf("resp=%+v err=%v", resp, err)
	})

	t.Run("Workflow_AppParameter", func(t *testing.T) {
		resp, err := client.WorkflowApp().AppParameter()
		t.Logf("resp=%+v err=%v", resp, err)
	})

	t.Run("Workflow_Status", func(t *testing.T) {
		resp, err := client.WorkflowApp().Status("efe3357d-fb0d-480e-974f-50acd137dfd6")
		t.Logf("resp=%v err=%v\n", resp, err)
	})

	t.Run("Workflow_Logs", func(t *testing.T) {
		resp, err := client.WorkflowApp().Logs("", types.StatusStopped, 0, 0)
		t.Logf("resp=%v err=%v\n", resp, err)
	})
}
