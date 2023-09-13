package fx

import (
	"fmt"
	"strings"
)

// Mapping from USD to other currencies
var usdRate map[string]float64 = map[string]float64{
	"USD": 1,
	"SGD": 1.36,
	"MYR": 4.68,
	"EUR": 0.93,
}

func GetRate(from string, to string) (float64, error) {
	from = strings.ToUpper(from)
	to = strings.ToUpper(to)

	fromRate, ok := usdRate[from]
	if !ok {
		return 0, fmt.Errorf("%s not found", from)
	}

	toRate, ok := usdRate[to]
	if !ok {
		return 0, fmt.Errorf("%s not found", to)
	}

	// from -> USD -> to
	return toRate / fromRate, nil
}

func Convert(from string, to string, amount float64) (float64, error) {
	rate, err := GetRate(from, to)
	if err != nil {
		return 0, err
	}

	return amount * rate, nil
}
