package entity

type HTTPError struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}

func (h *HTTPError) Error() string {
	return h.Message
}
