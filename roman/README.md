# TDD with Table-Driven test in Go

One of the most important things developing software is the way of testing functionalities of the program and maintain clean and readable the code. This post will cover how to write useful tests keeping the code clean using table-driven tests, solving the Roman Numerals conversion to decimal numbers.  

Before start explaining this concept, an introduction to software testing is needed. A good Software Engineer must write tests that cover all the functionalities of the code and make scalable solutions. That's why Test Driven Development (TDD) is one of the most popular process, where the requirements are described as specific test cases and the software is improved passing each test. This method has a very short development cycle:

1. Write a failing test.
2. Write the simplest code to pass it.
3. Refactor the code.

Following this development process the Roman numerals problem will be solved. 

The Roman numerals problem is very used to explain TDD, it consist in a function that converts a decimal number into roman. To test functions in Go, the [package test](https://golang.org/pkg/testing/) will be used. Let's create a package called `roman` and add the file `roman_numerals.go` with the code of the convert function, that passing a decimal number it returns the roman conversion and an error if the decimal number is negative:

```
func ConvertDecimalToRoman(decimalNumber int) (string, error)
```

Following the TDD process, the first step is create a failing test, so create a file `roman_numerals_simple_test.go` with the test. The input and the expected values of the test will be:

```
var (
    input    []int    = []int{1, 4, 6, 125, 47, 492}

    expected []string = []string{"I", "IV", "VI", "CXXV", "XLVII", "CDXCII"}
)
````
The test function will iterate through the inputs, converting to roman and compare them with the expected values:

```
func TestRomanToNumberSimple(t *testing.T) {
    for i := range input {
        result, _ := ConvertDecimalToRoman(input[i])
        if result != expected[i] {
            t.Errorf("ERROR -> Input: %v | Output: %v | Expected: %v",
                input[i], result, expected[i])
        }
    }
}
```

If this test is executed with the command `go test -v`, the result should be:

```
=== RUN   TestRomanToNumberSimple
--- FAIL: TestRomanToNumberSimple (0.00s)
    roman_numerals_simple_test.go:13: ERROR -> Input: 1 | Output:  | Expected: I
    roman_numerals_simple_test.go:13: ERROR -> Input: 4 | Output:  | Expected: IV
    roman_numerals_simple_test.go:13: ERROR -> Input: 6 | Output:  | Expected: VI
    roman_numerals_simple_test.go:13: ERROR -> Input: 125 | Output:  | Expected: CXXV
    roman_numerals_simple_test.go:13: ERROR -> Input: 47 | Output:  | Expected: XLVII
    roman_numerals_simple_test.go:13: ERROR -> Input: 492 | Output:  | Expected: CDXCII
```

Now, writing the code to solve the test and executing it again, the result will be:

```
=== RUN   TestRomanToNumberSimple
--- PASS: TestRomanToNumberSimple (0.00s)
```

But this test does not cover if the input number is negative, so add a new test with this new test case should be added. In Go to avoid the creation of multiple test functions, a subtest can be created using `t.Run()`:

```
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
```

Executing this, the result will be:

```
=== RUN   TestRomanToNumberSub
=== RUN   TestRomanToNumberSub/Test_positive_number
=== RUN   TestRomanToNumberSub/Test_negative_number
--- PASS: TestRomanToNumberSub (0.00s)
    --- PASS: TestRomanToNumberSub/Test_positive_number (0.00s)
    --- PASS: TestRomanToNumberSub/Test_negative_number (0.00s)
```

This could be enough for testing our method, but the test file is not so clean if we continue adding test cases and new functionalities. Here appears the Table-Driven test solution.

Table-Driven test is a way of design tests that relies in the creation of structs with the input and expected values. This will make more readable the code because the test cases are more clear. Let's refactor the test TestRomaToNumberSub using this concept. For this function, the test case table will be:

```
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
```

So with this strategy it is only needed to iterate through testCases and use the attributes of each struct:

```
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
```

The result of this test will be:

```
=== RUN   TestRomanToNumberSub
=== RUN   TestRomanToNumberSub/Testing_positive_numbers
=== RUN   TestRomanToNumberSub/Testing_negative_numbers
--- PASS: TestRomanToNumberSub (0.00s)
    --- PASS: TestRomanToNumberSub/Testing_positive_numbers (0.00s)
    --- PASS: TestRomanToNumberSub/Testing_negative_numbers (0.00s)
```

With this approach of Table-Driven tests you will make better tests covering more cases and the design is smarter and clearer. I hope this post was helpful! In my [GitHub](https://github.com/GonzMG/tdt-go) you will find the whole code.

### References
1. [Table-Driven test](https://github.com/golang/go/wiki/TableDrivenTests).
2. [Going from 0 to 1 in Parallel test execution in Go](https://eleni.blog/2019/05/11/parallel-test-execution-in-go/) by Eleni Fragkiadaki.  
