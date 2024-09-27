package numberutil

import (
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Add 进行加法运算
func Add(a, b float64) float64 {
	bfA := big.NewFloat(a)
	bfB := big.NewFloat(b)
	result := new(big.Float).Add(bfA, bfB)
	f64, _ := result.Float64()
	return f64
}

// Sub 进行减法运算
func Sub(a, b float64) float64 {
	bfA := big.NewFloat(a)
	bfB := big.NewFloat(b)
	result := new(big.Float).Sub(bfA, bfB)
	f64, _ := result.Float64()
	return f64
}

// Mul 进行乘法运算
func Mul(a, b float64) float64 {
	bfA := big.NewFloat(a)
	bfB := big.NewFloat(b)
	result := new(big.Float).Mul(bfA, bfB)
	f64, _ := result.Float64()
	return f64
}

// Div 进行除法运算，支持指定小数位数和舍弃方式
func Div(a, b float64, precision int, roundMode string) float64 {
	if b == 0 {
		return math.NaN() // 避免除以零
	}
	bfA := big.NewFloat(a)
	bfB := big.NewFloat(b)
	result := bfA.Quo(bfA, bfB)
	f64, _ := result.Float64()
	if precision > 0 {
		switch roundMode {
		case "up":
			return RoundUp(f64, precision)
		case "down":
			return RoundDown(f64, precision)
		default:
			return Round(f64, precision)
		}
	}
	return f64
}

// Round 四舍五入
func Round(value float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Round(value*factor) / factor
}

// RoundStr 四舍五入并返回字符串
func RoundStr(value float64, precision int) string {
	return strconv.FormatFloat(Round(value, precision), 'f', precision, 64)
}

// DecimalFormat 格式化数字
func DecimalFormat(format string, value int64) string {
	strValue := strconv.FormatInt(value, 10)
	var builder strings.Builder
	formatIndex, valueIndex := 0, 0

	for formatIndex < len(format) {
		ch := format[formatIndex]
		if ch == '#' || ch == '0' {
			if valueIndex < len(strValue) {
				builder.WriteByte(strValue[valueIndex])
				valueIndex++
			} else if ch == '0' {
				builder.WriteByte('0')
			}
		} else {
			builder.WriteByte(byte(ch))
		}
		formatIndex++
	}
	return builder.String()
}

// IsNumber 检查是否为数字
func IsNumber(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

// IsInteger 检查是否为整数
func IsInteger(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

// IsDouble 检查是否为浮点数
func IsDouble(value string) bool {
	return IsNumber(value)
}

// IsPrime 检查是否为质数
func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// GenerateRandomNumber 生成不重复随机数
func GenerateRandomNumber(min, max, count int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make(map[int]struct{})
	for len(numbers) < count {
		num := rand.Intn(max-min+1) + min
		numbers[num] = struct{}{}
	}
	result := make([]int, 0, len(numbers))
	for num := range numbers {
		result = append(result, num)
	}
	return result
}

// Factorial 计算阶乘
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Sqrt 计算平方根
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Divisor 计算最大公约数
func Divisor(a, b int) int {
	if b == 0 {
		return a
	}
	return Divisor(b, a%b)
}

// Multiple 计算最小公倍数
func Multiple(a, b int) int {
	return a * b / Divisor(a, b)
}

// GetBinaryStr 获取数字对应的二进制字符串
func GetBinaryStr(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

// BinaryToInt 二进制转 int
func BinaryToInt(s string) (int, error) {
	i64, err := strconv.ParseInt(s, 2, 0)
	return int(i64), err
}

// Compare 比较两个值的大小
func Compare(a, b float64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ToStr 数字转字符串，并自动去除尾部多余的 0
func ToStr(value float64) string {
	return strings.TrimRight(strings.TrimRight(strconv.FormatFloat(value, 'f', -1, 64), "0"), ".")
}

// 内部方法用于四舍五入、向上、向下取整
func round(value float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Round(value*factor) / factor
}

func RoundUp(value float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Ceil(value*factor) / factor
}

func RoundDown(value float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Floor(value*factor) / factor
}
