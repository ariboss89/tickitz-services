package main

import (
	"log"
	"os"

	"github.com/ariboss89/tickitz-services/internal/config"
	"github.com/ariboss89/tickitz-services/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           Tickitz Services
// @version         1.0
// @description     Backend Services For Tickitz Project
// @host      			72.62.120.77:8081
// @BasePath  			/

// @securityDefinitions.apikey	BearerAuth
// @in													header
// @name 												Authorization
// @description 								Type "Bearer" followed by space and JWT Token
func main() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Failed to Load env")
			return
		}
	}
	db, err := config.InitDb()
	rdb := config.InitRedis()
	defer rdb.Close()

	if err != nil {
		log.Println("Failed to Connect to Database")
		return
	}

	app := gin.Default()
	app.MaxMultipartMemory = 64 << 20
	router.Init(app, db, rdb)
	app.Run()
}
