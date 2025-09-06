package utils

import (
	"fmt"
	"math"
)

func IsWhole(n float64) bool {
	return math.Mod(n, 1) == 0
}

func FormatFloatUtil(f float64) string {
	if f == math.Trunc(f) {
		return fmt.Sprintf("%d", int64(f)) // no decimal part
	}
	return fmt.Sprintf("%.3f", f) // keeps decimals
}
