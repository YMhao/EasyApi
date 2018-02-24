package serv

// CallResponse 是api的应答格式
type CallResponse struct {
	HasError bool        `json:"hasError,omitempty"`
	Error    *APIError   `json:"error,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
