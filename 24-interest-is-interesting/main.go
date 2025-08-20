package main

import "fmt"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance >= 0 && balance < 1000:
		return 0.5
	case balance >= 1000 && balance < 5000:
		return 1.621
	case balance >= 5000:
		return 2.475
	}

	return 3.213
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance)) * (balance / 100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	year := 0
	totalBalance := balance
	for totalBalance < targetBalance {
		totalBalance += Interest(totalBalance)
		year += 1
	}

	return year
}

func main() {
	fmt.Println(InterestRate(200.75))
	// => 0.5

	fmt.Println(Interest(200.75))
	// => 1.003750

	fmt.Println(AnnualBalanceUpdate(200.75))
	// => 201.75375

	balance := 200.75
	targetBalance := 214.88
	fmt.Println(YearsBeforeDesiredBalance(balance, targetBalance))
	// => 14
}
