package money

import (
    "fmt"
    "math"
)

// CurrencyMismatchError represents an error when operations are attempted between different currencies
type CurrencyMismatchError struct {
    Currency1 string
    Currency2 string
}

func (e *CurrencyMismatchError) Error() string {
    return fmt.Sprintf("currency mismatch: %s vs %s", e.Currency1, e.Currency2)
}

// OverflowError represents an error when a monetary calculation would exceed int64 bounds
type OverflowError struct {
    Operation string
    Amount1   int64
    Amount2   int64
}

func (e *OverflowError) Error() string {
    return fmt.Sprintf("overflow detected in %s operation: %d, %d", e.Operation, e.Amount1, e.Amount2)
}

// ValidationError represents an error when input validation fails
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error for %s: %s", e.Field, e.Message)
}

// GetCurrencyError represents an error when a currency code is not found
type GetCurrencyError struct {
    Code string
}

func (e *GetCurrencyError) Error() string {
    return fmt.Sprintf("currency %s not found", e.Code)
}

// ValidateAmount checks if an amount is within safe bounds for monetary calculations
func ValidateAmount(amount int64) error {
    if amount < math.MinInt64/100 || amount > math.MaxInt64/100 {
        return &ValidationError{
            Field:   "amount",
            Message: "amount exceeds safe bounds for monetary calculations",
        }
    }
    return nil
}
