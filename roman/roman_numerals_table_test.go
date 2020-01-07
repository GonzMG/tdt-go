package roman

import "testing"

var testCases = []struct {
	// Atributes of the test
	inputDecimal int
	outputRoman  string
}{
	// Input and expected values of the test cases
	{1, "I"},
	{4, "IV"},
	{6, "VI"},
	{125, "CXXV"},
	{47, "XLVII"},
	{492, "CDXCII"},
}

func TestRomanToNumberTable(t *testing.T) {
	for _, testCase := range testCases {
		result, _ := ConvertDecimalToRoman(testCase.inputDecimal)
		if result != testCase.outputRoman {
			t.Errorf("ERROR -> Input: %v | Output: %v | Expected: %v",
				testCase.inputDecimal, result, testCase.outputRoman)
		}
	}
}
