package money

import "time"

// RoundingMethod defines available rounding methods
type RoundingMethod int

const (
    RoundHalfUp RoundingMethod = iota
    RoundHalfDown
    RoundHalfEven
    RoundUp
    RoundDown
    BrazilianRounding
)

// DefaultRoundingMethod sets the default rounding for operations
var DefaultRoundingMethod = RoundHalfUp

// CurrencyConverter defines the interface for getting exchange rates
type CurrencyConverter interface {
    GetRate(fromCurrency, toCurrency string, date *time.Time) (float64, error)
}

// DefaultConverter can be set by the application to handle currency conversions
var DefaultConverter CurrencyConverter

// WarnOnFloat64Constructor controls whether a warning appears when using float64 in constructors
var WarnOnFloat64Constructor = true

// Currency holds details about each currency
type Currency struct {
    Code             string
    Symbol           string
    Precision        int
    SingularName     string
    PluralName       string
    GroupSeparator   string
    DecimalSeparator string
    SymbolPosition   string // "before" or "after"
}

// CurrencyMap defines available currencies
// Currency list organized by regions and financial importance
var CurrencyMap = map[string]Currency{
    // Major World Currencies
    "USD": {"USD", "$", 2, "Dollar", "Dollars", ",", ".", "before"},      // US Dollar
    "EUR": {"EUR", "€", 2, "Euro", "Euros", ".", ",", "before"},          // Euro
    "JPY": {"JPY", "¥", 0, "Yen", "Yen", ",", "", "before"},             // Japanese Yen
    "GBP": {"GBP", "£", 2, "Pound", "Pounds", ",", ".", "before"},       // British Pound Sterling
    "CHF": {"CHF", "Fr.", 2, "Franc", "Francs", "'", ".", "before"},     // Swiss Franc

    // South American Currencies
    "BRL": {"BRL", "R$", 2, "Real", "Reais", ".", ",", "before"},        // Brazilian Real
    "ARS": {"ARS", "$", 0, "Peso", "Pesos", ".", ",", "before"},         // Argentine Peso
    "UYU": {"UYU", "$U", 2, "Peso", "Pesos", ".", ",", "before"},        // Uruguayan Peso
    "CLP": {"CLP", "$", 0, "Peso", "Pesos", ".", "", "before"},          // Chilean Peso
    
    // North American Currencies
    "CAD": {"CAD", "$", 2, "Dollar", "Dollars", ",", ".", "before"},     // Canadian Dollar
    "MXN": {"MXN", "$", 2, "Peso", "Pesos", ",", ".", "before"},        // Mexican Peso

    // Asia-Pacific Currencies
    "CNY": {"CNY", "¥", 2, "Yuan", "Yuan", ",", ".", "before"},          // Chinese Yuan (Renminbi)
    "HKD": {"HKD", "HK$", 2, "Dollar", "Dollars", ",", ".", "before"},   // Hong Kong Dollar
    "SGD": {"SGD", "$", 2, "Dollar", "Dollars", ",", ".", "before"},     // Singapore Dollar
    "INR": {"INR", "₹", 2, "Rupee", "Rupees", ",", ".", "before"},      // Indian Rupee
    "KRW": {"KRW", "₩", 0, "Won", "Won", ",", "", "before"},            // South Korean Won
    "TWD": {"TWD", "NT$", 2, "Dollar", "Dollars", ",", ".", "before"},   // New Taiwan Dollar
    "AUD": {"AUD", "$", 2, "Dollar", "Dollars", ",", ".", "before"},     // Australian Dollar
    "NZD": {"NZD", "$", 2, "Dollar", "Dollars", ",", ".", "before"},     // New Zealand Dollar
    
    // European Currencies (non-EUR)
    "SEK": {"SEK", "kr", 2, "Krona", "Kronor", " ", ",", "after"},      // Swedish Krona
    "NOK": {"NOK", "kr", 2, "Krone", "Kroner", " ", ",", "after"},      // Norwegian Krone
    "DKK": {"DKK", "kr", 2, "Krone", "Kroner", ".", ",", "after"},      // Danish Krone
}

// Money represents a monetary value in the smallest unit (e.g., cents)
type Money struct {
    amount   int64
    currency Currency
}

// MoneyFormatOptions defines how a Money instance is formatted
type MoneyFormatOptions struct {
    UseSymbol        bool   // Display the currency symbol or code
    ShowCents        bool   // Show decimal places even if 0
    SymbolPosition   string // "before" or "after" the amount
    GroupSeparator   string // Separator for thousands grouping
    DecimalSeparator string // Separator for decimal places
}
