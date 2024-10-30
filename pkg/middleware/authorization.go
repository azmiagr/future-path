package middleware

import (
	"errors"
	"future-path/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *middleware) OnlyAdmin(ctx *gin.Context) {
	admin, err := m.jwtAuth.GetLoginAdmin(ctx)
	if err != nil {
		response.Error(ctx, http.StatusForbidden, "failed get login admin", err)
		return
	}

	if admin.ID_Admin != 1 {
		response.Error(ctx, http.StatusForbidden, "this endpoint cannot be access", errors.New("user dont have access"))
		ctx.Abort()
		return
	}

	ctx.Next()
}
