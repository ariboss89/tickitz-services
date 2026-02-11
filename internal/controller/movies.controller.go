package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/service"
	"github.com/gin-gonic/gin"
)

type MovieController struct {
	movieService *service.MovieService
}

func NewMovieController(movieService *service.MovieService) *MovieController {
	return &MovieController{
		movieService: movieService,
	}
}

// Get Movies By Status godoc
// @Summary      Get Movies By Status
// @Tags         movies
// @Produce      json
// @Param        status		query string  false  "Status Movies (popular, upcoming or now_showing)"
// @Success      200  {object}  dto.Movies
// @Failure 		 500 {object} dto.ResponseError
// @Router       /movies [get]
func (m MovieController) GetMoviesByStatus(c *gin.Context) {
	status := c.Query("status")
	var moviesQuery dto.MoviesQuery
	if err := c.ShouldBindQuery(&moviesQuery); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}

	data, err := m.movieService.GetMoviesByStatus(c.Request.Context(), status)
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

// Get Movie Genres By Id godoc
// @Summary      Get Movie Genres By Id
// @Tags         movies
// @Produce      json
// @Param        id		path int  true  "Id Movie"
// @Success      200  {object}  dto.MovieGenres
// @Failure 		 404 {object} dto.ResponseError
// @Failure 		 500 {object} dto.ResponseError
// @Router       /movies/genres/{id} [get]
func (m MovieController) GetMovieGenresById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

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

	data, err := m.movieService.GetMovieGenresById(c.Request.Context(), id)
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

// Get Detail Movies By Id godoc
// @Summary      Get Movies Detail By Id
// @Tags         movies
// @Produce      json
// @Param        id		path int  true  "Id Movie"
// @Success      200  {object}  dto.MovieDetails
// @Failure 		 500 {object} dto.ResponseError
// @Failure 		 404 {object} dto.ResponseError
// @Router       /movies/{id} [get]
func (m MovieController) GetMovieDetailById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

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

	data, err := m.movieService.GetMovieDetailById(c.Request.Context(), id)
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

// Search Movies By Id godoc
// @Summary      Search Movies
// @Tags         movies
// @Produce      json
// @Param        title		query string  false  "Movie Title"
// @Param 			 genre query []string false "Genres" collectionFormat(multi)
// @Param        page		query int  true  "Page"
// @Success      200  {object}  dto.SearchMovies
// @Failure 		 500 {object} dto.ResponseError
// @Failure 		 404 {object} dto.ResponseError
// @Router       /movies/search [get]
func (m MovieController) SearchMoviesByTitleAndGenre(c *gin.Context) {
	title := c.Query("title")
	genre := c.QueryArray("genre")
	page := c.Query("page")

	intPage, _ := strconv.Atoi(page)

	var moviesQuery dto.MoviesQuery
	if err := c.ShouldBindQuery(&moviesQuery); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}

	data, err := m.movieService.SearchMoviesByTitleAndGenre(c.Request.Context(), title, genre, intPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}
	totalPage, err := m.movieService.GetTotalPage(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "OK",
		Success: true,
		Data:    []any{data},
		Meta: dto.PaginationMeta{
			Page:      intPage,
			TotalPage: totalPage,
		},
	})
}
