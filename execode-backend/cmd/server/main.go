package main

import (
	"context"
	_ "database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/ChampManZ/ExeCode/v2/docs/execode"
	"github.com/ChampManZ/ExeCode/v2/entities"
	"github.com/ChampManZ/ExeCode/v2/internal/api"
	"github.com/ChampManZ/ExeCode/v2/internal/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "gorm.io/gorm"
)

// @title       Execode API
// @version     1.0
// @description API for Execode code learning environment

// @host     localhost:8080
// @BasePath /
// @schemes  http
func main() {
	// Environment init
	env, err := api.GetEnv()
	if err != nil {
		log.Fatalf("failed to initialize environment variables: %v", err)
	}

	fmt.Println("Environment initialized...")

	// Database init
	err = entities.InitPostgresQL(env.PostgresURL, env.DatabaseUser, env.DatabasePassword, env.DatabaseName, env.PostgresPort)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	err = entities.AutoMigrate()
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	fmt.Println("Database initialized...")

	authEnv, _ := auth.GetEnv()
	err = authEnv.InitKeys()
	if err != nil {
		log.Fatalf("failed to initialize keys: %v", err)
	}

	fmt.Println("Starting server...")

	// Setup
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.JWTWithConfig(authEnv.JwtConfig()))

	// e.POST("/login", auth.LoginHandler)
	// e.GET("/refresh", auth.RefreshHandler)
	// Execute apis
	e.GET("/execute/runtimes", env.RuntimeHandler)
	e.POST("/execute", env.ExecuteHandler)

	// CRUD apis
	e.POST("/users", api.CreateUserHandler)
	e.GET("/users", api.GetUsersHandler)
	e.GET("/users/:userID", api.GetUserHandler)

	e.POST("/classes", api.CreateClassHandler)
	e.GET("/classes", api.GetClassesHandler)
	e.GET("/users/:userID/classes", api.GetUserClassesHandler)
	e.GET("/classes/:classID", api.GetClassHandler)
	e.DELETE("/classes/:classID", api.DeleteClassHandler)

	e.POST("/lectures", api.CreateLectureHandler)
	e.GET("/classes/:class/lectures", api.GetClassLecturesHandler)
	e.GET("/lectures/:lectureID", api.GetLectureHandler)
	e.DELETE("/lectures/:lectureID", api.DeleteLectureHandler)

	e.POST("/problems", api.CreateProblemHandler)
	e.GET("/problems", api.GetProblemsHandler)
	e.GET("/classes/:lecture/problems", api.GetClassProblemsHandler)
	e.GET("/problems/:problemID", api.GetProblemHandler)
	e.DELETE("/problems/:problemID", api.DeleteProblemHandler)

	// Utils
	e.GET("/", healthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	bindTo := fmt.Sprintf("%v:%d", env.BindURL, env.BindPort)
	go func() {
		if err := e.Start(bindTo); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Starting server shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	fmt.Println("Shutdown complete")
}

// healthCheck godoc
// @Summary     Show the status of server.
// @Description get the status of server.
// @Tags        Health
// @Accept      */*
// @Produce     json
// @Success     200 {object} main.healthCheck.response
// @Router      / [get]
func healthCheck(c echo.Context) error {
	type response struct {
		Message string `json:"message"`
	}

	return c.JSON(http.StatusOK, response{"server is up and running"})
}
