package util

import "fmt"

func CheckQuantity(stock, requested int) error {
	if stock-requested >= 0 {
		return nil
	}
	return fmt.Errorf("Insufficient stock. Available quantity: %d", stock)
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func ErrorResponse(err error) APIResponse {
	return APIResponse{
		Success: false,
		Message: "error",
		Error:   err.Error(),
	}
}

func SuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}
