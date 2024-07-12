package utils

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(message string, data interface{}) response {
	response := response{
		Status:  "success",
		Message: message,
	}

	if data != nil {
		response.Data = data
	}

	return response
}

func ResponseError(message string) response {
	response := response{
		Status:  "fail",
		Message: message,
	}

	return response
}
