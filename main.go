package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "eventbrite/docs"
)

// @title Eventbrite API
// @version 1.0
// @description This is a sample server for Eventbrite API.
// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Add your handlers here
	e.GET("/organizations/marketing", getMarketingTools)
	e.GET("/organizations/payouts", getPayoutsInfo)
	e.GET("/organizations/info", getOrgInfo)

	e.Start(":8080")
}

// getMarketingTools godoc
// @Summary Get marketing tools information
// @Tags Marketing Tools
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /organizations/marketing [get]
func getMarketingTools(c echo.Context) error {
	// Your logic here
	marketingTools := map[string]interface{}{
		"socialMedia":    "Facebook, Twitter, Instagram",
		"emailMarketing": "Mailchimp, SendGrid",
	}

	if marketingTools == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch marketing tools"})
	}

	return c.JSON(http.StatusOK, marketingTools)
}

// getPayoutsInfo godoc
// @Summary Get payouts information
// @Tags Payouts Information
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /organizations/payouts [get]
func getPayoutsInfo(c echo.Context) error {
	// Your logic here
	payoutsInfo := map[string]interface{}{
		"paymentMethods":  "PayPal, Bank Transfer",
		"payoutFrequency": "Monthly",
	}

	if payoutsInfo == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch payouts information"})
	}

	return c.JSON(http.StatusOK, payoutsInfo)
}

// getOrgInfo godoc
// @Summary Get organization information
// @Tags Organization Information
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /organizations/info [get]
func getOrgInfo(c echo.Context) error {
	// Your logic here
	orgInfo := map[string]interface{}{
		"orgName":       "Eventbrite Inc.",
		"contactEmail":  "info@eventbrite.com",
		"website":       "https://www.eventbrite.com",
		"foundingYear":  2006,
		"employeeCount": 1000,
		"headquarters":  "San Francisco, CA",
	}

	if orgInfo == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch organization information"})
	}

	return c.JSON(http.StatusOK, orgInfo)
}

