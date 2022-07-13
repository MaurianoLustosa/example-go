package main

import (
	"core-banking/log"
	"core-banking/math"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const port = 8000

var SEQ = 1

type (
	calculationResponse struct {
		ID               int     `json:"id"`
		TotalAmount      float64 `json:"totalAmount"`
		InstallmentValue float64 `json:"installmentValue"`
	}

	calculationRequest struct {
		Amount       float64 `json:"amount"`
		Installments int     `json:"installments"`
		InterestRate float64 `json:"interestRate"`
	}
)

func main() {
	log.Println("Starting application...")

	createdEcho := echo.New()
	handleMiddleWare(createdEcho)
	createRoutes(createdEcho)
	startApplication(createdEcho)
}

func startApplication(echo *echo.Echo) {
	echo.Start(fmt.Sprint(":", port))
}

func handleMiddleWare(echo *echo.Echo) {
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())
}

func createRoutes(echo *echo.Echo) {
	echo.POST("/calculate", handleCalculateValuesHttpRequest)
}

func handleCalculateValuesHttpRequest(context echo.Context) error {
	log.Println("Calculating values...\n")
	var result = getCalculatedValues(processQueryParameters(context))

	return context.JSON(http.StatusOK, result)
}

func processQueryParameters(context echo.Context) (amount float64, installments int, interestRate float64) {
	request := new(calculationRequest)
	context.Bind(request)

	return request.Amount,
		request.Installments,
		request.InterestRate
}

func getCalculatedValues(amount float64, installments int, interestRate float64) *calculationResponse {

	var totalAmount = math.CalculateTotalDueAmount(amount, interestRate, installments)
	var installmentValue = math.CalculateInstallmentValue(totalAmount, installments)

	log.LogCalculatedValues(totalAmount, installmentValue)

	return buildResultObject(totalAmount, installmentValue)
}

func buildResultObject(totalAmount float64, installmentValue float64) *calculationResponse {
	result := new(calculationResponse)
	result.ID = SEQ
	SEQ++
	result.TotalAmount = totalAmount
	result.InstallmentValue = installmentValue

	return result
}
