### Dify 应用类型


| 应用类型             | 名称     | 类型       | 核心接口                |
| -------------------- | -------- | ---------- | ----------------------- |
| 对话型应用           | 聊天助手 | Chatbot    | /v1/chat-messages       |
| 对话型应用           | Agent    | Agent      | /v1/chat-messages       |
| 文本生成型应用       | 文本生成 | Completion | /v1/completion-messages |
| 工作流编排对话型应用 | Chatflow | Chatflow   | /v1/chat-messages       |
| Workflow 应用        | 工作流   | Workflow   | /v1/workflows/run       |

### 各类型应用接口

#### 1、Chatbot/Agent

1. POST `/chat-messages` 发送对话消息
2. POST `/files/upload` 上传文件
3. POST `/chat-messages/:task_id/stop` 停止响应
4. POST `/messages/:message_id/feedbacks` 消息反馈(点赞)
5. GET `/app/feedbacks` 获取APP的消息点赞和反馈
6. GET `/messages/{message_id}/suggested` 获取下一轮建议问题列表
7. GET `/messages` 获取会话历史消息
8. GET `/conversations` 获取会话列表
9. DELETE `/conversations/:conversation_id` 删除会话
10. POST `/conversations/:conversation_id/name` 会话重命名
11. GET `/conversations/:conversation_id/variables` 获取对话变量
12. POST `/audio-to-text` 语音转文字
13. POST `/text-to-audio` 文字转语音
14. GET `/info` 获取应用基本信息
15. GET `/parameters` 获取应用参数
16. GET `/meta` 获取应用 Meta 信息
17. GET `/site` 获取应用 WebApp 设置

#### 2、Completion

1. POST `/completion-messages` 发送消息
2. POST `/files/upload` 上传文件
3. POST `/completion-messages/:task_id/stop` 停止响应
4. POST `/messages/:message_id/feedbacks` 消息反馈(点赞)
5. GET `/app/feedbacks` 获取APP的消息点赞和反馈
6. POST `/text-to-audio` 文字转语音
7. GET `/info` 获取应用基本信息
8. GET `/parameters` 获取应用参数
9. GET `/site` 获取应用 WebApp 设置
10. GET `/apps/annotations` 获取标注列表
11. POST `/apps/annotations` 创建标注
12. PUT `/apps/annotations/{annotation_id}` 更新标注
13. DELETE `/apps/annotations/{annotation_id}` 删除标注
14. POST `/apps/annotation-reply/{action}` 标注回复初始设置
15. GET `/apps/annotation-reply/{action}/status/{job_id}` 查询标注回复初始设置任务状态

#### 3、Chatflow

1. POST `/chat-messages` 发送对话消息
2. POST `/files/upload` 上传文件
3. POST `/chat-messages/:task_id/stop` 停止响应
4. POST `/messages/:message_id/feedbacks` 消息反馈(点赞)
5. GET `/app/feedbacks` 获取APP的消息点赞和反馈
6. GET `/messages/{message_id}/suggested` 获取下一轮建议问题列表
7. GET `/messages` 获取会话历史消息
8. GET `/conversations` 获取会话列表
9. DELETE `/conversations/:conversation_id` 删除会话
10. POST `/conversations/:conversation_id/name` 会话重命名
11. GET `/conversations/:conversation_id/variables` 获取对话变量
12. POST `/audio-to-text` 语音转文字
13. POST `/text-to-audio` 文字转语音
14. GET `/info` 获取应用基本信息
15. GET `/parameters` 获取应用参数
16. GET `/meta` 获取应用 Meta 信息
17. GET `/site` 获取应用 WebApp 设置
18. GET `/apps/annotations` 获取标注列表
19. POST `/apps/annotations` 创建标注
20. PUT `/apps/annotations/{annotation_id}` 更新标注
21. DELETE `/apps/annotations/{annotation_id}` 删除标注
22. POST `/apps/annotation-reply/{action}` 标注回复初始设置
23. GET `/apps/annotation-reply/{action}/status/{job_id}` 查询标注回复初始设置任务状态

#### 4、Workflow

1. POST `/workflows/run` 执行 workflow
2. GET `/workflows/run/:workflow_id` 获取 workflow 执行情况
3. POST `/workflows/tasks/:task_id/stop` 停止响应
4. POST `/files/upload` 上传文件
5. GET `/workflows/logs` 获取 workflow 日志
6. GET `/info` 获取应用基本信息
7. GET `/parameters` 获取应用参数
8. GET `/site` 获取应用 WebApp 设置

