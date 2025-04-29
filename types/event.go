package types

type StreamResponse struct {
	Event  string `json:"event"`   // 事件类型
	TaskId string `json:"task_id"` // 任务ID，用于请求跟踪下方的停止响应接口
}
type ChatbotAppStreamResponse struct {
	ConversationId string `json:"conversation_id"` // 会话ID
	MessageId      string `json:"message_id"`      // 消息唯一ID
	CreatedAt      int64  `json:"created_at"`      // 创建时间戳
}

type CompletionAppStreamResponse struct {
	MessageId string `json:"message_id"` // 消息唯一ID
	CreatedAt int64  `json:"created_at"` // 创建时间戳
}

type WorkflowAppStreamResponse struct {
	WorkflowRunId string `json:"workflow_run_id"`
}

// EventError ERROR = "error"
type EventError struct {
	StreamResponse

	Status  int    `json:"status"`  // HTTP 状态码
	Code    string `json:"code"`    // 错误码
	Message string `json:"message"` // 错误消息
}

// EventMessage MESSAGE = "message"
type EventMessage struct {
	StreamResponse
	ChatbotAppStreamResponse

	Id                   string   `json:"id"`
	Answer               string   `json:"answer"`
	FromVariableSelector []string `json:"from_variable_selector,omitempty"`
}

// EventMessageEnd MESSAGE_END = "message_end"
type EventMessageEnd struct {
	StreamResponse
	ChatbotAppStreamResponse

	Id       string   `json:"id"`
	Metadata Metadata `json:"metadata"` // 元数据
	Files    *[]File  `json:"files,omitempty"`
}

// EventTtsMessage TTS_MESSAGE = "tts_message"
type EventTtsMessage struct {
	StreamResponse
	ChatbotAppStreamResponse

	Audio string `json:"audio"`
}

// EventTtsMessageEnd TTS_MESSAGE_END = "tts_message_end"
type EventTtsMessageEnd struct {
	StreamResponse
	ChatbotAppStreamResponse

	Audio string `json:"audio"`
}

// EventMessageFile MESSAGE_FILE = "message_file"
type EventMessageFile struct {
	StreamResponse

	Id        string `json:"id"`
	Type      string `json:"type"`
	BelongsTo string `json:"belongs_to"`
	Url       string `json:"url"`
}

// EventMessageReplace MESSAGE_REPLACE = "message_replace"
type EventMessageReplace struct {
	StreamResponse

	Answer string `json:"answer"`
	Reason string `json:"reason"`
}

// EventAgentThought AGENT_THOUGHT = "agent_thought"
type EventAgentThought struct {
	StreamResponse
	ChatbotAppStreamResponse

	Id           string                 `json:"id"`
	Position     int                    `json:"position"`              // agent_thought在消息中的位置，如第一轮迭代position为1
	Thought      string                 `json:"thought,omitempty"`     // agent的思考内容
	Observation  string                 `json:"observation,omitempty"` // 工具调用的返回结果
	Tool         string                 `json:"tool,omitempty"`        // 使用的工具列表，以 ; 分割多个工具
	ToolLabels   map[string]interface{} `json:"tool_labels,omitempty"`
	ToolInput    string                 `json:"tool_input,omitempty"`    // 工具的输入，JSON格式的字符串
	MessageFiles []string               `json:"message_files,omitempty"` // 当前 agent_thought 关联的文件ID
}

// EventAgentMessage AGENT_MESSAGE = "agent_message"
type EventAgentMessage struct {
	StreamResponse
	ChatbotAppStreamResponse

	Id     string `json:"id"`
	Answer string `json:"answer"`
}

// EventWorkflowStarted WORKFLOW_STARTED = "workflow_started"
type EventWorkflowStarted struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		Id             string                 `json:"id"`              // workflow 执行 ID
		WorkflowId     string                 `json:"workflow_id"`     // 关联 Workflow ID
		SequenceNumber int                    `json:"sequence_number"` // 自增序号，App 内自增，从 1 开始
		Inputs         map[string]interface{} `json:"inputs"`          // 节点中所有使用到的前置节点变量内容
		CreatedAt      int64                  `json:"created_at"`      // 开始时间
	} `json:"data"`
}

