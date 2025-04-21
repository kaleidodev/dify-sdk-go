package types

type FileInfo struct {
	Id        string `json:"id"` // 文件ID
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Extension string `json:"extension"`
	MimeType  string `json:"mime_type"`
	CreatedBy string `json:"created_by"`
	CreatedAt int64  `json:"created_at"`
}

type AppInfo struct {
	Name        string   `json:"name"`        // 应用名称
	Description string   `json:"description"` // 应用描述
	Tags        []string `json:"tags"`        // 应用标签
}

type AppParameter struct {
	OpeningStatement              string   `json:"opening_statement"`   //开场白
	SuggestedQuestions            []string `json:"suggested_questions"` // 开场推荐问题列表
	SuggestedQuestionsAfterAnswer struct {
		Enabled bool `json:"enabled"` // 是否开启
	} `json:"suggested_questions_after_answer"` // 启用回答后给出推荐问题
	SpeechToText struct {
		Enabled bool `json:"enabled"`
	} `json:"speech_to_text"` // 语音转文本
	RetrieverResource struct {
		Enabled bool `json:"enabled"`
	} `json:"retriever_resource"` // 引用和归属
	AnnotationReply struct {
		Enabled bool `json:"enabled"`
	} `json:"annotation_reply"` // 标记回复
	UserInputForm []struct {
		TextInput struct {
			Label     string `json:"label"`      // 控件展示标签名
			Variable  string `json:"variable"`   // 控件 ID
			Required  bool   `json:"required"`   // 是否必填
			Default   string `json:"default"`    // 默认值
			MaxLength int    `json:"max_length"` // 最大长度
		} `json:"text-input,omitempty"` // 文本输入控件
		Paragraph struct {
			Label     string `json:"label"`      // 控件展示标签名
			Variable  string `json:"variable"`   // 控件 ID
			Required  bool   `json:"required"`   // 是否必填
			Default   string `json:"default"`    // 默认值
			MaxLength int    `json:"max_length"` // 最大长度
		} `json:"paragraph,omitempty"` // 段落文本输入控件
		Select struct {
			Label    string   `json:"label"`    // 控件展示标签名
			Variable string   `json:"variable"` // 控件 ID
			Required bool     `json:"required"` // 是否必填
			Default  string   `json:"default"`  // 默认值
			Options  []string `json:"options"`  // 选项值
		} `json:"select,omitempty"` // 下拉控件
		Number struct {
			Label    string `json:"label"`    // 控件展示标签名
			Variable string `json:"variable"` // 控件 ID
			Required bool   `json:"required"` // 是否必填
			Default  string `json:"default"`  // 默认值
		} `json:"number,omitempty"` // 数字输入控件
	} `json:"user_input_form"` // 用户输入表单配置
	FileUpload struct {
		AllowedFileExtensions    []string `json:"allowed_file_extensions"`
		AllowedFileTypes         []string `json:"allowed_file_types"`
		AllowedFileUploadMethods []string `json:"allowed_file_upload_methods"`
		Enabled                  bool     `json:"enabled"`
		Image                    struct {
			Detail          string   `json:"detail"`
			Enabled         bool     `json:"enabled"`
			NumberLimits    int      `json:"number_limits"`
			TransferMethods []string `json:"transfer_methods"`
		} `json:"image"`
		NumberLimits int `json:"number_limits"`
	} `json:"file_upload"` // 文件上传配置
	SystemParameters struct {
		FileSizeLimit           int `json:"file_size_limit"`       // 文档上传大小限制 (MB)
		ImageFileSizeLimit      int `json:"image_file_size_limit"` // 图片文件上传大小限制（MB）
		AudioFileSizeLimit      int `json:"audio_file_size_limit"` // 音频文件上传大小限制 (MB)
		VideoFileSizeLimit      int `json:"video_file_size_limit"` // 视频文件上传大小限制 (MB)
		WorkflowFileUploadLimit int `json:"workflow_file_upload_limit"`
	} `json:"system_parameters"`
	MoreLikeThis struct {
		Enabled bool `json:"enabled"`
	} `json:"more_like_this"`
	SensitiveWordAvoidance struct {
		Configs []interface{} `json:"configs"`
		Enabled bool          `json:"enabled"`
		Type    string        `json:"type"`
	} `json:"sensitive_word_avoidance"`
	TextToSpeech struct {
		Enabled  bool   `json:"enabled"`
		Language string `json:"language"`
		Voice    string `json:"voice"`
	} `json:"text_to_speech"`
}

