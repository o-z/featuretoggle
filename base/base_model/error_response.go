package base_model

type ContextError struct {
	ErrorCode string `json:"ErrorCode,omitempty"`
	ErrorDesc string `json:"ErrorDesc,omitempty"`
}
