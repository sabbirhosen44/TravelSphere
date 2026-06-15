package utils

import (
	"fmt"
	"strings"
	"time"
)

// FormatPopulation formats an integer with comma separators (e.g., 1234567 -> "1,234,567").
func FormatPopulation(pop int64) string {
	if pop == 0 {
		return "0"
	}
	neg := false
	if pop < 0 {
		neg = true
		pop = -pop
	}

	s := fmt.Sprintf("%d", pop)
	var parts []string
	for len(s) > 3 {
		parts = append([]string{s[len(s)-3:]}, parts...)
		s = s[:len(s)-3]
	}
	if len(s) > 0 {
		parts = append([]string{s}, parts...)
	}

	result := strings.Join(parts, ",")
	if neg {
		result = "-" + result
	}
	return result
}

// FormatCurrencies takes the REST Countries currency map and returns a formatted string.
// E.g., {"USD": map[string]interface{}{"name": "United States dollar", "symbol": "$"}} -> "United States dollar ($)"
func FormatCurrencies(currencies map[string]interface{}) string {
	if len(currencies) == 0 {
		return "N/A"
	}
	var list []string
	for _, val := range currencies {
		currMap, ok := val.(map[string]interface{})
		if !ok {
			continue
		}
		name, _ := currMap["name"].(string)
		symbol, _ := currMap["symbol"].(string)
		if name != "" && symbol != "" {
			list = append(list, fmt.Sprintf("%s (%s)", name, symbol))
		} else if name != "" {
			list = append(list, name)
		} else if symbol != "" {
			list = append(list, symbol)
		}
	}
	if len(list) == 0 {
		return "N/A"
	}
	return strings.Join(list, ", ")
}

// FormatLanguages returns a comma-separated list of languages from the REST Countries languages map.
func FormatLanguages(languages map[string]string) string {
	if len(languages) == 0 {
		return "N/A"
	}
	var list []string
	for _, lang := range languages {
		if lang != "" {
			list = append(list, lang)
		}
	}
	if len(list) == 0 {
		return "N/A"
	}
	return strings.Join(list, ", ")
}

// FormatDate formats a time.Time as "YYYY-MM-DD HH:MM:SS".
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