type FeedbackReq struct {
	MessageId string   `json:"message_id"` // 消息 ID
	Rating    Feedback `json:"rating"`     // 点赞 like, 点踩 dislike, 撤销点赞 null
	User      string   `json:"user"`       // 用户标识
	Content   string   `json:"content"`    // 消息反馈的具体信息
}

type ConversationRenameReq struct {
	ConversationId string `json:"conversation_id"` // 会话 ID
	Name           string `json:"name"`            // 选填）名称，若 auto_generate 为 true 时，该参数可不传
	AutoGenerate   bool   `json:"auto_generate"`   // 选填）自动生成标题，默认 false
	User           string `json:"user"`            // 用户标识
}

type ConversationRenameResp struct {
	Id           string      `json:"id"`           // 会话 ID
	Name         string      `json:"name"`         // 会话名称
	Inputs       interface{} `json:"inputs"`       // 用户输入参数
	Status       string      `json:"status"`       // 会话状态
	Introduction string      `json:"introduction"` // 开场白
	CreatedAt    int         `json:"created_at"`   // 创建时间
	UpdatedAt    int         `json:"updated_at"`   // 更新时间
}

type Text2Audio struct {
	MessageId string `json:"message_id"` // 消息ID
	Text      string `json:"text"`       // 语音生成内容 当MessageId非空时,使用MessageId的文本内容
	User      string `json:"user"`       // 用户标识
}

type AppMeta struct {
	ToolIcons map[string]interface{} `json:"tool_icons"`
}

type Conversation struct {
	Id           string      `json:"id"`                   // 会话 ID
	Name         string      `json:"name,omitempty"`       // 会话名称
	Inputs       interface{} `json:"inputs,omitempty"`     // 用户输入参数
	Status       string      `json:"status,omitempty"`     // 会话状态
	Introduction string      `json:"introduction"`         // 开场白
	CreatedAt    int         `json:"created_at,omitempty"` // 创建时间
	UpdatedAt    int         `json:"updated_at,omitempty"` // 更新时间
}
type ConversationListResp struct {
	Data    []Conversation `json:"data"`
	HasMore bool           `json:"has_more"`
	Limit   int            `json:"limit"`
}

type MessageHistory struct {
	Data []struct {
		Id             string      `json:"id"`              // 消息 ID
		ConversationId string      `json:"conversation_id"` // 会话 ID
		Inputs         interface{} `json:"inputs"`          // 用户输入参数
		Query          string      `json:"query"`           // 用户输入 / 提问内容
		MessageFiles   []struct {
			Id        string `json:"id"`
			Type      string `json:"type"`       // 文件类型
			Url       string `json:"url"`        // 预览图片地址
			BelongsTo string `json:"belongs_to"` // 文件归属方，user 或 assistant
		} `json:"message_files"` // 消息文件
		AgentThoughts []struct {
			Id          string        `json:"id"`
			MessageId   string        `json:"message_id"`
			Position    int           `json:"position"`
			Thought     string        `json:"thought"` // agent的思考内容
			Observation string        `json:"observation"`
			Tool        string        `json:"tool"`
			ToolInput   string        `json:"tool_input"`
			CreatedAt   int           `json:"created_at"`
			ChainId     interface{}   `json:"chain_id"`
			Files       []interface{} `json:"files"`
			ToolLabels  interface{}   `json:"tool_labels"`
		} `json:"agent_thoughts,omitempty"` //Agent思考内容 仅Agent类型有该内容
		Answer             string      `json:"answer"` // 回答消息内容
		CreatedAt          int         `json:"created_at"`
		Feedback           interface{} `json:"feedback"` //  反馈信息
		RetrieverResources []struct {
			Position     int     `json:"position"`
			DatasetId    string  `json:"dataset_id"`
			DatasetName  string  `json:"dataset_name"`
			DocumentId   string  `json:"document_id"`
			DocumentName string  `json:"document_name"`
			SegmentId    string  `json:"segment_id"`
			Score        float64 `json:"score"`
			Content      string  `json:"content"`
		} `json:"retriever_resources"` // 引用和归属分段列表
		Error           interface{} `json:"error"`
		ParentMessageId string      `json:"parent_message_id"`
		Status          string      `json:"status"`
	} `json:"data"`
	Limit   int  `json:"limit"`    // 返回条数
	HasMore bool `json:"has_more"` // 是否存在下一页
}

