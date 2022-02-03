package conditional

import (
	"strconv"
)

func If(score int) string {
	var result string

	if score >= 85 {
		result = "A"
	} else if score >= 75 {
		result = "B"
	} else if score >= 65 {
		result = "C"
	} else if score >= 55 {
		result = "D"
	} else {
		result = "E"
	}

	return result
}

func Switch(number int) string {
	var result string

	switch number {
	case 1:
		result = "Satu"
	case 2:
		result = "Dua"
	case 3:
		result = "Tiga"
	default:
		s := strconv.Itoa(number)
		result = s
	}

	return result
}
