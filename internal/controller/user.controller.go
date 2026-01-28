package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"time"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/service"
	"github.com/ariboss89/tickitz-services/pkg"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Get User Profiles By Email
// @Summary      Get User Profiles By Email
// @Tags         user
// @Produce      json
// @Success      200  {object}  dto.Users
// @Failure 		 404 {object} dto.ResponseError
// @Failure 		 403 {object} dto.ResponseError
// @Failure 		 500 {object} dto.ResponseError
// @Router       /user/profile [get]
// @security			BearerAuth
func (u UserController) GetUserProfileByEmail(c *gin.Context) {
	token, isExist := c.Get("token")
	if !isExist {
		c.JSON(http.StatusForbidden, dto.ResponseError{
			Msg:     "Forbidden Access",
			Success: false,
			Error:   "Access Denied",
		})
		return
	}
	accessToken, _ := token.(pkg.JWTClaims)

	data, err := u.userService.GetUserProfileByEmail(c.Request.Context(), accessToken.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "OK",
		"data": data,
	})
}

// Get User History
// @Summary      Get User History
// @Tags         user
// @Produce      json
// @Success      200  {object}  dto.History
// @Failure 		 403 {object} dto.ResponseError
// @Failure 		 500 {object} dto.ResponseError
// @Router       /user/history [get]
// @security			BearerAuth
func (u UserController) GetHistory(c *gin.Context) {
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

	data, err := u.userService.GetHistory(c.Request.Context(), accessToken.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "OK",
		"data": data,
	})
}

// Update Profile
// @Summary      Update Profile
// @Tags         user
// @Accept			 json
// @Produce      json
// @Param        first_name	formData string false  "First Name"
// @Param        last_name	formData string false  "Last Name"
// @Param        image	formData file false  "Update Profile Body"
// @Success      200  {object}  dto.MovieGenres
// @Failure 		 400 {object} dto.ResponseError
// @Failure 		 403 {object} dto.ResponseError
// @Failure 		 404 {object} dto.ResponseError
// @Failure 		 500 {object} dto.ResponseError
// @Router       /user/ [patch]
// @Security			BearerAuth
func (u UserController) UpdateProfile(c *gin.Context) {
	const maxSize = 1 * 1024 * 1024
	var updateUser dto.UpdateProfile
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

	if err := c.ShouldBindWith(&updateUser, binding.FormMultipart); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
			Data:    []any{},
		})
		return
	}
	if updateUser.ImageFile != nil {
		ext := path.Ext(updateUser.ImageFile.Filename)
		re := regexp.MustCompile("^[.](jpg|png)$")
		if !re.Match([]byte(ext)) {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "File have to be jpg or png",
				Error:   "Bad Request",
				Success: false,
			})
			return
		}
		// validasi ukuran
		if updateUser.ImageFile.Size > maxSize {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "File maximum 1 MB",
				Error:   "Bad Request",
				Success: false,
			})
			return
		}

		data, errPhoto := u.userService.GetUserProfileByEmail(c.Request.Context(), accessToken.Email)

		if errPhoto != nil {
			// Unexpected Error
			log.Println("photo is not uploaded yet")
		}

		if data[0].Image != nil {
			prevPhoto := data[0].Image
			filePath := "public" + *prevPhoto
			absPath, err := filepath.Abs(filePath)
			errPath := os.Remove(absPath)

			if errPath != nil {
				// Log a fatal error if removal fails (e.g., file not found, permission issues)
				log.Println(err)
			}
		}

		filename := fmt.Sprintf("%d_profile_%d%s", time.Now().UnixNano(), accessToken.Id, ext)
		updateUser.ImageName = filename

		if e := c.SaveUploadedFile(updateUser.ImageFile, filepath.Join("public", "profile", filename)); e != nil {
			log.Println(e.Error())
			c.JSON(http.StatusInternalServerError, dto.ResponseError{
				Msg:     "Internal Server Error",
				Success: false,
				Error:   "internal server error",
			})
			return
		}
	}

	email := accessToken.Email

	if email != "" {
		if err := u.userService.UpdateProfile(c.Request.Context(), updateUser, email); err != nil {
			// Unexpected Error
			c.JSON(http.StatusInternalServerError, dto.ResponseError{
				Msg:     "Internal Server Error",
				Success: false,
				Error:   "internal server error",
			})
			return
		}

		//absPath, _ := filepath.Abs("public/profile/" + updateUser.ImageName)

		photoUrl := fmt.Sprintf("http://localhost:8002/static/img/profile/%s", updateUser.ImageName)

		c.JSON(http.StatusOK, gin.H{
			"msg": "Update successfully",
			// "img": absPath,
			"img": photoUrl,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusForbidden, dto.ResponseError{
			Msg:     "Forbidden Access",
			Success: false,
			Error:   "Access Denied",
		})
		return
	}
}
