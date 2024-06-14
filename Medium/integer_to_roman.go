package main

// romanValues maps Roman numeral symbols to their values.
var romanValues = []struct {
	value  int
	symbol string
}{
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

// intToRoman converts an integer to a Roman numeral.
func intToRoman(num int) string {
	result := ""

	// Iterate over the romanValues from largest to smallest.
	for _, rv := range romanValues {
		// While num is greater than or equal to the current value,
		// append the corresponding symbol to the result and subtract the value from num.
		for num >= rv.value {
			result += rv.symbol
			num -= rv.value
		}
	}

	return result
}
