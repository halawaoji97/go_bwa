package main

import (
	"bwa_start_up/auth"
	"bwa_start_up/handler"
	"bwa_start_up/user"
	"fmt"
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
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo4fQ.lP8NxXCsv6DuYdd_rSWEvxsI5d__MahboLYain-rF8Tc")
	if err != nil {
		fmt.Println("invla")
	}

	fmt.Println(token.Valid)

	userHandler := handler.NewUserHandler(userService, authService)
	userService.SaveAvatar(2, "images/profil.jpeg")

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/check_email", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)
	router.Run()
}

// handler -> service -> repo -> db
