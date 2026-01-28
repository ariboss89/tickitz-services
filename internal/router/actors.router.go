package router

import (
	"github.com/ariboss89/tickitz-services/internal/controller"
	"github.com/ariboss89/tickitz-services/internal/middleware"
	"github.com/ariboss89/tickitz-services/internal/repository"
	"github.com/ariboss89/tickitz-services/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func RegisterActorRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	actorRouter := app.Group("/actors")

	actorRepository := repository.NewActorsRepository(db)
	actorService := service.NewActorService(actorRepository)
	actorController := controller.NewActorController(actorService)
	actorRouter.Use(middleware.VerifyJWT)
	actorRouter.Use(middleware.IsBlackListed(rdb))

	actorRouter.GET("/", actorController.GetAllActors)
}
