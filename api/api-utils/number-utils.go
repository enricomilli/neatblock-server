package apiutil

import (
	"math"
	"strings"
	"time"
	"unicode"
)

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func UnixToDate(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format("2006-01-02")
}

func RemoveControlCharacters(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return -1
		}
		return r
	}, s)
}
