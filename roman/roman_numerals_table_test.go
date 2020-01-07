package roman

import "testing"

func TestRomanToNumberTable(t *testing.T) {
	testCases := []struct {
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

	for _, testCase := range testCases {
		result, _ := ConvertDecimalToRoman(testCase.inputDecimal)
		if result != testCase.outputRoman {
			t.Errorf("ERROR -> Input: %v | Output: %v | Expected: %v",
				testCase.inputDecimal, result, testCase.outputRoman)
		}
	}
}

func TestRomanToNumberSubTable(t *testing.T) {
	testCases := []struct {
		// NOTE: we are going to group the input/outputs into negative
		// and positive input cases.
		nameTest     string
		inputDecimal []int
		outputRoman  []string
		outputError  []error
	}{
		{"Testing positive numbers", []int{1, 125, 492}, []string{"I", "CXXV", "CDXCII"}, []error{nil, nil, nil}},
		{"Testing negative numbers", []int{-1, -125}, []string{"", ""}, []error{errorNegativeNumber, errorNegativeNumber}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.nameTest, func(t *testing.T) {
			for i := range testCase.inputDecimal {
				result, err := ConvertDecimalToRoman(testCase.inputDecimal[i])
				if result != testCase.outputRoman[i] || err != testCase.outputError[i] {
					t.Errorf("ERROR -> Input: %v | Output: %v | Expected: %v",
						testCase.inputDecimal[i], result, testCase.outputRoman[i])
				}
			}
		})
	}
}
