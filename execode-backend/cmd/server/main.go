package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/ChampManZ/ExeCode/v2/docs/execode"
	"github.com/ChampManZ/ExeCode/v2/internal/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title       Execode API
// @version     1.0
// @description API for Execode code learning environment

// @host     localhost:8080
// @BasePath /
// @schemes  http
func main() {
	// Setup
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.GET("/execute/runtimes", api.RuntimeHandler)
	e.POST("/execute", api.ExecuteHandler)

	e.GET("/", healthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
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
