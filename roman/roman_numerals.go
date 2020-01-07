package roman

import "errors"

var errorNegativeNumber error = errors.New("The input can not be lower than 0")

var decimalNumerals []int = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
var romanNumerals []string = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC",
	"C", "CD", "D", "CM", "M"}

func ConvertDecimalToRoman(decimalNumber int) (string, error) {
	romanNumber := ""
	lastIndex := 0
	if decimalNumber < 0 {
		return "", errorNegativeNumber
	}
	for index := range decimalNumerals {
		if decimalNumber == decimalNumerals[index] {
			romanNumber = romanNumerals[index]
			return romanNumber, nil
		}
		if decimalNumerals[index] < decimalNumber {
			romanNumber = romanNumerals[index]
			lastIndex = index
		}
	}
	remainningDecimal := decimalNumber - decimalNumerals[lastIndex]
	result, _ := ConvertDecimalToRoman(remainningDecimal)
	romanNumber += result
	return romanNumber, nil
}