type File struct {
	Type           string `json:"type"`                     // 支持类型 image-图片 document-文档 audio-音频 video-视频 custom-其它
	TransferMethod string `json:"transfer_method"`          // 传递方式 remote_url-图片地址 local_file-上传文件
	Url            string `json:"url,omitempty"`            // 图片地址
	UploadFileId   string `json:"upload_file_id,omitempty"` // 上传文件ID
}

type ChatRequest struct {
	Query            string                 `json:"query"`                        // 用户输入/提问内容
	Inputs           map[string]interface{} `json:"inputs"`                       // 允许传入App定义的各变量值，默认{}
	ResponseMode     string                 `json:"response_mode,omitempty"`      // streaming-流式模式 blocking-阻塞模式(Agent模式不支持)
	User             string                 `json:"user"`                         // 用户标识
	ConversationId   string                 `json:"conversation_id,omitempty"`    // [选填]之前的会话ID，可基于之前聊天记录继续对话
	Files            []File                 `json:"files,omitempty"`              // 上传的文件
	AutoGenerateName *bool                  `json:"auto_generate_name,omitempty"` // [选填]自动生成标题，默认true
}

type CompletionRequest struct {
	Query        string                 `json:"query"`                   // 用户输入/提问内容
	Inputs       map[string]interface{} `json:"inputs"`                  // 允许传入App定义的各变量值，默认{}
	ResponseMode string                 `json:"response_mode,omitempty"` // streaming-流式模式 blocking-阻塞模式(Agent模式不支持)
	User         string                 `json:"user"`                    // 用户标识
	Files        []File                 `json:"files,omitempty"`         // 上传的文件
}

type WorkflowRequest struct {
	Inputs       map[string]interface{} `json:"inputs"`                  // 允许传入App定义的各变量值，默认{}
	ResponseMode string                 `json:"response_mode,omitempty"` // streaming-流式模式 blocking-阻塞模式(Agent模式不支持)
	User         string                 `json:"user"`                    // 用户标识
}

type Usage struct {
	PromptTokens        int     `json:"prompt_tokens"`
	PromptUnitPrice     string  `json:"prompt_unit_price"`
	PromptPriceUnit     string  `json:"prompt_price_unit"`
	PromptPrice         string  `json:"prompt_price"`
	CompletionTokens    int     `json:"completion_tokens"`
	CompletionUnitPrice string  `json:"completion_unit_price"`
	CompletionPriceUnit string  `json:"completion_price_unit"`
	CompletionPrice     string  `json:"completion_price"`
	TotalTokens         int     `json:"total_tokens"`
	TotalPrice          string  `json:"total_price"`
	Currency            string  `json:"currency"`
	Latency             float64 `json:"latency"`
}

type RetrieverResource struct {
	Position     int     `json:"position"`
	DatasetId    string  `json:"dataset_id"`
	DatasetName  string  `json:"dataset_name"`
	DocumentId   string  `json:"document_id"`
	DocumentName string  `json:"document_name"`
	SegmentId    string  `json:"segment_id"`
	Score        float64 `json:"score"`
	Content      string  `json:"content"`
}

type Metadata struct {
	Usage              Usage               `json:"usage"`               // 模型用量信息
	RetrieverResources []RetrieverResource `json:"retriever_resources"` // 引用和归属分段列表
}

