package model

type ErrFields map[string]string

// ResponseError contains error related fields
type ResponseError struct {
	Msg    string      `json:"message"` // Error message
	Status int         `json:"status"`  // Http status code
	Data   interface{} `json:"data"`    // data
}

// Error returns a general response error
func Error(msg string, status int) ResponseError {
	return ResponseError{msg, status, nil}
}
