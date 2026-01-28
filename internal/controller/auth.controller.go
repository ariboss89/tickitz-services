package controller

import (
	"net/http"
	"strings"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/service"
	"github.com/ariboss89/tickitz-services/pkg"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type AuthController struct {
	authService *service.AuthService
	redis       *redis.Client
}

func NewAuthController(authService *service.AuthService, rdb *redis.Client) *AuthController {
	return &AuthController{
		authService: authService,
		redis:       rdb,
	}
}

// Register User godoc
// @Summary      Register new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        auth	 body dto.NewUser  true  "User Registration Body"
// @Success      200  {object}  dto.RegisterResponse
// @Failure 		 400 {object} dto.ResponseError
// @Failure 		 500 {object} dto.ResponseError
// @Router       /auth/register [post]
func (a AuthController) Register(c *gin.Context) {
	var newUser dto.NewUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		str := err.Error()
		if strings.Contains(str, "Field") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Invalid Body",
				Success: false,
				Error:   "invalid body",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}
	data, err := a.authService.Register(c.Request.Context(), newUser)
	if err != nil {
		if strings.Contains(err.Error(), "user exist") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Bad Request",
				Success: false,
				Error:   "Email already in use",
			})
			return
		}
		// Unexpected Error
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg":  "Registered successfully",
		"data": data,
	})
}

// Login User godoc
// @Summary      Login new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        auth	 body dto.Login  true  "User Login Body"
// @Success      200  {object}  dto.LoginResponse
// @Failure 		 400 {object} dto.ResponseError
// @Failure 		 500 {object} dto.ResponseError
// @Failure 		 403 {object} dto.ResponseError
// @Router       /auth/login [post]
func (a AuthController) Login(c *gin.Context) {
	var login dto.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		str := err.Error()
		if strings.Contains(str, "Field") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Invalid Body",
				Success: false,
				Error:   "invalid body",
			})
			return
		}
	}
	data, err := a.authService.Login(c.Request.Context(), login)
	if err != nil {
		if strings.Contains(err.Error(), "user exist") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Bad Request",
				Success: false,
				Error:   "Email already in use",
			})
			return
		}
		if strings.Contains(err.Error(), "username or password is wrong") {
			c.JSON(http.StatusForbidden, dto.ResponseError{
				Msg:     "username or password is wrong",
				Success: false,
				Error:   "username or password is wrong",
			})
			return
		}
		// Unexpected Error
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "Login Success",
		"data": data,
	})
}

// Update Password godoc
// @Summary      Change Password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        auth	 body dto.UpdatePassword  true  "User Update Body"
// @Success      200  {object}  dto.UpdatePassword
// @Failure 		 400 {object} dto.ResponseError
// @Failure 		 500 {object} dto.ResponseError
// @Router       /auth/password [patch]
// @security		 BearerAuth
func (a AuthController) UpdatePassword(c *gin.Context) {
	var updatePassword dto.UpdatePassword
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

	accessToken, _ := token.(pkg.JWTClaims)

	if err := c.ShouldBindJSON(&updatePassword); err != nil {
		str := err.Error()
		if strings.Contains(str, "Field") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Invalid Body",
				Success: false,
				Error:   "invalid body",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}
	email := accessToken.Email
	data, err := a.authService.UpdatePassword(c.Request.Context(), updatePassword, email)
	if err != nil {
		// Unexpected Error
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}
	if data.Email != accessToken.Email {
		c.JSON(http.StatusBadRequest, dto.ResponseError{
			Msg:     "Old Password Is Not Match ",
			Success: false,
			Error:   "old password is not match",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Update successfully",
			"data": data,
		})
	}
}

// Logout User godoc
// @Summary      Logout user
// @Tags         auth
// @Produce      json
// @Success      201  {object}  dto.LogoutResponse
// @Failure 		 400 {object} dto.ResponseError
// @Failure 		 500 {object} dto.ResponseError
// @Failure 		 403 {object} dto.ResponseError
// @Router       /auth/logout [delete]
// @security		 BearerAuth
func (a AuthController) Logout(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	tokenJWT := strings.Split(auth, "Bearer ")

	token := tokenJWT[1]

	// if err := a.authService.LogoutUser(c.Request.Context(), token); err != nil {
	_, err := a.authService.LogoutUser(c.Request.Context(), token)

	if err != nil {
		str := err.Error()
		if strings.Contains(str, "token obsoleted") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Token Obsoleted",
				Success: false,
				Error:   "logout failed",
			})
			return
		}
		// Unexpected Error
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}
	c.JSON(http.StatusCreated, dto.LogoutResponse{
		Msg: "Logout successfully",
	})
}
