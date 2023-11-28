package main

import (
	_ "eventbrite/docs"
	"eventbrite/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)


// @title Eventbrite API
// @version 1.0
// @description This is a sample server for Eventbrite API.
// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Add your handlers here
	e.GET("/organizations/marketing", handlers.GetMarketingTools)
	e.GET("/organizations/payouts", handlers.GetPayoutsInfo)
	e.GET("/organizations/info", handlers.GetOrgInfo)

	e.Start(":8080")
}
