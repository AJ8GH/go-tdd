package propertybasedtests

import "strings"

var romanToArabic = []struct {
	roman  string
	arabic int
}{
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

func ToRoman(n int) string {
	var roman strings.Builder
	for _, v := range romanToArabic {
		for n >= v.arabic {
			roman.WriteString(v.roman)
			n -= v.arabic
		}
	}
	return roman.String()
}
