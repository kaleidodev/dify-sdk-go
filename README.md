# dify-sdk-go

dify SDK的go语言版本(dify版本 >= 1.1.3,低版本未测试)
使用go官方自带库构建，无任何三方库依赖
囊括dify应用的所有接口

### Dify应用类型

Dify一共有5种应用类型，具体如下：


| 应用类型             | 名称     | 类型       |
| -------------------- | -------- | ---------- |
| 对话型应用           | 聊天助手 | Chatbot    |
| 对话型应用           | Agent    | Agent      |
| 文本生成型应用       | 文本生成 | Completion |
| 工作流编排对话型应用 | Chatflow | Chatflow   |
| Workflow应用         | 工作流   | Workflow   |

### 应用接口与SDK函数对应关系

每种应用类型拥有的API接口并不完全相同，下表列出了Dify每种应用的Api接口和在SDK中对应的函数关系。


| 功能                         | SDK函数                              | Chatbot/Agent | Completion | Chatflow | Workflow | Dify接口                                             |
| ---------------------------- | ------------------------------------ | ------------- | ---------- | -------- | -------- | ---------------------------------------------------- |
| 发送对话消息                 | Run/RunBlock                         | 1             |            | 1        |          | POST`/chat-messages`                                 |
| 发送消息                     | Run/RunBlock                         |               | 1          |          |          | POST`/completion-messages`                           |
| 执行workflow                 | Run/RunBlock                         |               |            |          | 1        | POST`/workflows/run`                                 |
| --                           |                                      |               |            |          |          |                                                      |
| 停止响应                     | Stop                                 | 1             |            | 1        |          | POST`/chat-messages/:task_id/stop`                   |
| 停止响应                     | Stop                                 |               | 1          |          |          | POST`/completion-messages/:task_id/stop`             |
| 停止响应                     | Stop                                 |               |            |          | 1        | POST`/workflows/tasks/:task_id/stop`                 |
| --                           |                                      |               |            |          |          |                                                      |
| 上传文件                     | UploadFile                           | 1             | 1          | 1        | 1        | POST`/files/upload`                                  |
| 获取应用基本信息             | AppInfo                              | 1             | 1          | 1        | 1        | GET`/info`                                           |
| 获取应用参数                 | AppParameter                         | 1             | 1          | 1        | 1        | GET`/parameters`                                     |
| 获取应用WebApp设置           | AppSite                              | 1             | 1          | 1        | 1        | GET`/site`                                           |
| --                           |                                      |               |            |          |          |                                                      |
| 获取workflow执行情况         | Status                               |               |            |          | 1        | GET`/workflows/run/:workflow_id`                     |
| 消息反馈(点赞)               | MsgFeedback                          | 1             | 1          | 1        |          | POST`/messages/:message_id/feedbacks`                |
| 获取APP的消息点赞和反馈      | AppFeedback                          | 1             | 1          | 1        |          | GET`/app/feedbacks`                                  |
| 获取下一轮建议问题列表       | SuggestQuestionList                  | 1             |            | 1        |          | GET`/messages/{message_id}/suggested`                |
| 获取会话历史消息             | History/HistoryPro                   | 1             |            | 1        |          | GET`/messages`                                       |
| 获取workflow日志             | Logs                                 |               |            |          | 1        | GET`/workflows/logs`                                 |
| 获取会话列表                 | ConversationList/ConversationListPro | 1             |            | 1        |          | GET`/conversations`                                  |
| 删除会话                     | ConversationDel                      | 1             |            | 1        |          | DELETE`/conversations/:conversation_id`              |
| 会话重命名                   | ConversationRename                   | 1             |            | 1        |          | POST`/conversations/:conversation_id/name`           |
| 获取对话变量                 | ConversationVars                     | 1             |            | 1        |          | GET`/conversations/:conversation_id/variables`       |
| 语音转文字                   | AudioToText                          | 1             |            | 1        |          | POST`/audio-to-text`                                 |
| 文字转语音                   | TextToAudio                          | 1             | 1          | 1        |          | POST`/text-to-audio`                                 |
| 获取应用Meta信息             | AppMeta                              | 1             |            | 1        |          | GET`/meta`                                           |
| 获取标注列表                 | AnnotationList                       |               | 1          | 1        |          | GET`/apps/annotations`                               |
| 创建标注                     | AnnotationCreate                     |               | 1          | 1        |          | POST`/apps/annotations`                              |
| 更新标注                     | AnnotationUpdate                     |               | 1          | 1        |          | PUT`/apps/annotations/{annotation_id}`               |
| 删除标注                     | AnnotationDel                        |               | 1          | 1        |          | DELETE`/apps/annotations/{annotation_id}`            |
| 标注回复初始设置             | AnnotationReplySetting               |               | 1          | 1        |          | POST`/apps/annotation-reply/{action}`                |
| 查询标注回复初始设置任务状态 | AnnotationReplySettingJobStatus      |               | 1          | 1        |          | GET`/apps/annotation-reply/{action}/status/{job_id}` |
| --                           |                                      |               |            |          |          |                                                      |

