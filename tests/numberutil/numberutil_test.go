package numberutil_test

import (
	"VeloCore/pkg/util/numberutil"
	"testing"
)

func TestAdd(t *testing.T) {
	result := numberutil.Add(1.5, 2.3)
	expected := 3.8
	if result != expected {
		t.Errorf("Add(1.5, 2.3) = %f; want %f", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := numberutil.Sub(5.5, 2.2)
	expected := 3.3
	if result != expected {
		t.Errorf("Sub(5.5, 2.2) = %f; want %f", result, expected)
	}
}

func TestMul(t *testing.T) {
	result := numberutil.Mul(3.0, 2.5)
	expected := 7.5
	if result != expected {
		t.Errorf("Mul(3.0, 2.5) = %f; want %f", result, expected)
	}
}

func TestDiv(t *testing.T) {
	result := numberutil.Div(10.0, 4.0, 2, "")
	expected := 2.5
	if result != expected {
		t.Errorf("Div(10.0, 4.0, 2, '') = %f; want %f", result, expected)
	}
}

func TestRound(t *testing.T) {
	result := numberutil.Round(3.14159, 2)
	expected := 3.14
	if result != expected {
		t.Errorf("Round(3.14159, 2) = %f; want %f", result, expected)
	}
}

func TestDecimalFormat(t *testing.T) {
	result := numberutil.DecimalFormat("###,###", 123456)
	expected := "123,456"
	if result != expected {
		t.Errorf("DecimalFormat('###,###', 123456) = %s; want %s", result, expected)
	}
}

func TestIsNumber(t *testing.T) {
	result := numberutil.IsNumber("123.45")
	if !result {
		t.Errorf("IsNumber('123.45') = false; want true")
	}
}

func TestIsInteger(t *testing.T) {
	result := numberutil.IsInteger("123")
	if !result {
		t.Errorf("IsInteger('123') = false; want true")
	}
}

func TestIsPrime(t *testing.T) {
	result := numberutil.IsPrime(7)
	if !result {
		t.Errorf("IsPrime(7) = false; want true")
	}
}

func TestGenerateRandomNumber(t *testing.T) {
	numbers := numberutil.GenerateRandomNumber(1, 100, 10)
	if len(numbers) != 10 {
		t.Errorf("GenerateRandomNumber(1, 100, 10) generated %d numbers; want 10", len(numbers))
	}
}

func TestFactorial(t *testing.T) {
	result := numberutil.Factorial(5)
	expected := 120
	if result != expected {
		t.Errorf("Factorial(5) = %d; want %d", result, expected)
	}
}

func TestSqrt(t *testing.T) {
	result := numberutil.Sqrt(16.0)
	expected := 4.0
	if result != expected {
		t.Errorf("Sqrt(16.0) = %f; want %f", result, expected)
	}
}

func TestDivisor(t *testing.T) {
	result := numberutil.Divisor(36, 60)
	expected := 12
	if result != expected {
		t.Errorf("Divisor(36, 60) = %d; want %d", result, expected)
	}
}

func TestMultiple(t *testing.T) {
	result := numberutil.Multiple(12, 15)
	expected := 60
	if result != expected {
		t.Errorf("Multiple(12, 15) = %d; want %d", result, expected)
	}
}

func TestGetBinaryStr(t *testing.T) {
	result := numberutil.GetBinaryStr(10)
	expected := "1010"
	if result != expected {
		t.Errorf("GetBinaryStr(10) = %s; want %s", result, expected)
	}
}

func TestBinaryToInt(t *testing.T) {
	result, _ := numberutil.BinaryToInt("1010")
	expected := 10
	if result != expected {
		t.Errorf("BinaryToInt('1010') = %d; want %d", result, expected)
	}
}

func TestCompare(t *testing.T) {
	result := numberutil.Compare(3.0, 3.0)
	if result != 0 {
		t.Errorf("Compare(3.0, 3.0) = %d; want 0", result)
	}
}

func TestToStr(t *testing.T) {
	result := numberutil.ToStr(123.45000)
	expected := "123.45"
	if result != expected {
		t.Errorf("ToStr(123.45000) = %s; want %s", result, expected)
	}
}
