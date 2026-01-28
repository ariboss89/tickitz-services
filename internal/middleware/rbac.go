package middleware

import (
	"net/http"
	"slices"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/pkg"
	"github.com/gin-gonic/gin"
)

func AuthRole(role ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, isExist := c.Get("token")
		if !isExist {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
				Msg:     "Forbidden Access",
				Success: false,
				Data:    []any{},
				Error:   "Access Denied",
			})
			return
		}
		accessToken, ok := token.(pkg.JWTClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
				Msg:     "Internal Server Error",
				Success: false,
				Data:    []any{},
				Error:   "internal server error",
			})
			return
		}
		isAuthorized := slices.Contains(role, accessToken.Role)

		if !isAuthorized {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
				Msg:     "Forbidden Access",
				Success: false,
				Data:    []any{},
				Error:   "Access Denied",
			})
			return
		}
		c.Next()
	}
}
