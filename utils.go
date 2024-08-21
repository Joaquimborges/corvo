package corvo

import (
	"strconv"
	"strings"
)

func parseBrazilianStrAmountToFloat(value string) (float64, error) {
	value = strings.ReplaceAll(value, ",", ".")
	return strconv.ParseFloat(value, 64)
}
