package errors

import (
	"fmt"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	Err     string `json:"error"`
}

func (app *AppError) Error() string {
	if app.Error != nil {
		fmt.Sprintf("%s: %s", app.Message, app.Error)
	}
	return app.Message
}
