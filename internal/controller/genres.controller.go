package controller

import (
	"net/http"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/service"
	"github.com/gin-gonic/gin"
)

type GenreController struct {
	genreService *service.GenreService
}

func NewGenreController(genreService *service.GenreService) *GenreController {
	return &GenreController{
		genreService: genreService,
	}
}

// Get All Genres godoc
// @Summary      Get All Genres
// @Tags         genre
// @Produce      json
// @Success      200  {object}  dto.Genres
// @Failure 		 500 {object} dto.ResponseError
// @Router       /genres [get]
func (m GenreController) GetAllGenres(c *gin.Context) {
	data, err := m.genreService.GetAllGenres(c.Request.Context())
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
