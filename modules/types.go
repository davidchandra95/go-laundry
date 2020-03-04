package modules

type APIResponse struct {
	Status            int         `json:"status"`
	Config            struct{}    `json:"config"`
	Data              interface{} `json:"data,omitempty"`
	ServerProcessTime string      `json:"server_process_time"`
	MessageError      []string    `json:"message_error,omitempty"`
}
