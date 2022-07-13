package math

const monthsInAYear = 12
const daysInAMonth = 30
const daysInAYear = monthsInAYear * daysInAMonth

func CalculateTotalDueAmount(amount float64, interestRate float64, installments int) float64 {
	return amount + calculateTotalInterest(amount, interestRate, calculateInterestDays(installments))
}

func CalculateInstallmentValue(totalAmount float64, installments int) float64 {
	return totalAmount / float64(installments)
}

func calculateInterestDays(installments int) int {
	return installments * daysInAMonth
}

func calculateTotalInterest(amount float64, interestRate float64, interestDays int) float64 {
	return amount * ((interestRate / 100) * (float64(interestDays) / daysInAYear))
}
