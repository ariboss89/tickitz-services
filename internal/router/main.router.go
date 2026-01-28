package router

import (
	"log"
	"net/http"

	_ "github.com/ariboss89/tickitz-services/docs"
	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	app.Use(middleware.CORSMiddleware, MyMiddleware)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Static("/static/img", "public")

	RegisterMovieRouter(app, db, rdb)
	RegisterGenreRouter(app, db)
	RegisterActorRouter(app, db, rdb)
	RegisterAuthRouter(app, db, rdb)
	RegisterUserRouter(app, db)
	RegisterAdminRouter(app, db, rdb)
	RegisterOrderRouter(app, db, rdb)

	app.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Msg:     "No route here, go to http://localhost:8002/swagger/index.html",
			Error:   "route not found",
			Success: false,
		})
	})
}

func MyMiddleware(c *gin.Context) {
	// sebelum controller di alur request
	// request -> middleware -> controller
	log.Println("BEFORE")
	c.Next()
	log.Println("AFTER")
	// sesudah controller di alur response
	// controller -> middleware -> response
}

// func GetRootMiddleware(c *gin.Context) {
// 	defer c.Next()
// 	log.Printf("HOST: %s\n", c.GetHeader("Origin"))
// }
