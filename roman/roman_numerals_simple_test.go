package roman

import "testing"

var (
	input    []int    = []int{1, 4, 6, 125, 47, 492}
	expected []string = []string{"I", "IV", "VI", "CXXV", "XLVII", "CDXCII"}
)

func TestRomanToNumberSimple(t *testing.T) {
	for i := range input {
		result := ConvertDecimalToRoman(input[i])
		if result != expected[i] {
			t.Errorf("ERROR -> Input: %v | Output: %v | Expected: %v",
				input[i], result, expected[i])
		}
	}
}
