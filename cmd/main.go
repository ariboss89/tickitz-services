package main

import (
	"log"

	"github.com/ariboss89/tickitz-services/internal/config"
	"github.com/ariboss89/tickitz-services/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           Tickitz Services
// @version         1.0
// @description     Backend Services For Tickitz Project
// @host      			localhost:8002
// @BasePath  			/

// @securityDefinitions.apikey	BearerAuth
// @in													header
// @name 												Authorization
// @description 								Type "Bearer" followed by space and JWT Token
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to Load env")
		return
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
	app.Run(":8002")
}
