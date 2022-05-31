package currency

import "strings"

var currencySymbolMap = map[string]string{
	"CNY": "¥",
	"USD": "$",
}

func CurrencySymbolMap() map[string]string {
	return currencySymbolMap
}

func GetSymbol(currency string) string {
	return currencySymbolMap[strings.ToUpper(currency)]
}
