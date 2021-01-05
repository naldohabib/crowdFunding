package main

import (
	"crowFunding/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"

	userHandler "crowFunding/user/handler"
	userService "crowFunding/user/service"
	userRepo "crowFunding/user/repository"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	conStr := viper.GetString("database.postgres")
	port := viper.GetString("port.port")

	db, err := gorm.Open("postgres", conStr)
	if err != nil {
		log.Fatal("Error When Connect to DB " + conStr + " : " + err.Error())
	}

	defer db.Close()

	db.Debug().AutoMigrate(
		&model.User{})

	router := gin.Default()

	userRepo := userRepo.CreateUserRepo(db)
	userService := userService.CreateUserService(userRepo)
	userHandler.CreateUserHandler(router, userService)

	fmt.Println("Starting Web Server at port : " + port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
