package money

// Helper function for rounding monetary values.
// The value parameter is the amount to round multiplied by 10 to handle an extra decimal place during calculations.
// For example, to round 1.234, pass 12340. The function will return 123 (representing 1.23).
//
// Available rounding methods:
//   - RoundHalfUp: Rounds up for fractions >= 0.5 (e.g., 1.25 -> 1.3, 1.24 -> 1.2)
//   - RoundHalfDown: Rounds down for fractions <= 0.5 (e.g., 1.25 -> 1.2, 1.26 -> 1.3)
//   - RoundHalfEven: Rounds to nearest even number for exact halves (e.g., 1.25 -> 1.2, 1.35 -> 1.4)
//   - RoundDown: Always rounds down (truncates) (e.g., 1.29 -> 1.2)
//   - RoundUp: Rounds up if there's any fraction (e.g., 1.21 -> 1.3)
//   - BrazilianRounding: Rounds to the nearest 0.05 for final amounts
//     Brazilian rounding examples:
//     1.21 -> 1.20
//     1.22 -> 1.20
//     1.23 -> 1.25
//     1.24 -> 1.25
//     1.26 -> 1.25
//     1.27 -> 1.25
//     1.28 -> 1.30
func round(value int64, method RoundingMethod) int64 {
	switch method {
	case RoundHalfUp:
		if value%10 >= 5 {
			return value/10 + 1
		}
		return value / 10

	case RoundHalfDown:
		if value%10 > 5 {
			return value/10 + 1
		}
		return value / 10

	case RoundHalfEven:
		mod := value % 10
		quotient := value / 10
		if mod > 5 || (mod == 5 && quotient%2 != 0) {
			return quotient + 1
		}
		return quotient

	case RoundDown:
		return value / 10

	case RoundUp:
		if value%10 != 0 {
			return value/10 + 1
		}
		return value / 10

	case BrazilianRounding:
		// Brazilian rounding works on the final amount (not intermediate calculations)
		// and rounds to the nearest 0.05
		// First, get the last digit (hundredths position)
		lastDigit := value % 10
		// Get the second to last digit (tenths position)
		secondLastDigit := (value / 10) % 10

		// Combine them to get the last two digits as a number between 0-99
		lastTwoDigits := secondLastDigit*10 + lastDigit

		// Round to nearest 5
		var roundedLastTwoDigits int64
		switch {
		case lastTwoDigits < 25:
			roundedLastTwoDigits = 20
		case lastTwoDigits < 75:
			roundedLastTwoDigits = 50
		default:
			roundedLastTwoDigits = 0
			value += 100 // Carry over to next whole number
		}

		// Replace last two digits with rounded value
		return (value/100)*100 + roundedLastTwoDigits/10

	default:
		return value / 10
	}
}

// Additional helper function to support Brazilian rounding edge cases
// This function should be called when displaying final amounts in BRL
func formatBrazilianAmount(amount int64) int64 {
	if amount < 0 {
		return -formatBrazilianAmount(-amount)
	}

	// Round to nearest 0.05
	cents := amount % 100
	wholePart := amount - cents

	var roundedCents int64
	switch {
	case cents < 3:
		roundedCents = 0
	case cents < 8:
		roundedCents = 5
	case cents < 13:
		roundedCents = 10
	case cents < 18:
		roundedCents = 15
	case cents < 23:
		roundedCents = 20
	case cents < 28:
		roundedCents = 25
	case cents < 33:
		roundedCents = 30
	case cents < 38:
		roundedCents = 35
	case cents < 43:
		roundedCents = 40
	case cents < 48:
		roundedCents = 45
	case cents < 53:
		roundedCents = 50
	case cents < 58:
		roundedCents = 55
	case cents < 63:
		roundedCents = 60
	case cents < 68:
		roundedCents = 65
	case cents < 73:
		roundedCents = 70
	case cents < 78:
		roundedCents = 75
	case cents < 83:
		roundedCents = 80
	case cents < 88:
		roundedCents = 85
	case cents < 93:
		roundedCents = 90
	case cents < 98:
		roundedCents = 95
	default:
		return wholePart + 100
	}

	return wholePart + roundedCents
}