// EventWorkflowFinished WORKFLOW_FINISHED = "workflow_finished"
type EventWorkflowFinished struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		Id              string                 `json:"id"`
		WorkflowId      string                 `json:"workflow_id"`
		SequenceNumber  int                    `json:"sequence_number"`
		Status          string                 `json:"status"`
		Outputs         map[string]interface{} `json:"outputs,omitempty"`
		Error           string                 `json:"error,omitempty"`
		ElapsedTime     float64                `json:"elapsed_time"`
		TotalTokens     int                    `json:"total_tokens"`
		TotalSteps      int                    `json:"total_steps"`
		CreatedBy       map[string]interface{} `json:"created_by,omitempty"`
		CreatedAt       int64                  `json:"created_at"`
		FinishedAt      int64                  `json:"finished_at"`
		ExceptionsCount int                    `json:"exceptions_count,omitempty"`
		Files           []File                 `json:"files,omitempty"`
	} `json:"data"`
}

// EventNodeStarted NODE_STARTED = "node_started"
type EventNodeStarted struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		Id                        string                 `json:"id"`
		NodeId                    string                 `json:"node_id"`
		NodeType                  string                 `json:"node_type"`
		Title                     string                 `json:"title"`
		Index                     int                    `json:"index"`
		PredecessorNodeId         string                 `json:"predecessor_node_id,omitempty"`
		Inputs                    map[string]interface{} `json:"inputs,omitempty"`
		CreatedAt                 int64                  `json:"created_at"`
		Extras                    map[string]interface{} `json:"extras,omitempty"`
		ParallelId                string                 `json:"parallel_id,omitempty"`
		ParallelStartNodeId       string                 `json:"parallel_start_node_id,omitempty"`
		ParentParallelId          string                 `json:"parent_parallel_id,omitempty"`
		ParentParallelStartNodeId string                 `json:"parent_parallel_start_node_id,omitempty"`
		IterationId               string                 `json:"iteration_id,omitempty"`
		LoopId                    string                 `json:"loop_id,omitempty"`
		ParallelRunId             string                 `json:"parallel_run_id,omitempty"`
		AgentStrategy             struct {
			Name string `json:"name"`
			Icon string `json:"icon,omitempty"`
		} `json:"agent_strategy,omitempty"`
	} `json:"data"`
}

// EventNodeFinished NODE_FINISHED = "node_finished"
type EventNodeFinished struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		Id                        string                 `json:"id"`
		NodeId                    string                 `json:"node_id"`
		NodeType                  string                 `json:"node_type"`
		Title                     string                 `json:"title"`
		Index                     int                    `json:"index"`
		PredecessorNodeId         string                 `json:"predecessor_node_id,omitempty"`
		Inputs                    map[string]interface{} `json:"inputs,omitempty"`
		ProcessData               map[string]interface{} `json:"process_data,omitempty"`
		Outputs                   map[string]interface{} `json:"outputs,omitempty"`
		Status                    string                 `json:"status"`
		Error                     string                 `json:"error,omitempty"`
		ElapsedTime               float64                `json:"elapsed_time"`
		ExecutionMetadata         map[string]interface{} `json:"execution_metadata,omitempty"`
		CreatedAt                 int64                  `json:"created_at"`
		FinishedAt                int64                  `json:"finished_at"`
		Files                     []File                 `json:"files,omitempty"`
		ParallelId                string                 `json:"parallel_id,omitempty"`
		ParallelStartNodeId       string                 `json:"parallel_start_node_id,omitempty"`
		ParentParallelId          string                 `json:"parent_parallel_id,omitempty"`
		ParentParallelStartNodeId string                 `json:"parent_parallel_start_node_id,omitempty"`
		IterationId               string                 `json:"iteration_id,omitempty"`
		LoopId                    string                 `json:"loop_id,omitempty"`
	} `json:"data"`
}

