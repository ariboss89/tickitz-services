package controller

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/service"
	"github.com/ariboss89/tickitz-services/pkg"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AdminController struct {
	adminService *service.AdminService
}

func NewAdminController(adminService *service.AdminService) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

// GetAllMovies godoc
// @Summary      Show All Movies For Admin
// @Tags         admin
// @Produce      json
// @Success      200  {object}  dto.Movies
// @Failure 		 500 {object} dto.ResponseError
// @Router       /admin/movies/ [get]
// @security 		 BearerAuth
func (m AdminController) GetAllMovies(c *gin.Context) {
	data, err := m.adminService.GetAllMovies(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ResponseError{
			Msg:     "please login first as admin",
			Success: false,
			Error:   err.Error(),
		})

		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "internal server error",
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "OK",
		"data": data,
	})
}

// Post Movies
// @Summary      Post Movies
// @Tags         admin
// @Accept			 json
// @Produce      json
// @Param        title	formData string false  "Title"
// @Param        synopsis	formData string false  "Synopsis"
// @Param        poster_file	formData file false  "Poster Photo"
// @Param        background_file	formData file false  "Background Photo"
// @Param        release_date	formData string false  "Release Date"
// @Param        duration	formData int false  "Duration Movie"
// @Param        status	formData string false  "Status(upcoming or now_showing)"
// @Param        rating	formData number false  "Rating"
// @Success      200  {object}  dto.Response
// @Failure 		 400 {object} dto.ResponseError
// @Failure 		 403 {object} dto.ResponseError
// @Failure 		 404 {object} dto.ResponseError
// @Failure 		 500 {object} dto.ResponseError
// @Router       /admin/movies/ [post]
// @Security			BearerAuth
func (a AdminController) PostMovie(c *gin.Context) {
	const maxSize = 2 * 1024 * 1024
	var postMovie dto.PostMovies
	token, isExist := c.Get("token")
	if !isExist {
		c.AbortWithStatusJSON(http.StatusForbidden, dto.ResponseError{
			Msg:     "Forbidden Access",
			Success: false,
			Error:   "Access Denied",
		})
		return
	}
	accessToken, _ := token.(pkg.JWTClaims)

	if err := c.ShouldBindWith(&postMovie, binding.FormMultipart); err != nil {
		log.Println(err.Error(), "1")
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}

	if postMovie.BackgroundFile != nil && postMovie.PosterFile != nil {
		extPoster := path.Ext(postMovie.PosterFile.Filename)
		extBg := path.Ext(postMovie.BackgroundFile.Filename)
		re := regexp.MustCompile("^[.](jpg|png)$")
		if !re.Match([]byte(extBg)) || !re.Match([]byte(extPoster)) {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "File have to be jpg or png",
				Error:   "Bad Request",
				Success: false,
			})
			return
		}
		//validasi ukuran
		if postMovie.BackgroundFile.Size > maxSize || postMovie.PosterFile.Size > maxSize {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "File maximum 1 MB",
				Error:   "Bad Request",
				Success: false,
			})
			return
		}

		filenamePoster := fmt.Sprintf("%d_poster_%d%s", time.Now().UnixNano(), accessToken.Id, extPoster)
		filenameBg := fmt.Sprintf("%d_background_%d%s", time.Now().UnixNano(), accessToken.Id, extBg)
		postMovie.Poster_Url = filenamePoster
		postMovie.Background_Url = filenameBg

		if e := c.SaveUploadedFile(postMovie.PosterFile, filepath.Join("public", "poster", filenamePoster)); e != nil {
			log.Printf(e.Error(), 2)
			c.JSON(http.StatusInternalServerError, dto.ResponseError{
				Msg:     "Internal Server Error",
				Success: false,
				Error:   "internal server error",
			})
			return
		}

		if e := c.SaveUploadedFile(postMovie.BackgroundFile, filepath.Join("public", "background", filenameBg)); e != nil {
			log.Printf(e.Error(), 3)
			c.JSON(http.StatusInternalServerError, dto.ResponseError{
				Msg:     "Internal Server Error",
				Success: false,
				Error:   "internal server error",
			})
			return
		}
	}

	if err := c.ShouldBindWith(&postMovie, binding.FormMultipart); err != nil {
		str := err.Error()
		log.Println(3)
		log.Println(str)
		if strings.Contains(str, "Field") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Invalid Body",
				Success: false,
				Error:   "invalid body",
			})
			return
		}
		if strings.Contains(str, "Empty") {
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

	if err := a.adminService.PostMovie(c.Request.Context(), postMovie); err != nil {
		str := err.Error()
		log.Println(4)
		log.Println(str)
		if strings.Contains(str, "empty") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Invalid Body",
				Success: false,
				Error:   str,
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
	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Movies Inserted Successfully",
		Success: true,
		Data:    []any{},
	})
}

// UpdateMovie godoc
// @Summary      Update Movies For Admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        id		path int  true  "Movie Id"
// @Param        movies	 body dto.UpdateMovies  true  "Update Movies Body"
// @Success      200  {object}  dto.UpdateMovies
// @Failure 		 400 {object} dto.ResponseError
// @Failure 		 500 {object} dto.ResponseError
// @Router       /admin/movies/update/{id} [patch]
// @security 		 BearerAuth
func (a AdminController) UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	strId, _ := strconv.Atoi(id)

	var updateMovie dto.UpdateMovies

	if err := c.ShouldBindJSON(&updateMovie); err != nil {
		str := err.Error()
		if strings.Contains(str, "Field") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Invalid Body",
				Success: false,
				Error:   "invalid body",
			})
			return
		}
		if strings.Contains(str, "Empty") {
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

	if err := a.adminService.UpdateMovie(c.Request.Context(), updateMovie, strId); err != nil {
		str := err.Error()
		if strings.Contains(str, "empty") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Invalid Body",
				Success: false,
				Error:   str,
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
	c.JSON(http.StatusOK, gin.H{
		"msg":  "Movies Updated Successfully",
		"data": []any{},
	})
}

// DeleteMovie godoc
// @Summary      Delete Movies For Admin
// @Tags         admin
// @Produce      json
// @Param        id		path int  true  "Movie Id"
// @Success      200  {object}  dto.UpdateMovies
// @Failure 		 400 {object} dto.ResponseError
// @Failure 		 500 {object} dto.ResponseError
// @Router       /admin/movies/delete/{id} [delete]
// @security 		 BearerAuth
func (a AdminController) DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	strId, _ := strconv.Atoi(id)

	var moviesParam dto.MoviesParam

	if err := c.ShouldBindUri(&moviesParam); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}

	data, err := a.adminService.DeleteMovie(c.Request.Context(), strId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "Deleted Successfully",
		"data": data,
	})
}
