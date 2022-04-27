package server

import (
	"github.com/gin-gonic/gin"
	"github.com/guil95/mutantdna/internal/domain"
	"net/http"
)

func handleError(err error, ctx *gin.Context) {
	switch err.(type) {
	case domain.HumanError:
		ctx.Status(http.StatusForbidden)
		return
	default:
		ctx.Status(http.StatusInternalServerError)
		return
	}
}
