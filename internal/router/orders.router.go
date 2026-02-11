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

func RegisterOrderRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	orderRouter := app.Group("/order")

	orderRepository := repository.NewOrdersRepository(db)
	orderService := service.NewOrderService(orderRepository, db)
	orderController := controller.NewOrderController(orderService)

	orderRouter.Use(middleware.VerifyJWT)
	orderRouter.Use(middleware.IsBlackListed(rdb))

	orderRouter.GET("/schedule", middleware.AuthRole("user"), orderController.GetSchedules)
	orderRouter.POST("/", middleware.AuthRole("user"), orderController.CreateOrder)
}
