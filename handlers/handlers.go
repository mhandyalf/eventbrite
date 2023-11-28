package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetMarketingTools godoc
// @Summary Get marketing tools information
// @Description Get information about the marketing tools used by the organization.
// @Tags Marketing Tools
// @Produce json
// @Param orgID path int true "Organization ID"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 404 {object} map[string]string "Organization not found"} "Organization not found"
// @Failure 500 {object} map[string]string "Failed to fetch marketing tools"} "Failed to fetch marketing tools"
// @Router /organizations/{orgID}/marketing [get]
func GetMarketingTools(c echo.Context) error {
	marketingTools := map[string]interface{}{
		"socialMedia":    "Facebook, Twitter, Instagram",
		"emailMarketing": "Mailchimp, SendGrid",
	}

	if marketingTools == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch marketing tools"})
	}

	return c.JSON(http.StatusOK, marketingTools)
}

// GetPayoutsInfo godoc
// @Summary Get payouts information
// @Description Get information about the payouts used by the organization.
// @Tags Payouts Information
// @Produce json
// @Param orgID path int true "Organization ID"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 404 {object} map[string]string "Organization not found"} "Organization not found"
// @Failure 500 {object} map[string]string "Failed to fetch payouts information"} "Failed to fetch payouts information"
// @Router /organizations/{orgID}/payouts [get]
func GetPayoutsInfo(c echo.Context) error {
	payoutsInfo := map[string]interface{}{
		"paymentMethods":  "PayPal, Bank Transfer",
		"payoutFrequency": "Monthly",
	}

	if payoutsInfo == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch payouts information"})
	}

	return c.JSON(http.StatusOK, payoutsInfo)
}

// GetOrgInfo godoc
// @Summary Get organization information
// @Description Get information about the organization.
// @Tags Organization Information
// @Produce json
// @Param orgID path int true "Organization ID"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 404 {object} map[string]string "Organization not found"} "Organization not found"
// @Failure 500 {object} map[string]string "Failed to fetch organization information"} "Failed to fetch organization information"
// @Router /organizations/{orgID}/info [get]
func GetOrgInfo(c echo.Context) error {
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
