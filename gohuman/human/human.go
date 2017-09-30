// Package human provides functions for converting a number to human readable form
package human

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Bytes transforms a number into a string with 1024 as base.
// Arguments: n - number to transform, decimals - number of decimals to show,
// long - if true, show full names, otherwise show abbreviated.
func Bytes(n int64, decimals int, long bool) string {
	padding := " "
	levels := []string{
		"B", "KB", "MB", "GB", "TB", "PB", "EB", /* "ZB", "YB" will overflow int64 */
	}
	if long {
		levels = []string{
			"bytes", "kilobyte", "megabyte", "gigabyte", "terabyte", "petabyte", "exabyte",
		}
	}

	return human(n, levels, 1024, decimals, padding)
}

// Kilos transforms a number into a string with 1000 as base
// Arguments: n - number to transform, decimals - number of decimals to show,
// long - if true, show full names, otherwise show abbreviated.
func Kilos(n int64, decimals int, long bool) string {
	padding := " "
	levels := []string{
		"", "k", "M", "G", "T", "P", "E", /* "Z", "Y" will overflow int64 */
	}
	if long {
		levels = []string{
			"s", "kilo", "mega", "giga", "tera", "peta", "exa",
		}
	}

	return human(n, levels, 1000, decimals, padding)
}

// human - generic number to string conversion function
func human(n int64, levels []string, base int, decimals int, padding string) string {

	format := fmt.Sprintf("%%.%df%v%%v", decimals, padding)
	res := strconv.FormatInt(n, 10)

	for i, level := range levels {
		nom := float64(n)
		den := math.Pow(float64(base), float64(i))
		val := nom / den
		if val < float64(base) {
			res = fmt.Sprintf(format, val, level)
			break
		}
	}

	return strings.TrimSpace(res)
}
