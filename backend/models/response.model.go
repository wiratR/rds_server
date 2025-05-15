package models

import "github.com/gofrs/uuid"

type ResponseSuccessToken struct {
	Status string `json:"status"`
	Data   struct {
		Token string `json:"token"`
	} `json:"data"`
}

type ResponseSuccessUser struct {
	Status string `json:"status"`
	Data   struct {
		AccountId uuid.UUID    `json:"account_id"`
		User      UserResponse `json:"user"`
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

type ResponseSuccessAccount struct {
	Status string `json:"status"`
	Data   struct {
		User         UserResponse    `json:"user"`
		Account      AccountResponse `json:"account"`
		TxnHistories []TxnHistory    `json:"txn_histories"`
	} `json:"data"`
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
