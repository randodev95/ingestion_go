package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidBindCheck(ctx *gin.Context, str interface{}) (*gin.Context, error) {
	if err := ctx.BindJSON(&str); err != nil {
		ctx.JSON(http.StatusBadRequest, "request validation error")
		return ctx, err
	}
	if validatorErr := validate.Struct(&str); validatorErr != nil {
		ctx.JSON(http.StatusBadRequest, "request validation error")
		return ctx, validatorErr
	}

	return ctx, nil
}
