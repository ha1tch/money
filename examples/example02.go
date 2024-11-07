package main

import (
	"fmt"
	"time"

	"github.com/ha1tch/money"
)

// Example currency converter implementation
type ExampleConverter struct{}

func (c *ExampleConverter) GetRate(from, to string, date *time.Time) (float64, error) {
	// Example fixed rates (in real usage, you'd fetch these from an API)
	rates := map[string]map[string]float64{
		"USD": {
			"EUR": 0.85,
			"GBP": 0.73,
			"JPY": 110.0,
		},
		"EUR": {
			"USD": 1.18,
			"GBP": 0.86,
			"JPY": 129.5,
		},
	}

	if fromRates, ok := rates[from]; ok {
		if rate, ok := fromRates[to]; ok {
			return rate, nil
		}
	}
	return 0, fmt.Errorf("rate not found for %s to %s", from, to)
}

func main() {
	// Set up the converter
	money.DefaultConverter = &ExampleConverter{}

	// Example 1: Create amounts in different currencies
	usd, _ := money.NewFromFloat(100.00, "USD")
	eur, _ := money.NewFromFloat(100.00, "EUR")
	gbp, _ := money.NewFromFloat(100.00, "GBP")
	jpy, _ := money.NewFromFloat(10000, "JPY")

	fmt.Println("Original amounts:")
	fmt.Printf("USD: %s\n", usd.Format())
	fmt.Printf("EUR: %s\n", eur.Format())
	fmt.Printf("GBP: %s\n", gbp.Format())
	fmt.Printf("JPY: %s\n", jpy.Format())

	// Example 2: Direct currency conversion
	eurToUsd, _ := eur.ConvertTo("USD", 1.18)
	fmt.Printf("\nEUR 100 to USD: %s\n", eurToUsd.Format())

	// Example 3: Currency conversion via reference currency (USD)
	now := time.Now()
	eurToJpy, _ := eur.ConvertViaReference("JPY", "USD", &now)
	fmt.Printf("EUR 100 to JPY via USD: %s\n", eurToJpy.Format())

	// Example 4: Different formatting styles for the same amount
	amount, _ := money.NewFromFloat(1234567.89, "USD")

	// Default format
	fmt.Printf("\nDefault format: %s\n", amount.Format())

	// European style format
	europeanStyle := money.MoneyFormatOptions{
		UseSymbol:        true,
		ShowCents:        true,
		SymbolPosition:   "after",
		GroupSeparator:   ".",
		DecimalSeparator: ",",
	}
	fmt.Printf("European style: %s\n", amount.FormatWithOptions(europeanStyle))

	// Minimal format (no symbol, no grouping)
	minimalStyle := money.MoneyFormatOptions{
		UseSymbol:        false,
		ShowCents:        true,
		GroupSeparator:   "",
		DecimalSeparator: ".",
	}
	fmt.Printf("Minimal style: %s\n", amount.FormatWithOptions(minimalStyle))

	// Example 5: Currency validation
	// Try to add different currencies (will fail)
	_, err := usd.Add(eur)
	if err != nil {
		fmt.Printf("\nCurrency mismatch error: %v\n", err)
	}

	// Example 6: Brazilian Real rounding
	brl, _ := money.NewFromFloat(10.23, "BRL") // Will round to 10.25
	fmt.Printf("\nBrazilian Real with rounding: %s\n", brl.Format())

	// Example 7: Japanese Yen (no decimal places)
	yenAmount, _ := money.NewFromFloat(1234.56, "JPY") // Will round to 1235
	fmt.Printf("Japanese Yen (no decimals): %s\n", yenAmount.Format())

	// Example 8: Swiss Franc (special grouping)
	chf, _ := money.NewFromFloat(1234567.89, "CHF")
	fmt.Printf("Swiss Franc (with special grouping): %s\n", chf.Format())
}
