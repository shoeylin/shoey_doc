package problem0008

import (
	"math"
	"strings"
)

func myAtoi(str string) int {

	//除去 str 首尾的空格
	s := strings.TrimSpace(str)
	if len(s) == 0 {
		return 0
	}

	//取出 s 的符號和主體 x
	sign, x := getSign(s)

	//裁剪x丟棄混入的非數字字符
	x = trim(x)

	//根據sign和x 轉換成int
	return convert(sign, x)
}

func getSign(s string) (int, string) {
	sign := 1
	switch s[0] {
	case '-':
		s = s[1:]
		sign = -1.0
	case '+':
		s = s[1:]
	default:
	}
	return sign, s
}

func trim(s string) string {
	for i := range s {
		if s[i] < '0' || '9' < s[i] {
			return s[:i]
		}
	}
	return s
}

func convert(sign int, s string) int {
	base := 1 * sign
	res := 0
	yes := false

	for i := len(s) - 1; i >= 0; i-- {
		if len(s)>48{
			if base>0{
				return math.MaxInt32
			}
		}
		//為了防止特別長的數字 甚至超過float64的範圍 所以每一步都要檢查是否溢出
		res, yes = isOverflow(res + (int(s[i])-48)*base)
		if yes {
			return res
		}
		base *= 10
	}
	return res
}

func isOverflow(i int) (int, bool) {
	switch {
	case i > math.MaxInt32:
		return math.MaxInt32, true
	case i < math.MinInt32:
		return math.MinInt32, true
	default:
		return i, false
	}
}