// EventNodeRetry NODE_RETRY = "node_retry"
type EventNodeRetry struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		Id                        string                 `json:"id"`
		NodeId                    string                 `json:"node_id"`
		NodeType                  string                 `json:"node_type"`
		Title                     string                 `json:"title"`
		Index                     int                    `json:"index"`
		PredecessorNodeId         string                 `json:"predecessor_node_id,omitempty"`
		Inputs                    map[string]interface{} `json:"inputs,omitempty"`
		ProcessData               map[string]interface{} `json:"process_data,omitempty"`
		Outputs                   map[string]interface{} `json:"outputs,omitempty"`
		Status                    string                 `json:"status"`
		Error                     string                 `json:"error,omitempty"`
		ElapsedTime               float64                `json:"elapsed_time"`
		ExecutionMetadata         map[string]interface{} `json:"execution_metadata,omitempty"`
		CreatedAt                 int64                  `json:"created_at"`
		FinishedAt                int64                  `json:"finished_at"`
		Files                     []File                 `json:"files,omitempty"`
		ParallelId                string                 `json:"parallel_id,omitempty"`
		ParallelStartNodeId       string                 `json:"parallel_start_node_id,omitempty"`
		ParentParallelId          string                 `json:"parent_parallel_id,omitempty"`
		ParentParallelStartNodeId string                 `json:"parent_parallel_start_node_id,omitempty"`
		IterationId               string                 `json:"iteration_id,omitempty"`
		LoopId                    string                 `json:"loop_id,omitempty"`
		RetryIndex                int64                  `json:"retry_index"`
	} `json:"data"`
}

// EventParallelBranchStarted PARALLEL_BRANCH_STARTED = "parallel_branch_started"
type EventParallelBranchStarted struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		ParallelId                string `json:"parallel_id"`
		ParallelBranchId          string `json:"parallel_branch_id"`
		ParentParallelId          string `json:"parent_parallel_id,omitempty"`
		ParentParallelStartNodeId string `json:"parent_parallel_start_node_id,omitempty"`
		IterationId               string `json:"iteration_id,omitempty"`
		LoopId                    string `json:"loop_id,omitempty"`
		CreatedAt                 int64  `json:"created_at"`
	}
}

// EventParallelBranchFinished PARALLEL_BRANCH_FINISHED = "parallel_branch_finished"
type EventParallelBranchFinished struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		ParallelId                string `json:"parallel_id"`
		ParallelBranchId          string `json:"parallel_branch_id"`
		ParentParallelId          string `json:"parent_parallel_id,omitempty"`
		ParentParallelStartNodeId string `json:"parent_parallel_start_node_id,omitempty"`
		IterationId               string `json:"iteration_id,omitempty"`
		LoopId                    string `json:"loop_id,omitempty"`
		Status                    string `json:"status"`
		Error                     string `json:"error,omitempty"`
		CreatedAt                 int64  `json:"created_at"`
	} `json:"data"`
}

// EventIterationStarted ITERATION_STARTED = "iteration_started"
type EventIterationStarted struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		Id                  string                 `json:"id"`
		NodeId              string                 `json:"node_id"`
		NodeType            string                 `json:"node_type"`
		Title               string                 `json:"title"`
		CreatedAt           int64                  `json:"created_at"`
		Extras              map[string]interface{} `json:"extras,omitempty"`
		Metadata            Metadata               `json:"metadata,omitempty"`
		Inputs              map[string]interface{} `json:"inputs,omitempty"`
		ParallelId          string                 `json:"parallel_id,omitempty"`
		ParallelStartNodeId string                 `json:"parallel_start_node_id,omitempty"`
	} `json:"data"`
}

// EventIterationNext ITERATION_NEXT = "iteration_next"
type EventIterationNext struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		Id                  string                 `json:"id"`
		NodeId              string                 `json:"node_id"`
		NodeType            string                 `json:"node_type"`
		Title               string                 `json:"title"`
		Index               int                    `json:"index"`
		CreatedAt           int64                  `json:"created_at"`
		PreIterationOutput  interface{}            `json:"pre_iteration_output,omitempty"`
		Extras              map[string]interface{} `json:"extras,omitempty"`
		ParallelId          string                 `json:"parallel_id,omitempty"`
		ParallelStartNodeId string                 `json:"parallel_start_node_id,omitempty"`
		ParallelModeRunId   string                 `json:"parallel_mode_run_id,omitempty"`
		Duration            float64                `json:"duration,omitempty"`
	} `json:"data"`
}

// EventIterationCompleted ITERATION_COMPLETED = "iteration_completed"
type EventIterationCompleted struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		Id                  string                 `json:"id"`
		NodeId              string                 `json:"node_id"`
		NodeType            string                 `json:"node_type"`
		Title               string                 `json:"title"`
		Outputs             map[string]interface{} `json:"outputs,omitempty"`
		CreatedAt           int64                  `json:"created_at"`
		Extras              map[string]interface{} `json:"extras,omitempty"`
		Inputs              map[string]interface{} `json:"inputs,omitempty"`
		Status              string                 `json:"status"` //running succeeded failed exception retry
		Error               string                 `json:"error,omitempty"`
		ElapsedTime         float64                `json:"elapsed_time"`
		TotalTokens         int                    `json:"total_tokens"`
		ExecutionMetadata   map[string]interface{} `json:"execution_metadata,omitempty"`
		FinishedAt          int64                  `json:"finished_at"`
		Steps               int                    `json:"steps"`
		ParallelId          string                 `json:"parallel_id,omitempty"`
		ParallelStartNodeId string                 `json:"parallel_start_node_id,omitempty"`
	} `json:"data"`
}

