package money

import (
    "fmt"
    "sort"
)

// MoneySlice represents a slice of Money pointers that can be sorted
type MoneySlice []*Money

func (ms MoneySlice) Len() int { return len(ms) }

func (ms MoneySlice) Swap(i, j int) { ms[i], ms[j] = ms[j], ms[i] }

func (ms MoneySlice) Less(i, j int) bool {
    // Note: This assumes all Money values in the slice have the same currency
    // It's the caller's responsibility to ensure this
    return ms[i].amount < ms[j].amount
}

// ValidateMoneySlice checks if all Money values in a slice have the same currency
func ValidateMoneySlice(slice MoneySlice) error {
    if len(slice) <= 1 {
        return nil
    }

    currency := slice[0].currency
    for _, money := range slice[1:] {
        if money.currency != currency {
            return &CurrencyMismatchError{
                Currency1: currency.Code,
                Currency2: money.currency.Code,
            }
        }
    }
    return nil
}

// Sum returns the sum of a slice of Money values
func Sum(slice MoneySlice) (*Money, error) {
    if len(slice) == 0 {
        return nil, &ValidationError{
            Field:   "slice",
            Message: "empty slice provided",
        }
    }

    if err := ValidateMoneySlice(slice); err != nil {
        if cmErr, ok := err.(*CurrencyMismatchError); ok {
            return nil, cmErr
        }
        return nil, fmt.Errorf("sum validation failed: %w", err)
    }

    result := &Money{
        amount:   0,
        currency: slice[0].currency,
    }

    for _, money := range slice {
        var err error
        result, err = result.Add(money)
        if err != nil {
            return nil, fmt.Errorf("sum operation failed: %w", err)
        }
    }

    return result, nil
}

// Average returns the arithmetic mean of a slice of Money values
func Average(slice MoneySlice) (*Money, error) {
    if len(slice) == 0 {
        return nil, &ValidationError{
            Field:   "slice",
            Message: "empty slice provided",
        }
    }

    sum, err := Sum(slice)
    if err != nil {
        return nil, fmt.Errorf("average calculation failed: %w", err)
    }

    // Use Multiply to handle rounding according to DefaultRoundingMethod
    return sum.Multiply(1.0 / float64(len(slice))), nil
}

// SortMoneySlice sorts a slice of Money values in ascending order
// Returns error if currencies don't match
func SortMoneySlice(slice MoneySlice) error {
    if err := ValidateMoneySlice(slice); err != nil {
        if cmErr, ok := err.(*CurrencyMismatchError); ok {
            return cmErr
        }
        return fmt.Errorf("sort validation failed: %w", err)
    }
    sort.Sort(slice)
    return nil
}

// SortMoneySliceDescending sorts a slice of Money values in descending order
// Returns error if currencies don't match
func SortMoneySliceDescending(slice MoneySlice) error {
    if err := ValidateMoneySlice(slice); err != nil {
        if cmErr, ok := err.(*CurrencyMismatchError); ok {
            return cmErr
        }
        return fmt.Errorf("sort validation failed: %w", err)
    }
    sort.Sort(sort.Reverse(slice))
    return nil
}

// Filter returns a new MoneySlice containing only the elements that satisfy the predicate
func Filter(slice MoneySlice, predicate func(*Money) bool) MoneySlice {
    result := make(MoneySlice, 0, len(slice))
    for _, money := range slice {
        if predicate(money) {
            result = append(result, money)
        }
    }
    return result
}

// Example predicates that can be used with Filter
func IsPositivePredicate(m *Money) bool { return m.IsPositive() }
func IsNegativePredicate(m *Money) bool { return m.IsNegative() }
func IsZeroPredicate(m *Money) bool     { return m.IsZero() }

// Map applies a transformation function to each element in the slice
// Note: The transformation must maintain the same currency
func Map(slice MoneySlice, transform func(*Money) *Money) (MoneySlice, error) {
    if len(slice) == 0 {
        return MoneySlice{}, nil
    }

    result := make(MoneySlice, len(slice))
    for i, money := range slice {
        transformed := transform(money)
        if transformed.currency != money.currency {
            return nil, &ValidationError{
                Field:   "transform",
                Message: "transformation must maintain the same currency",
            }
        }
        result[i] = transformed
    }
    return result, nil
}
