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

// Usage
// func someFunction() error {
// 	err := doSomething()
// 	if err != nil {
// 		return &HTTPError{
// 			Status:     http.StatusText(http.StatusBadRequest),
// 			StatusCode: http.StatusBadRequest,
// 			Message:    err.Error(),
// 			Result:     nil,
// 		}
// 	}
// 	// ...
// 	return nil
// }
