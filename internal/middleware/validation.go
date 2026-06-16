package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidRequest(v *validator.Validate, data interface{}) []ErrorResponse {
	var errors []ErrorResponse

	err := v.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ErrorResponse{
				Field:   err.Field(),
				Message: getErrorMessage(err),
			})
		}
	}
	return errors
}

// getErrorMessage retorna mensagem de erro amigável
func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "Campo obrigatório"
	case "email":
		return "Email inválido"
	case "min":
		return "Valor mínimo: " + err.Param()
	case "max":
		return "Valor máximo: " + err.Param()
	case "uuid":
		return "UUID inválido"
	case "oneof":
		return "Valor deve ser um dos permitidos: " + err.Param()
	default:
		return "Validação falhou: " + err.Tag()
	}
}

func ValidationErrorResponse(c *gin.Context, errors []ErrorResponse) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   "Validacao Falhou",
		"errors":  errors,
	})
}
