package romannumerals

type tuple struct {
	value int
	digit string
}

func Encode(n int) (string, bool) {
	conversions := []tuple{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var roman string
	if n > 0 {
		for _, conversion := range conversions {
			for n >= conversion.value {
				roman += conversion.digit
				n -= conversion.value
			}
		}
		return roman, true
	}
	return "", false
}

func Decode(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	canConvert := true
	conversions := []tuple{
		{1, "I"},
		{5, "V"},
		{10, "X"},
		{50, "L"},
		{100, "C"},
		{500, "D"},
		{1000, "M"},
	}
	var values []int
	var result int
	for i := range s {
		for j := range conversions {
			if conversions[j].digit == string(s[i]) {
				values = append(values, conversions[j].value)
			}
		}
		if len(values)-1 != i {
			canConvert = false
		}
	}
	if !canConvert {
		return 0, canConvert
	}

	for i := 0; i < len(values)-1; i++ {
		if values[i] < values[i+1] {
			result += values[i] * -1
			continue
		}
		result += values[i]
	}
	result += values[len(values)-1]

	return result, canConvert
}
