package roman

var decimalNumerals []int = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
var romanNumerals []string = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC",
	"C", "CD", "D", "CM", "M"}

func ConvertDecimalToRoman(decimalNumber int) string {
	romanNumber := ""
	lastIndex := 0
	for index := range decimalNumerals {
		if decimalNumber == decimalNumerals[index] {
			romanNumber = romanNumerals[index]
			return romanNumber
		}
		if decimalNumerals[index] < decimalNumber {
			romanNumber = romanNumerals[index]
			lastIndex = index
		}
	}
	remainningDecimal := decimalNumber - decimalNumerals[lastIndex]
	romanNumber += ConvertDecimalToRoman(remainningDecimal)
	return romanNumber
}
