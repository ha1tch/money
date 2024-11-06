package money

import (
    "fmt"
    "math"
    "strings"
    "time"
)

// GetCurrency retrieves the Currency struct from CurrencyMap
func GetCurrency(code string) (Currency, error) {
    if code == "" {
        return Currency{}, &ValidationError{
            Field:   "currency code",
            Message: "currency code cannot be empty",
        }
    }

    currency, exists := CurrencyMap[code]
    if !exists {
        return Currency{}, &GetCurrencyError{
            Code: code,
        }
    }
    return currency, nil
}

// ConvertTo converts Money to another currency given an exchange rate
func (m *Money) ConvertTo(targetCurrency string, rate float64) (*Money, error) {
    targetCurrencyObj, err := GetCurrency(targetCurrency)
    if err != nil {
        return nil, err
    }

    // Special handling for Brazilian Real conversions
    if targetCurrencyObj.Code == "BRL" {
        targetAmount := round(int64(float64(m.amount)*rate*10), DefaultRoundingMethod) / 10
        // Apply Brazilian rounding to the final amount
        targetAmount = formatBrazilianAmount(targetAmount)
        return &Money{amount: targetAmount, currency: targetCurrencyObj}, nil
    }

    targetAmount := round(int64(float64(m.amount)*rate*10), DefaultRoundingMethod) / 10
    return &Money{amount: targetAmount, currency: targetCurrencyObj}, nil
}

// ConvertViaReference converts to another currency using a reference currency and optional date.
// It uses the DefaultConverter to get exchange rates. Returns an error if no converter is configured
// or if any conversion fails.
func (m *Money) ConvertViaReference(targetCurrency, referenceCurrency string, date *time.Time) (*Money, error) {
    if DefaultConverter == nil {
        return nil, &ValidationError{
            Field:   "converter",
            Message: "no currency converter configured",
        }
    }

    toReferenceRate, err := DefaultConverter.GetRate(m.currency.Code, referenceCurrency, date)
    if err != nil {
        return nil, &ValidationError{
            Field:   "exchange_rate",
            Message: fmt.Sprintf("could not fetch rate from %s to %s: %v", m.currency.Code, referenceCurrency, err),
        }
    }

    referenceAmount := round(int64(float64(m.amount)*toReferenceRate*10), DefaultRoundingMethod) / 10
    fromReferenceRate, err := DefaultConverter.GetRate(referenceCurrency, targetCurrency, date)
    if err != nil {
        return nil, &ValidationError{
            Field:   "exchange_rate",
            Message: fmt.Sprintf("could not fetch rate from %s to %s: %v", referenceCurrency, targetCurrency, err),
        }
    }

    targetCurrencyObj, err := GetCurrency(targetCurrency)
    if err != nil {
        return nil, err
    }

    targetAmount := round(int64(float64(referenceAmount)*fromReferenceRate*10), DefaultRoundingMethod) / 10
    
    // Apply Brazilian rounding if converting to BRL
    if targetCurrencyObj.Code == "BRL" {
        targetAmount = formatBrazilianAmount(targetAmount)
    }

    return &Money{amount: targetAmount, currency: targetCurrencyObj}, nil
}

// FormatWithOptions formats Money with custom options
func (m *Money) FormatWithOptions(opts MoneyFormatOptions) string {
    amount := m.amount

    // Apply Brazilian rounding for BRL currency on final display
    if m.currency.Code == "BRL" {
        amount = formatBrazilianAmount(amount)
    }

    units := amount / int64(math.Pow10(m.currency.Precision))
    decimals := amount % int64(math.Pow10(m.currency.Precision))
    if decimals < 0 {
        decimals = -decimals
    }

    // Handle negative amounts
    negative := amount < 0
    if negative {
        units = -units
    }

    var result string
    
    // Add negative sign if necessary
    if negative {
        result += "-"
    }

    // Add currency symbol if requested
    if opts.UseSymbol {
        if opts.SymbolPosition == "before" {
            result += m.currency.Symbol + " "
        }
    }

    // Format the main amount
    amountStr := fmt.Sprintf("%d", units)
    if opts.GroupSeparator != "" {
        amountStr = addThousandsSeparator(amountStr, opts.GroupSeparator)
    }
    result += amountStr

    // Add decimal places if needed
    if opts.ShowCents && m.currency.Precision > 0 {
        decimalPart := fmt.Sprintf("%0*d", m.currency.Precision, decimals)
        result += opts.DecimalSeparator + decimalPart
    }

    // Add currency symbol after if specified
    if opts.UseSymbol && opts.SymbolPosition == "after" {
        result += " " + m.currency.Symbol
    }

    return result
}

// Format returns a string representation using default formatting options for the currency
func (m *Money) Format() string {
    return m.FormatWithOptions(MoneyFormatOptions{
        UseSymbol:        true,
        ShowCents:        true,
        SymbolPosition:   m.currency.SymbolPosition,
        GroupSeparator:   m.currency.GroupSeparator,
        DecimalSeparator: m.currency.DecimalSeparator,
    })
}

// Helper for thousands separator
func addThousandsSeparator(s, sep string) string {
    var result strings.Builder
    n := len(s) % 3
    if n > 0 {
        result.WriteString(s[:n])
        if len(s) > n {
            result.WriteString(sep)
        }
    }
    for i := n; i < len(s); i += 3 {
        result.WriteString(s[i : i+3])
        if i+3 < len(s) {
            result.WriteString(sep)
        }
    }
    return result.String()
}
