package roman

import "testing"

var (
	input    []string = []string{"I", "IV", "CXXV", "XLVII", "CDXCII"}
	expected []int    = []int{1, 4, 125, 47, 492}
)

func TestRomanToNumberSimple(t *testing.T) {
	for i := range input {
		result := ConvertRomanToDecimal(input[i])
		if result != expected[i] {
			t.Errorf("ERROR -> Input: %v | Output: %v | Expected: %v",
				input[i], result, expected[i])
		}
	}
}
