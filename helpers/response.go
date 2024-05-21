package helpers

import "gin_social/dto"

type ResponseWithData struct{
	Code int `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
	Paginate *dto.Paginate `json:"paginate, omitempty"`
	Data any `json:"data"`
}

type ResponseWithoutData struct{
	Code int `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
}

func Response(params *dto.ResponseParam) any {
	var response any
	var status string

	if params.StatusCode >= 200 && params.StatusCode <= 299 {
		status = "success"
	} else {
		status = "failed"
	}

	if response != nil {
		response = &ResponseWithData{
			Code: params.StatusCode,
			Status: status,
			Message: params.Message,
			Paginate: params.Paginate,
			Data: params.Data,
		}
	} else {
		response = &ResponseWithoutData{
			Code: params.StatusCode,
			Status: status,
			Message: params.Message,
		}
	}

	return response
}