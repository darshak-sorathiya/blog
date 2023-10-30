package model

// SuccessResponseModel contains success response data
type SuccessResponseModel struct {
	Msg    string      `json:"message"` // Error message
	Status int         `json:"status"`  // Http status code
	Data   interface{} `json:"data"`    // data
}

// SuccessResponse takes necessary details in args and returns SuccessResponseModel structure of data
func SuccessResponse(status int, msg string, data interface{}) SuccessResponseModel {
	return SuccessResponseModel{msg, status, data}
}
