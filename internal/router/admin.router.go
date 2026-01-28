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

func RegisterAdminRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	adminRouter := app.Group("/admin")
	adminRepository := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepository)
	adminController := controller.NewAdminController(adminService)
	adminRouter.Use(middleware.VerifyJWT)
	adminRouter.Use(middleware.IsBlackListed(rdb))

	adminRouter.GET("/movies", middleware.AuthRole("admin"), adminController.GetAllMovies)
	adminRouter.POST("/movies", middleware.AuthRole("admin"), adminController.PostMovie)
	adminRouter.PATCH("/movies/update/:id", middleware.AuthRole("admin"), adminController.UpdateMovie)
	adminRouter.PATCH("/movies/delete/:id", middleware.AuthRole("admin"), adminController.DeleteMovie)
}
