package types

type AnnotationListResp struct {
	Data    []Annotation `json:"data"`
	HasMore bool         `json:"has_more"`
	Limit   int          `json:"limit"`
	Page    int          `json:"page"`
	Total   int          `json:"total"`
}

type Annotation struct {
	Answer    string `json:"answer"`
	CreatedAt int    `json:"created_at"`
	HitCount  int    `json:"hit_count"`
	Id        string `json:"id"`
	Question  string `json:"question"`
}

type AnnotationSetting struct {
	EmbeddingProviderName string  `json:"embedding_provider_name"` // 嵌入模型提供商 对应的是provider字段
	EmbeddingModelName    string  `json:"embedding_model_name"`    // 嵌入模型 对应的是model字段
	ScoreThreshold        float64 `json:"score_threshold"`         // 相似度阈值，当相似度大于该阈值时，系统会自动回复，否则不回复
}

type AnnotationSettingJobResp struct {
	JobId     string `json:"job_id"`
	JobStatus string `json:"job_status"` // waiting
}

type AnnotationSettingJobStatusResp struct {
	JobId     string `json:"job_id"`
	JobStatus string `json:"job_status"` // 执行状态，如 completed
	ErrorMsg  string `json:"error_msg"`
}
