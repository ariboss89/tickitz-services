package router

import (
	"github.com/ariboss89/tickitz-services/internal/controller"
	"github.com/ariboss89/tickitz-services/internal/repository"
	"github.com/ariboss89/tickitz-services/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func RegisterMovieRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	moviesRouter := app.Group("/movies")

	movieRepository := repository.NewMoviesRepository(db)
	movieService := service.NewMovieService(movieRepository, rdb)
	movieController := controller.NewMovieController(movieService)

	moviesRouter.GET("", movieController.GetMoviesByStatus)
	moviesRouter.GET("/genres/:id", movieController.GetMovieGenresById)
	moviesRouter.GET("/search", movieController.SearchMoviesByTitleAndGenre)
	moviesRouter.GET("/:id", movieController.GetMovieDetailById)
}
