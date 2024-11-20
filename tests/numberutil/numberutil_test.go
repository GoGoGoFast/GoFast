package numberutil_test

import (
	"GoAllInOne/pkg/util/numberutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 3.0, numberutil.Add(1, 2))
}

func TestSub(t *testing.T) {
	assert.Equal(t, -1.0, numberutil.Sub(1, 2))
}

func TestMul(t *testing.T) {
	assert.Equal(t, 2.0, numberutil.Mul(1, 2))
}

func TestDiv(t *testing.T) {
	assert.Equal(t, 0.5, numberutil.Div(1, 2, 2, ""))
	assert.Equal(t, 0.52, numberutil.Div(1, 1.95, 2, "up"))
	assert.Equal(t, 0.51, numberutil.Div(1, 1.95, 2, ""))
	assert.Equal(t, 0.51, numberutil.Div(1, 1.95, 2, "down"))
}

func TestRound(t *testing.T) {
	assert.Equal(t, 1.23, numberutil.Round(1.234, 2))
}

func TestRoundStr(t *testing.T) {
	assert.Equal(t, "1.23", numberutil.RoundStr(1.234, 2))
}

func TestDecimalFormat(t *testing.T) {
	assert.Equal(t, "00123", numberutil.DecimalFormat("00000", 123))
}

func TestIsNumber(t *testing.T) {
	assert.True(t, numberutil.IsNumber("123.45"))
	assert.False(t, numberutil.IsNumber("abc"))
}

func TestIsInteger(t *testing.T) {
	assert.True(t, numberutil.IsInteger("123"))
	assert.False(t, numberutil.IsInteger("123.45"))
}

func TestIsDouble(t *testing.T) {
	assert.True(t, numberutil.IsDouble("123.45"))
	assert.False(t, numberutil.IsDouble("abc"))
}

func TestIsPrime(t *testing.T) {
	assert.True(t, numberutil.IsPrime(7))
	assert.False(t, numberutil.IsPrime(4))
}

func TestGenerateRandomNumber(t *testing.T) {
	numbers := numberutil.GenerateRandomNumber(1, 10, 5)
	assert.Equal(t, 5, len(numbers))
}

func TestFactorial(t *testing.T) {
	assert.Equal(t, 120, numberutil.Factorial(5))
}

func TestSqrt(t *testing.T) {
	assert.Equal(t, 2.0, numberutil.Sqrt(4))
}

func TestDivisor(t *testing.T) {
	assert.Equal(t, 2, numberutil.Divisor(4, 2))
}

func TestMultiple(t *testing.T) {
	assert.Equal(t, 6, numberutil.Multiple(2, 3))
}

func TestGetBinaryStr(t *testing.T) {
	assert.Equal(t, "1010", numberutil.GetBinaryStr(10))
}

func TestBinaryToInt(t *testing.T) {
	result, err := numberutil.BinaryToInt("1010")
	assert.NoError(t, err)
	assert.Equal(t, 10, result)
}

func TestCompare(t *testing.T) {
	assert.Equal(t, -1, numberutil.Compare(1, 2))
	assert.Equal(t, 0, numberutil.Compare(2, 2))
	assert.Equal(t, 1, numberutil.Compare(3, 2))
}

func TestToStr(t *testing.T) {
	assert.Equal(t, "123.45", numberutil.ToStr(123.450000))
}

func TestIsEven(t *testing.T) {
	assert.True(t, numberutil.IsEven(2))
	assert.False(t, numberutil.IsEven(3))
}

func TestIsOdd(t *testing.T) {
	assert.True(t, numberutil.IsOdd(3))
	assert.False(t, numberutil.IsOdd(2))
}

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 0, numberutil.Fibonacci(0))
	assert.Equal(t, 1, numberutil.Fibonacci(1))
	assert.Equal(t, 1, numberutil.Fibonacci(2))
	assert.Equal(t, 2, numberutil.Fibonacci(3))
	assert.Equal(t, 3, numberutil.Fibonacci(4))
	assert.Equal(t, 5, numberutil.Fibonacci(5))
}
