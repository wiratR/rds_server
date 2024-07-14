package models

type ResponseSuccessToken struct {
	Status string `json:"status"`
	Data   struct {
		Token string `json:"token"`
	} `json:"data"`
}

type ResponseSuccessUser struct {
	Status string `json:"status"`
	Data   struct {
		User UserResponse `json:"user"`
	} `json:"data"`
}

type ResponseError struct {
	Status string `json:"status"`
	Data   struct {
		Message string `json:"message"`
	} `json:"data"`
}

type ResponseBody struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ApiResponse(isSuccess bool, data interface{}) ResponseBody {
	status := ""
	if isSuccess {
		status = "success"
	} else {
		status = "failed"
	}
	return ResponseBody{
		Status: status,
		Data:   data,
	}
}
