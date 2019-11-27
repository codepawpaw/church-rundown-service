package handler

import (
	"net/http"

	dto "../../dto"
)

func construct(data []byte, err error) *dto.HttpResponse {
	httpStatus := http.StatusOK
	errorMessage := ""
	responseData := string(data)

	if err != nil {
		responseData = ""
		httpStatus = http.StatusInternalServerError
		errorMessage = err.Error()
	}

	response := &dto.HttpResponse{
		Data:         responseData,
		ErrorMessage: errorMessage,
		Status:       httpStatus,
	}

	return response
}
