package roman

import "testing"

func TestRomanToNumberSimple(t *testing.T) {

	input := []int{1, 4, 6, 125, 47, 492}
	expected := []string{"I", "IV", "VI", "CXXV", "XLVII", "CDXCII"}

	for i := range input {
		result, _ := ConvertDecimalToRoman(input[i])
		if result != expected[i] {
			t.Errorf("ERROR -> Input: %v | Output: %v | Expected: %v",
				input[i], result, expected[i])
		}
	}
}

func TestRomanToNumberSub(t *testing.T) {
	t.Run("Test positive number", func(t *testing.T) {
		result, _ := ConvertDecimalToRoman(14)
		if result != "XIV" {
			t.Errorf("ERROR -> Input: %v | Output: %v | Expected: %v",
				14, result, "XVI")
		}
	})
	t.Run("Test negative number", func(t *testing.T) {
		_, err := ConvertDecimalToRoman(-1)
		if err == nil {
			t.Error("The test does not fail")
		}
	})
}