### 创建Client

Client的配置定义在dify.ClientConfig这个结构体中：

```golang
type ClientConfig struct {
	ApiServer string        // [必填]API服务器 eg: http://your.domain.com/v1 注意是包括/v1的
	ApiKey    string        // [必填]API密钥
	User      string        // 用户标识 部分接口需要传入用户标识,这里设置后,调用接口时user字段可留空
	Debug     bool          // 是否打印原始请求及响应
	Timeout   time.Duration // 超时时间,默认300秒
	Transport *http.Transport
}
NewClient(config ClientConfig) (*base.Client, error) 
```

ClientConfig有两个必填参数ApiServer和ApiKey，由于很多接口都需要传入User参数，所以建议在创建客户端时同时把User的值也设置了，这样后面使用时，遇到User参数的地方可以传入空字符串。
其他参数可以根据需要进行设置，所以我们可以这样构建一个客户端：

```golang
        client,err:=dify.NewClient(dify.ClientConfig{
		ApiServer: "http://your.domain/v1", // 注意这里是包括/v1的
		ApiKey:    "your-api-key",
		User: "demo-client",
	})
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
```

客户端创建后，根据你应用的类型，调用XxxApp函数，然后就可以调用应用拥有的具体功能函数了

- 如果是 聊天助手 Chatbot 类型，则是 client.ChatbotApp()
- 如果是 Agent 类型，则是 client.AgentApp()
- 如果是 文本生成 Completion 类型，则是 client.CompletionApp()
- 如果是 Chatflow 类型，则是 client.ChatflowApp()
- 如果是 工作流 Workflow 类型，则是 client.WorkflowApp()

需要注意的是，对于流式调用，这里提供了三种结果的输出方式:

- 方式一：将SSE Event事件解析为一个大而全的结构体，再通过channel输出
  优点是输出是固定的结构体，但这个结构体字段很多，很多字段会是空值，使用时不方便判断哪些字段是有用的，哪些是没用的
  调用示例：`eventCh := client.AgentApp().Run(ctx, types.ChatRequest{}).ParseToStructCh()`
- 方式二：将SSE Event事件中的输出以文本字符串方式，通过channel提供出来
  优点是只输出最终文本内容，其他内容不输出，使用最简单
  调用示例：`eventCh := client.AgentApp().Run(ctx, types.ChatRequest{}).SimplePrint()`
- 方式三：将SSE Event按事件类型，解析为具体的结构体，然后通过channel提供
  优点是不同的event事件类型，对应不同的结构体，更加精准，但是使用前需要做类型断言 如：`msg.Data.(*types.EventMessage)`
  调用示例：`eventCh := client.AgentApp().Run(ctx, types.ChatRequest{}).ParseToEventCh()`

### 一个完整的示例

阻塞式调用示例：

```golang
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
```

流式调用示例：

```golang
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
		User:   "demo-client",
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

```

### 待完善的功能

1、文字转语音接口(没有模型供调试)
2、语音转文字接口
