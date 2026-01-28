package router

import (
	"github.com/ariboss89/tickitz-services/internal/controller"
	"github.com/ariboss89/tickitz-services/internal/middleware"
	"github.com/ariboss89/tickitz-services/internal/repository"
	"github.com/ariboss89/tickitz-services/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUserRouter(app *gin.Engine, db *pgxpool.Pool) {
	userRouter := app.Group("/user")

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	userRouter.Use(middleware.VerifyJWT)

	rbacMiddleware := middleware.AuthRole("user", "admin")

	userRouter.GET("/profile", rbacMiddleware, userController.GetUserProfileByEmail)
	userRouter.GET("/history", middleware.AuthRole("user"), userController.GetHistory)
	userRouter.PATCH("/", rbacMiddleware, userController.UpdateProfile)
}
