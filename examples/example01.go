package main

import (
	"fmt"

	"github.com/ha1tch/money" // assuming this is your package name
)

func main() {
	// Example 1: Creating money values
	dollars, _ := money.New(1099, "USD")         // $10.99
	euros, _ := money.NewFromFloat(20.50, "EUR") // €20.50
	fmt.Println("Formatted amounts:", dollars.Format(), euros.Format())

	// Example 2: Basic arithmetic
	price, _ := money.New(500, "USD") // $5.00
	tax, _ := money.New(45, "USD")    // $0.45
	total, _ := price.Add(tax)
	fmt.Println("Total with tax:", total.Format()) // $5.45

	// Example 3: Applying a discount
	originalPrice, _ := money.New(2500, "USD")                   // $25.00
	discounted, _ := originalPrice.ApplyPercentageDiscount(20.0) // 20% off
	fmt.Println("After 20% discount:", discounted.Format())

	// Example 4: Comparison operations
	amount1, _ := money.New(1000, "USD") // $10.00
	amount2, _ := money.New(2000, "USD") // $20.00
	isLess, _ := amount1.LessThan(amount2)
	fmt.Println("Is $10 less than $20?", isLess)

	// Example 5: Currency formatting options
	yen, _ := money.New(1000, "JPY") // ¥1,000 (JPY has 0 decimal places)
	fmt.Println("Default yen format:", yen.Format())
	fmt.Println("Custom yen format:", yen.FormatWithOptions(money.MoneyFormatOptions{
		UseSymbol:      true,
		ShowCents:      false,
		SymbolPosition: "before",
		GroupSeparator: ",",
	}))

	// Example 6: Sign operations
	profit, _ := money.New(1500, "USD") // $15.00
	loss, _ := money.New(-500, "USD")   // -$5.00
	fmt.Println("Is profit positive?", profit.IsPositive())
	fmt.Println("Is loss negative?", loss.IsNegative())

	// Example 7: Multiplication
	price, _ = money.New(399, "USD") // $3.99
	quantity := 3.0
	total = price.Multiply(quantity)
	fmt.Println("Total for 3 items:", total.Format())

	// Example 8: Absolute value
	negative, _ := money.New(-750, "USD") // -$7.50
	absolute := negative.Abs()
	fmt.Println("Absolute value:", absolute.Format())

	// Example 9: Brazilian Real with special rounding
	real, _ := money.New(123, "BRL") // R$1.23 -> will round to R$1.25
	fmt.Println("Brazilian Real with rounding:", real.Format())

	// Example 10: Zero checks
	zero, _ := money.New(0, "USD")
	amount, _ := money.New(100, "USD")
	fmt.Println("Is zero?", zero.IsZero())
	fmt.Println("Is positive?", amount.IsZero())
}
