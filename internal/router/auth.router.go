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

func RegisterAuthRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	authRouter := app.Group("/auth")

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, rdb)
	authController := controller.NewAuthController(authService, rdb)

	authRouter.POST("/register", authController.Register)
	authRouter.POST("/login", authController.Login)

	authRouter.Use(middleware.VerifyJWT)
	authRouter.Use(middleware.IsBlackListed(rdb))
	rbacMiddleware := middleware.AuthRole("user", "admin")

	authRouter.PATCH("/password", rbacMiddleware, authController.UpdatePassword)
	authRouter.DELETE("/logout", rbacMiddleware, authController.Logout)
}
