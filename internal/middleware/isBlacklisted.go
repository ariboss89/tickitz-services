package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func IsBlackListed(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Get("tokenJWT")
		rkey := "ari:tickitz:logout" + fmt.Sprint(token)
		rsc := rdb.Get(c, rkey)

		if rsc.Err() == nil {
			tokenStore := rsc.Val()
			if token == tokenStore {
				log.Println("token is blacklisted")
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseError{
					Msg:     "Unauthorized Access",
					Success: false,
					Error:   "Invalid Token",
				})
				return
			}
		}
		c.Next()
	}
}
