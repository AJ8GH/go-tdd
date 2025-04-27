package propertybasedtests

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		arabic int
		want   string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{12, "XII"},
		{13, "XIII"},
		{14, "XIV"},
		{15, "XV"},
		{19, "XIX"},
		{20, "XX"},
		{39, "XXXIX"},
		{40, "XL"},
		{48, "XLVIII"},
		{49, "XLIX"},
		{50, "L"},
		{50, "L"},
		{90, "XC"},
		{98, "XCVIII"},
		{99, "XCIX"},
		{100, "C"},
		{400, "CD"},
		{500, "D"},
		{798, "DCCXCVIII"},
		{900, "CM"},
		{1000, "M"},
		{1006, "MVI"},
		{1984, "MCMLXXXIV"},
		{2014, "MMXIV"},
		{3999, "MMMCMXCIX"},
	}

	for _, test := range cases {
		got := ToRoman(test.arabic)
		if got != test.want {
			t.Run(test.want, func(t *testing.T) {
				t.Errorf(`got "%s", want "%s"`, got, test.want)
			})
		}
	}
}
