package responses

import (
	"encoding/json"
	"net/http"
	"task_API/pkg/errors"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *Error      `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

type Meta struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
	Total   int `json:"total,omitempty"`
}

func WriteSuccess(writer http.ResponseWriter, data interface{}, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	response := Response{
		Success: true,
		Data:    data,
	}

	json.NewEncoder(writer).Encode(response)
}

func WriterError(writer http.ResponseWriter, appError *errors.AppError) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(appError.Code)

	response := Response{
		Success: false,
		Error: &Error{
			Code:    appError.Code,
			Message: appError.Message,
			Details: appError.Err.Error(),
		},
	}

	json.NewEncoder(writer).Encode(response)
}

func Paginated(writer http.ResponseWriter, data interface{}, page, perPage, total int) {
	WriteSuccess(writer, data, http.StatusOK)
}