// EventLoopStarted LOOP_STARTED = "loop_started"
type EventLoopStarted struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		Id                  string                 `json:"id"`
		NodeId              string                 `json:"node_id"`
		NodeType            string                 `json:"node_type"`
		Title               string                 `json:"title"`
		CreatedAt           int64                  `json:"created_at"`
		Extras              map[string]interface{} `json:"extras,omitempty"`
		Metadata            Metadata               `json:"metadata,omitempty"`
		Inputs              map[string]interface{} `json:"inputs,omitempty"`
		ParallelId          string                 `json:"parallel_id,omitempty"`
		ParallelStartNodeId string                 `json:"parallel_start_node_id,omitempty"`
	} `json:"data"`
}

// EventLoopNext LOOP_NEXT = "loop_next"
type EventLoopNext struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		Id                  string                 `json:"id"`
		NodeId              string                 `json:"node_id"`
		NodeType            string                 `json:"node_type"`
		Title               string                 `json:"title"`
		Index               int                    `json:"index"`
		CreatedAt           int64                  `json:"created_at"`
		PreLoopOutput       interface{}            `json:"pre_loop_output,omitempty"`
		Extras              map[string]interface{} `json:"extras,omitempty"`
		ParallelId          string                 `json:"parallel_id,omitempty"`
		ParallelStartNodeId string                 `json:"parallel_start_node_id,omitempty"`
		ParallelModeRunId   string                 `json:"parallel_mode_run_id,omitempty"`
		Duration            float64                `json:"duration,omitempty"`
	} `json:"data"`
}

// EventLoopCompleted LOOP_COMPLETED = "loop_completed"
type EventLoopCompleted struct {
	StreamResponse
	WorkflowAppStreamResponse // workflow_run_id

	Data struct {
		Id                  string                 `json:"id"`
		NodeId              string                 `json:"node_id"`
		NodeType            string                 `json:"node_type"`
		Title               string                 `json:"title"`
		Outputs             map[string]interface{} `json:"outputs,omitempty"`
		CreatedAt           int64                  `json:"created_at"`
		Extras              map[string]interface{} `json:"extras,omitempty"`
		Inputs              map[string]interface{} `json:"inputs,omitempty"`
		Status              string                 `json:"status"` //running succeeded failed exception retry
		Error               string                 `json:"error,omitempty"`
		ElapsedTime         float64                `json:"elapsed_time"`
		TotalTokens         int                    `json:"total_tokens"`
		ExecutionMetadata   map[string]interface{} `json:"execution_metadata,omitempty"`
		FinishedAt          int64                  `json:"finished_at"`
		Steps               int                    `json:"steps"`
		ParallelId          string                 `json:"parallel_id,omitempty"`
		ParallelStartNodeId string                 `json:"parallel_start_node_id,omitempty"`
	} `json:"data"`
}

// EventTextChunk TEXT_CHUNK = "text_chunk"
type EventTextChunk struct {
	StreamResponse

	Data struct {
		Text                 string   `json:"text"`
		FromVariableSelector []string `json:"from_variable_selector,omitempty"`
	} `json:"data"`
}

// EventTextReplace TEXT_REPLACE = "text_replace"
type EventTextReplace struct {
	StreamResponse

	Data struct {
		Text string `json:"text"`
	} `json:"data"`
}

// EventAgentLog AGENT_LOG = "agent_log"
type EventAgentLog struct {
	StreamResponse

	Data struct {
		NodeExecutionId string                 `json:"node_execution_id"`
		Id              string                 `json:"id"`
		Label           string                 `json:"label"`
		ParentId        string                 `json:"parent_id,omitempty"`
		Error           string                 `json:"error,omitempty"`
		Status          string                 `json:"status"`
		Data            map[string]interface{} `json:"data"`
		Metadata        Metadata               `json:"metadata,omitempty"`
		NodeId          string                 `json:"node_id"`
	} `json:"data"`
}
