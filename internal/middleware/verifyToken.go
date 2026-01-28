package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	result := strings.Split(bearerToken, " ")
	if result[0] != "Bearer" {
		log.Println("token is not bearer token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Msg:     "Unauthorized Access",
			Success: false,
			Data:    []any{},
			Error:   "Invalid Token",
		})
		return
	}

	var jc pkg.JWTClaims
	_, err := jc.VerifyToken(result[1])
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, jwt.ErrTokenExpired) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Msg:     "Unauthorized Access",
				Success: false,
				Data:    []any{},
				Error:   "Expired Token, Please Login Again",
			})
			return
		}
		if errors.Is(err, jwt.ErrTokenInvalidIssuer) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Msg:     "Unauthorized Access",
				Success: false,
				Data:    []any{},
				Error:   "Invalid Token, Please Login Again",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Data:    []any{},
			Error:   "internal server error",
		})
		return
	}
	c.Set("token", jc)
	c.Set("tokenJWT", result[1])
	c.Next()
}