type ChunkChatCompletionResponse struct {
	// 事件类型 message/agent_message/agent_thought/message_file/message_end/tts_message/tts_message_end/message_replace/error/ping
	Event                string   `json:"event"`                     // 事件类型
	TaskId               string   `json:"task_id,omitempty"`         // 任务 ID
	MessageId            string   `json:"message_id,omitempty"`      // 消息唯一 ID
	ConversationId       string   `json:"conversation_id,omitempty"` // 会话 ID
	Answer               string   `json:"answer,omitempty"`          // LLM 返回文本块内容
	CreatedAt            int      `json:"created_at,omitempty"`      // 创建时间戳
	Id                   string   `json:"id,omitempty"`
	Position             int64    `json:"position,omitempty"`      //agent_thought在消息中的位置
	Thought              string   `json:"thought,omitempty"`       // agent的思考内容
	Observation          string   `json:"observation,omitempty"`   // 工具调用的返回结果
	Tool                 string   `json:"tool,omitempty"`          // 使用的工具列表
	ToolInput            string   `json:"tool_input,omitempty"`    // 工具的输入，JSON格式的字符串
	MessageFiles         []string `json:"message_files,omitempty"` // 当前 agent_thought 关联的文件ID
	Type                 string   `json:"type,omitempty"`          // 文件类型，目前仅为image
	BelongsTo            string   `json:"belongs_to,omitempty"`    // 文件归属
	Url                  string   `json:"url,omitempty"`           // 文件访问地址
	Metadata             Metadata `json:"metadata,omitempty"`      // 元数据
	Audio                string   `json:"audio,omitempty"`         // 语音合成之后的音频块使用 Base64 编码之后的文本内容
	FromVariableSelector []string `json:"from_variable_selector,omitempty"`
	WorkflowRunId        string   `json:"workflow_run_id,omitempty"` // workflow 执行 ID
	Data                 struct {
		Id                string                 `json:"id,omitempty"`                  // workflow 执行 ID
		WorkflowId        string                 `json:"workflow_id,omitempty"`         // 关联 Workflow ID
		SequenceNumber    int                    `json:"sequence_number,omitempty"`     // 自增序号，App 内自增，从 1 开始
		CreatedAt         int                    `json:"created_at,omitempty"`          // 开始时间
		NodeId            string                 `json:"node_id,omitempty"`             // 节点 ID
		NodeType          string                 `json:"node_type,omitempty"`           // 节点类型
		Title             string                 `json:"title,omitempty"`               // 节点名称
		Index             int                    `json:"index,omitempty"`               // 执行序号，用于展示 Tracing Node 顺序
		PredecessorNodeId string                 `json:"predecessor_node_id,omitempty"` // 前置节点 ID，用于画布展示执行路径
		Inputs            map[string]interface{} `json:"inputs,omitempty"`              // 节点中所有使用到的前置节点变量内容
		Outputs           struct {
			Text  string `json:"text"`
			Usage struct {
				PromptTokens        int     `json:"prompt_tokens"`
				PromptUnitPrice     string  `json:"prompt_unit_price"`
				PromptPriceUnit     string  `json:"prompt_price_unit"`
				PromptPrice         string  `json:"prompt_price"`
				CompletionTokens    int     `json:"completion_tokens"`
				CompletionUnitPrice string  `json:"completion_unit_price"`
				CompletionPriceUnit string  `json:"completion_price_unit"`
				CompletionPrice     string  `json:"completion_price"`
				TotalTokens         int     `json:"total_tokens"`
				TotalPrice          string  `json:"total_price"`
				Currency            string  `json:"currency"`
				Latency             float64 `json:"latency"`
			} `json:"usage"`
			FinishReason string `json:"finish_reason"`
		} `json:"outputs,omitempty"` // Optional 输出内容
		Status            string      `json:"status,omitempty"`       // 执行状态 running / succeeded / failed / stopped
		Error             interface{} `json:"error,omitempty"`        // Optional 错误原因
		ElapsedTime       float64     `json:"elapsed_time,omitempty"` // Optional 耗时(s)
		TotalTokens       int         `json:"total_tokens,omitempty"` // Optional 总使用 tokens
		TotalSteps        int         `json:"total_steps,omitempty"`  // 总步数（冗余），默认 0
		FinishedAt        int         `json:"finished_at,omitempty"`  // 结束时间
		ExecutionMetadata struct {
			TotalTokens int    `json:"total_tokens"` // optional 总使用 tokens
			TotalPrice  string `json:"total_price"`  // optional 总费用
			Currency    string `json:"currency"`     // optional 货币
		} `json:"execution_metadata,omitempty"`
		ProcessData struct {
			ModelMode string `json:"model_mode"`
			Prompts   []struct {
				Role  string        `json:"role"`
				Text  string        `json:"text"`
				Files []interface{} `json:"files"`
			} `json:"prompts"`
			ModelProvider string `json:"model_provider"`
			ModelName     string `json:"model_name"`
		} `json:"process_data,omitempty"`
		Files                     []interface{} `json:"files,omitempty"`
		ParallelId                interface{}   `json:"parallel_id,omitempty"`
		ParallelStartNodeId       interface{}   `json:"parallel_start_node_id,omitempty"`
		ParentParallelId          interface{}   `json:"parent_parallel_id,omitempty"`
		ParentParallelStartNodeId interface{}   `json:"parent_parallel_start_node_id,omitempty"`
		IterationId               interface{}   `json:"iteration_id,omitempty"`
		LoopId                    interface{}   `json:"loop_id,omitempty"`
	} `json:"data,omitempty"`
	Status  int    `json:"status,omitempty"`  // HTTP 状态码
	Code    string `json:"code,omitempty"`    // 错误码
	Message string `json:"message,omitempty"` // 错误消息
}

