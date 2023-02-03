package main

import (
	"log"
	"yuhuuuMit/config"
	userHandler "yuhuuuMit/feature/user/handler"
	userRepository "yuhuuuMit/feature/user/repository"
	userService "yuhuuuMit/feature/user/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	c := config.LoadEnv()    // panggil fungsi baca env
	db := config.OpenDB(c)   // panggil fungsi konek ke database
	config.GormMigration(db) // auto migration

	v := validator.New()

	userRep := userRepository.New(db)
	userSrv := userService.New(userRep, v)
	userHdl := userHandler.New(&userSrv)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom}, method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	e.POST("/register", userHdl.RegisterHdl())
	e.POST("/login", userHdl.LoginHdl())

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
