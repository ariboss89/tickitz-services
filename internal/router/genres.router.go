package router

import (
	"github.com/ariboss89/tickitz-services/internal/controller"
	"github.com/ariboss89/tickitz-services/internal/repository"
	"github.com/ariboss89/tickitz-services/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterGenreRouter(app *gin.Engine, db *pgxpool.Pool) {
	genresRouter := app.Group("/genres")

	genreRepository := repository.NewGenresRepository(db)
	genreService := service.NewGenreService(genreRepository)
	genreController := controller.NewGenreController(genreService)

	genresRouter.GET("/", genreController.GetAllGenres)
}
