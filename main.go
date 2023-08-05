package main

import (
	"bwa_start_up/handler"
	"bwa_start_up/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa_start_up?charset=utf8mb4&parseTime=True&loc=Local"

	// dsn := "host=localhost user=oziduhu.halawa password=iceM1422 dbname=bwa_campaign port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	router.Run()
}

// handler -> service -> repo -> db