type ChatCompletionResponse struct {
	Id             string   `json:"id"`                        // 同MessageId
	TaskId         string   `json:"task_id"`                   // 任务id
	MessageId      string   `json:"message_id"`                // 消息唯一ID
	ConversationId string   `json:"conversation_id,omitempty"` // 会话ID  Completion应用无该字段
	Mode           string   `json:"mode"`                      // App 模式，固定为 chat；Completion 应用固定为 completion;Chatflow 应用固定为 advanced-chat
	Answer         string   `json:"answer"`                    // 完整回复内容
	Event          string   `json:"event"`                     // 固定为 message
	Metadata       Metadata `json:"metadata"`                  // 元数据
	CreatedAt      int      `json:"created_at"`                // 消息创建时间戳
}

type WorkflowResponse struct {
	WorkflowRunId string `json:"workflow_run_id"` // workflow 执行 ID
	TaskId        string `json:"task_id"`         // 任务 ID
	Data          struct {
		Id          string                 `json:"id"`           // workflow 执行 ID
		WorkflowId  string                 `json:"workflow_id"`  // 关联 Workflow ID
		Status      string                 `json:"status"`       // 执行状态, running / succeeded / failed / stopped
		Outputs     map[string]interface{} `json:"outputs"`      // Optional 输出内容 json
		Error       string                 `json:"error"`        // Optional 错误原因
		ElapsedTime float64                `json:"elapsed_time"` // Optional 耗时(s)
		TotalTokens int                    `json:"total_tokens"` // Optional 总使用 tokens
		TotalSteps  int                    `json:"total_steps"`  // 总步数（冗余），默认 0
		CreatedAt   int                    `json:"created_at"`   // 开始时间
		FinishedAt  int                    `json:"finished_at"`  // 结束时间
	} `json:"data"`
}

type WorkflowStatus struct {
	Id          string  `json:"id"`           // workflow 执行 ID
	WorkflowId  string  `json:"workflow_id"`  // 关联的 Workflow ID
	Status      string  `json:"status"`       // 执行状态 running / succeeded / failed / stopped
	Inputs      string  `json:"inputs"`       // 任务输入内容 json字符串
	Outputs     string  `json:"outputs"`      // 任务输出内容 json字符串
	Error       string  `json:"error"`        // 错误原因
	TotalSteps  int     `json:"total_steps"`  // 任务执行总步数
	TotalTokens int     `json:"total_tokens"` // 任务执行总 tokens
	CreatedAt   string  `json:"created_at"`   // 任务开始时间
	FinishedAt  string  `json:"finished_at"`  // 任务结束时间
	ElapsedTime float64 `json:"elapsed_time"` // 耗时(s)
}

type WorkflowLogs struct {
	Page    int  `json:"page"`     // 当前页码
	Limit   int  `json:"limit"`    // 每页条数
	Total   int  `json:"total"`    // 总条数
	HasMore bool `json:"has_more"` // 是否还有更多数据
	Data    []struct {
		Id          string `json:"id"` // 标识
		WorkflowRun struct {
			Id          string  `json:"id"`              // 标识
			Version     string  `json:"version"`         // 版本
			Status      string  `json:"status"`          // 执行状态, running / succeeded / failed / stopped
			Error       string  `json:"error,omitempty"` // 错误
			ElapsedTime float64 `json:"elapsed_time"`    // 耗时，单位秒
			TotalTokens int     `json:"total_tokens"`    // 消耗的token数量
			TotalSteps  int     `json:"total_steps"`     // 执行步骤长度
			CreatedAt   int     `json:"created_at"`      // 开始时间
			FinishedAt  int     `json:"finished_at"`     // 结束时间
		} `json:"workflow_run"` // Workflow 执行日志
		CreatedFrom      string `json:"created_from"`                 // 来源
		CreatedByRole    string `json:"created_by_role"`              // 角色
		CreatedByAccount string `json:"created_by_account,omitempty"` // 帐号
		CreatedByEndUser struct {
			Id          string `json:"id"`           // 标识
			Type        string `json:"type"`         // 类型
			IsAnonymous bool   `json:"is_anonymous"` // 是否匿名
			SessionId   string `json:"session_id"`   // 会话标识
		} `json:"created_by_end_user"` // 用户
		CreatedAt int `json:"created_at"` // 创建时间
	} `json:"data"`
}
