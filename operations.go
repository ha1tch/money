package money

import (
    "fmt"
    "math"
    "os"
)

// New creates a Money instance from an integer amount
func New(amount int64, currencyCode string) (*Money, error) {
    currency, err := GetCurrency(currencyCode)
    if err != nil {
        return nil, err
    }
    return &Money{amount: amount, currency: currency}, nil
}

// NewFromFloat creates a Money instance from a float64, applying DefaultRoundingMethod
func NewFromFloat(amount float64, currencyCode string) (*Money, error) {
    if WarnOnFloat64Constructor {
        fmt.Fprintln(os.Stderr, "Warning: Using float64 in money calculations may lead to precision issues.")
    }

    currency, err := GetCurrency(currencyCode)
    if err != nil {
        return nil, err
    }

    factor := math.Pow(10, float64(currency.Precision))
    scaledAmount := round(int64(amount*factor*10), DefaultRoundingMethod) / 10
    return &Money{amount: scaledAmount, currency: currency}, nil
}

// Add adds two Money instances
func (m *Money) Add(other *Money) (*Money, error) {
    if m.currency != other.currency {
        return nil, &CurrencyMismatchError{
            Currency1: m.currency.Code,
            Currency2: other.currency.Code,
        }
    }

    // Check for overflow
    if (other.amount > 0 && m.amount > math.MaxInt64-other.amount) ||
        (other.amount < 0 && m.amount < math.MinInt64-other.amount) {
        return nil, &OverflowError{
            Operation: "addition",
            Amount1:   m.amount,
            Amount2:   other.amount,
        }
    }

    return &Money{amount: m.amount + other.amount, currency: m.currency}, nil
}

// Subtract subtracts another Money from the current Money
func (m *Money) Subtract(other *Money) (*Money, error) {
    if m.currency != other.currency {
        return nil, &CurrencyMismatchError{
            Currency1: m.currency.Code,
            Currency2: other.currency.Code,
        }
    }

    // Check for overflow
    if (other.amount < 0 && m.amount > math.MaxInt64+other.amount) ||
        (other.amount > 0 && m.amount < math.MinInt64+other.amount) {
        return nil, &OverflowError{
            Operation: "subtraction",
            Amount1:   m.amount,
            Amount2:   other.amount,
        }
    }

    return &Money{amount: m.amount - other.amount, currency: m.currency}, nil
}

// Multiply multiplies Money by a factor and rounds the result
func (m *Money) Multiply(factor float64) *Money {
    scaledAmount := round(int64(float64(m.amount)*factor*10), DefaultRoundingMethod) / 10
    return &Money{amount: scaledAmount, currency: m.currency}
}

// ApplyPercentageDiscount applies a percentage discount to Money
func (m *Money) ApplyPercentageDiscount(percentage float64) (*Money, error) {
    if percentage < 0 || percentage > 100 {
        return nil, &ValidationError{
            Field:   "percentage",
            Message: "percentage must be between 0 and 100",
        }
    }
    discountAmount := m.Multiply(percentage / 100)
    return m.Subtract(discountAmount)
}

// Equals checks if two Money instances are equal
func (m *Money) Equals(other *Money) (bool, error) {
    if m.currency != other.currency {
        return false, &CurrencyMismatchError{
            Currency1: m.currency.Code,
            Currency2: other.currency.Code,
        }
    }
    return m.amount == other.amount, nil
}

// GreaterThan checks if this Money is greater than another
func (m *Money) GreaterThan(other *Money) (bool, error) {
    if m.currency != other.currency {
        return false, &CurrencyMismatchError{
            Currency1: m.currency.Code,
            Currency2: other.currency.Code,
        }
    }
    return m.amount > other.amount, nil
}

// LessThan checks if this Money is less than another
func (m *Money) LessThan(other *Money) (bool, error) {
    if m.currency != other.currency {
        return false, &CurrencyMismatchError{
            Currency1: m.currency.Code,
            Currency2: other.currency.Code,
        }
    }
    return m.amount < other.amount, nil
}

// Abs returns the absolute value of Money
func (m *Money) Abs() *Money {
    if m.amount < 0 {
        return &Money{
            amount:   -m.amount,
            currency: m.currency,
        }
    }
    return &Money{
        amount:   m.amount,
        currency: m.currency,
    }
}

// Sign returns:
// -1 if the amount is negative
// 0 if the amount is zero
// 1 if the amount is positive
func (m *Money) Sign() int {
    if m.amount < 0 {
        return -1
    }
    if m.amount > 0 {
        return 1
    }
    return 0
}

// IsZero returns true if the amount is zero
func (m *Money) IsZero() bool {
    return m.amount == 0
}

// IsPositive returns true if the amount is greater than zero
func (m *Money) IsPositive() bool {
    return m.amount > 0
}

// IsNegative returns true if the amount is less than zero
func (m *Money) IsNegative() bool {
    return m.amount < 0
}
