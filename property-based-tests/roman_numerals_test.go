package propertybasedtests

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		arabic uint16
		roman  string
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
		t.Run(fmt.Sprintf(`Convert "%d" to "%s`, test.arabic, test.roman), func(t *testing.T) {
			got := ToRoman(test.arabic)
			if got != test.roman {
				t.Errorf(`got "%s", want "%s"`, got, test.roman)
			}
		})

		t.Run(fmt.Sprintf(`Convert "%s" to "%d`, test.roman, test.arabic), func(t *testing.T) {
			got := ToArabic(test.roman)
			if got != test.arabic {
				t.Errorf(`got "%d", want "%d"`, got, test.arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ToRoman(arabic)
		fromRoman := ToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
