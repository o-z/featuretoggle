package base_model

type CustomResponse struct {
	ResponseTime  int64
	StatusCode    int
	Data          interface{}    `json:"Data,omitempty"`
	ContextErrors []ContextError `json:"ContextErrors,omitempty"`
}
