package expression_calculator
import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	testTable := []struct {
		str           string
		expected      int
		expectedError error
	}{
		{
			str:           "1 +  3  - 2",
			expected:      2,
			expectedError: nil,
		},
		{
			str:           "        +   2   -     1",
			expected:      1,
			expectedError: nil,
		},
		{
			str:           "++2",
			expected:      0,
			expectedError: fmt.Errorf("unexpected token at pos: %v", 1),
		},
	}
	for _, testcase := range testTable {
		result, err := Evaluate(testcase.str)
		t.Logf("сalling str : %s, error : %v, result : %d\n", testcase.str, err, result)
		//Assert
		assert.Equal(t, testcase.expectedError, err,
			fmt.Sprintf("incorrect error : expected : %v, got : %v", testcase.expectedError, err))
		assert.Equal(t, testcase.expected, result,
			fmt.Sprintf("incorrect result. Expected %d. Got %d", testcase.expected, result))

	}
}
func benchExpr(expr string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Evaluate(expr)
	}
}
func BenchmarkExpression_Evaluate3(b *testing.B) {
	benchExpr("1+1", b)
}
func BenchmarkExpression_Evaluate10(b *testing.B) {
	benchExpr(" 2+  1 - 2", b)
}
func BenchmarkExpression_Evaluate30(b *testing.B) {
	benchExpr("7+1 +  6  - 7 - 1 +3 + 1 + 1+2", b)
}

func ExampleMain() {
	strings := [3]string{
		"1 +  3  - 2",
		"        +   2   -     1",
		"++2",
	}
	for _, value := range strings {
		fmt.Println(Evaluate(value))
	}
	// Output:
	// 2 <nil>
	//1 <nil>
	// 0 unexpected token at pos: 1
}