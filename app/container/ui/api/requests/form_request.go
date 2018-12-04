package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

type FormRequest struct{}

func (this *FormRequest) Errors(c *gin.Context, errors error) {
	errorsNew := make(map[string]string)

	for _, error := range errors.(validator.ValidationErrors) {
		errorsNew[error.Field] = error.Tag
	}

	c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errorsNew})
}
