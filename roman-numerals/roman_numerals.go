package romannumerals

import (
	"errors"
)

type romanConversion struct {
	roman  string
	arabic int
}

var romanConversionTable []romanConversion = []romanConversion{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// ToRomanNumeral converts the given number into roman within 0 to 3000
func ToRomanNumeral(num int) (string, error) {
	if num > 3000 || num <= 0 {
		return "", errors.New("number beyond range")
	}
	roman := ""
	// loop through the conversion table
	for _, conversion := range romanConversionTable {
		// decrease the highest number possible, as many times as possible
		for num >= conversion.arabic {
			roman += conversion.roman
			num -= conversion.arabic
		}
		if num == 0 {
			break
		}
	}
	return roman, nil
}