### 应用接口矩阵


| Dify 接口                                                                         | Chatbot/Agent | Completion | Chatflow | Workflow | SDK 对应函数                         |
| --------------------------------------------------------------------------------- | ------------- | ---------- | -------- | -------- | ------------------------------------ |
| POST`/chat-messages` 发送对话消息                                                 | 1             |            | 1        |          | Run/RunBlock                         |
| POST`/completion-messages` 发送消息                                               |               | 1          |          |          |                                      |
| POST`/workflows/run` 执行 workflow                                                |               |            |          | 1        |                                      |
|                                                                                   |               |            |          |          |                                      |
| POST`/chat-messages/:task_id/stop` 停止响应                                       | 1             |            | 1        |          | Stop                                 |
| POST`/completion-messages/:task_id/stop` 停止响应                                 |               | 1          |          |          |                                      |
| POST`/workflows/tasks/:task_id/stop` 停止响应                                     |               |            |          | 1        |                                      |
|                                                                                   |               |            |          |          |                                      |
| POST`/files/upload` 上传文件                                                      | 1             | 1          | 1        | 1        | UploadFile                           |
| GET`/info` 获取应用基本信息                                                       | 1             | 1          | 1        | 1        | AppInfo                              |
| GET`/parameters` 获取应用参数                                                     | 1             | 1          | 1        | 1        | AppParameter                         |
| GET`/site` 获取应用 WebApp 设置                                                   | 1             | 1          | 1        | 1        | AppSite                              |
|                                                                                   |               |            |          |          |                                      |
| GET`/workflows/run/:workflow_id` 获取 workflow 执行情况                           |               |            |          | 1        | Status                               |
| POST`/messages/:message_id/feedbacks` 消息反馈(点赞)                              | 1             | 1          | 1        |          | MsgFeedback                          |
| GET`/app/feedbacks` 获取APP的消息点赞和反馈                                       | 1             | 1          | 1        |          | AppFeedback                          |
| GET`/messages/{message_id}/suggested` 获取下一轮建议问题列表                      | 1             |            | 1        |          | SuggestQuestionList                  |
| GET`/messages` 获取会话历史消息                                                   | 1             |            | 1        |          | History/HistoryPro                   |
| GET`/workflows/logs` 获取 workflow 日志                                           |               |            |          | 1        | Logs                                 |
| GET`/conversations` 获取会话列表                                                  | 1             |            | 1        |          | ConversationList/ConversationListPro |
| DELETE`/conversations/:conversation_id` 删除会话                                  | 1             |            | 1        |          | ConversationDel                      |
| POST`/conversations/:conversation_id/name` 会话重命名                             | 1             |            | 1        |          | ConversationRename                   |
| GET`/conversations/:conversation_id/variables` 获取对话变量                       | 1             |            | 1        |          | ConversationVars                     |
| POST`/audio-to-text` 语音转文字                                                   | 1             |            | 1        |          | AudioToText                          |
| POST`/text-to-audio` 文字转语音                                                   | 1             | 1          | 1        |          | TextToAudio                          |
| GET`/meta` 获取应用 Meta 信息                                                     | 1             |            | 1        |          | AppMeta                              |
| GET`/apps/annotations` 获取标注列表                                               |               | 1          | 1        |          | AnnotationList                       |
| POST`/apps/annotations` 创建标注                                                  |               | 1          | 1        |          | AnnotationCreate                     |
| PUT`/apps/annotations/{annotation_id}` 更新标注                                   |               | 1          | 1        |          | AnnotationUpdate                     |
| DELETE`/apps/annotations/{annotation_id}` 删除标注                                |               | 1          | 1        |          | AnnotationDel                        |
| POST`/apps/annotation-reply/{action}` 标注回复初始设置                            |               | 1          | 1        |          | AnnotationReplySetting               |
| GET`/apps/annotation-reply/{action}/status/{job_id}` 查询标注回复初始设置任务状态 |               | 1          | 1        |          | AnnotationReplySettingJobStatus      |
|                                                                                   |               |            |          |          |                                      |

### dify 接口 api 文档

源码路径：dify/web/app/components/develop/template/ & dify/api/core/app/entities/task_entities.py
每个版本的放于本目录相头版本号文件夹内
