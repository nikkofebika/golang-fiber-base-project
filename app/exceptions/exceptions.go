package exceptions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type BaseException struct {
	StatusCode int                 `json:"-"`
	Message    string              `json:"message,omitempty"`
	Errors     map[string][]string `json:"errors,omitempty"`
}

func (be *BaseException) Error() string {
	return be.Message
}

func NewHttpException(statusCode int, message ...string) *BaseException {
	if statusCode < 100 || statusCode > 599 {
		statusCode = fiber.StatusInternalServerError
	}

	defaultMessage := utils.StatusMessage(statusCode)

	if len(message) > 0 && message[0] != "" {
		defaultMessage = message[0]
	}

	return &BaseException{
		StatusCode: statusCode,
		Message:    defaultMessage,
	}
}

func NewBadRequestException(message ...string) *BaseException {
	statusCode := fiber.StatusBadRequest
	defaultMessage := utils.StatusMessage(statusCode)

	if len(message) > 0 && message[0] != "" {
		defaultMessage = message[0]
	}

	return &BaseException{
		StatusCode: statusCode,
		Message:    defaultMessage,
	}
}

func NewNotFoundException(message ...string) *BaseException {
	statusCode := fiber.StatusNotFound
	defaultMessage := utils.StatusMessage(statusCode)

	if len(message) > 0 && message[0] != "" {
		defaultMessage = message[0]
	}

	return &BaseException{
		StatusCode: statusCode,
		Message:    defaultMessage,
	}
}

func NewUnauthorizedException() *BaseException {
	statusCode := fiber.StatusUnauthorized
	return &BaseException{
		StatusCode: statusCode,
		Message:    utils.StatusMessage(statusCode),
	}
}

type ValidationException struct {
	*BaseException
}

func NewValidationException(errors map[string][]string) *ValidationException {
	statusCode := fiber.StatusUnprocessableEntity
	return &ValidationException{
		BaseException: &BaseException{
			StatusCode: statusCode,
			Message:    utils.StatusMessage(statusCode),
			Errors:     errors,
		},
	}
}
