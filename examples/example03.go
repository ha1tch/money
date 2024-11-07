package main

import (
	"fmt"

	"github.com/ha1tch/money"
)

func main() {
	// Basic Arithmetic Operations
	fmt.Println("=== Basic Arithmetic Operations ===")

	// Addition
	price, _ := money.New(2000, "USD") // $20.00
	tax, _ := money.New(160, "USD")    // $1.60
	total, _ := price.Add(tax)
	fmt.Printf("Addition: %s + %s = %s\n",
		price.Format(), tax.Format(), total.Format())

	// Subtraction
	payment, _ := money.New(2500, "USD") // $25.00
	change, _ := payment.Subtract(total)
	fmt.Printf("Subtraction: %s - %s = %s\n",
		payment.Format(), total.Format(), change.Format())

	// Multiplication
	quantity := 3.0
	itemPrice, _ := money.New(1099, "USD") // $10.99
	subtotal := itemPrice.Multiply(quantity)
	fmt.Printf("Multiplication: %s × %.1f = %s\n",
		itemPrice.Format(), quantity, subtotal.Format())

	// Percentage Operations
	fmt.Println("\n=== Percentage Operations ===")

	originalPrice, _ := money.New(5000, "USD") // $50.00
	// Apply 20% discount
	discounted, _ := originalPrice.ApplyPercentageDiscount(20.0)
	fmt.Printf("20%% discount on %s = %s\n",
		originalPrice.Format(), discounted.Format())

	// Comparison Operations
	fmt.Println("\n=== Comparison Operations ===")

	amount1, _ := money.New(1000, "USD") // $10.00
	amount2, _ := money.New(2000, "USD") // $20.00
	amount3, _ := money.New(1000, "USD") // $10.00

	isLess, _ := amount1.LessThan(amount2)
	isGreater, _ := amount2.GreaterThan(amount1)
	isEqual, _ := amount1.Equals(amount3)

	fmt.Printf("%s < %s: %t\n", amount1.Format(), amount2.Format(), isLess)
	fmt.Printf("%s > %s: %t\n", amount2.Format(), amount1.Format(), isGreater)
	fmt.Printf("%s = %s: %t\n", amount1.Format(), amount3.Format(), isEqual)

	// Sign Operations
	fmt.Println("\n=== Sign Operations ===")

	profit, _ := money.New(1500, "USD") // $15.00
	loss, _ := money.New(-500, "USD")   // -$5.00
	zero, _ := money.New(0, "USD")      // $0.00

	fmt.Printf("%s is positive: %t\n", profit.Format(), profit.IsPositive())
	fmt.Printf("%s is negative: %t\n", loss.Format(), loss.IsNegative())
	fmt.Printf("%s is zero: %t\n", zero.Format(), zero.IsZero())
	fmt.Printf("Sign of %s: %d\n", loss.Format(), loss.Sign())
	fmt.Printf("Absolute value of %s: %s\n", loss.Format(), loss.Abs().Format())

	// Collection Operations
	fmt.Println("\n=== Collection Operations ===")

	// Create a slice of money values
	payments := money.MoneySlice{
		func() *money.Money { m, _ := money.New(1000, "USD"); return m }(),
		func() *money.Money { m, _ := money.New(2000, "USD"); return m }(),
		func() *money.Money { m, _ := money.New(-500, "USD"); return m }(),
		func() *money.Money { m, _ := money.New(1500, "USD"); return m }(),
	}

	// Sum
	total, _ = money.Sum(payments)
	fmt.Printf("Sum of payments: %s\n", total.Format())

	// Average
	avg, _ := money.Average(payments)
	fmt.Printf("Average payment: %s\n", avg.Format())

	// Sort ascending
	money.SortMoneySlice(payments)
	fmt.Print("Sorted payments (ascending): ")
	for _, p := range payments {
		fmt.Printf("%s ", p.Format())
	}
	fmt.Println()

	// Sort descending
	money.SortMoneySliceDescending(payments)
	fmt.Print("Sorted payments (descending): ")
	for _, p := range payments {
		fmt.Printf("%s ", p.Format())
	}
	fmt.Println()

	// Filter positive values
	positivePayments := money.Filter(payments, money.IsPositivePredicate)
	fmt.Print("Positive payments only: ")
	for _, p := range positivePayments {
		fmt.Printf("%s ", p.Format())
	}
	fmt.Println()

	// Map (double all values)
	doubled, _ := money.Map(payments, func(m *money.Money) *money.Money {
		return m.Multiply(2.0)
	})
	fmt.Print("Doubled payments: ")
	for _, p := range doubled {
		fmt.Printf("%s ", p.Format())
	}
	fmt.Println()

	// Error Handling Examples
	fmt.Println("\n=== Error Handling Examples ===")

	// Currency mismatch
	eur, _ := money.New(1000, "EUR")
	_, err := amount1.Add(eur)
	fmt.Printf("Adding different currencies: %v\n", err)

	// Empty slice
	_, err = money.Sum(money.MoneySlice{})
	fmt.Printf("Sum of empty slice: %v\n", err)

	// Invalid currency
	_, err = money.New(1000, "INVALID")
	fmt.Printf("Invalid currency: %v\n", err)

	// Invalid discount percentage
	_, err = amount1.ApplyPercentageDiscount(150.0)
	fmt.Printf("Invalid discount percentage: %v\n", err)
}
