package problem0007

import (
	"math"
)

func reverse(x int) int {
	sign := 1

	//處理負數
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 {
		//取出 x 的末尾
		temp := x % 10
		//放入 res 的開頭
		res = res*10 + temp
		// x 去除末尾
		x = x / 10
	}

	//還原 x 的符號到 res
	res = sign * res

	// 處理 res 的溢出問題
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}

	return res
}
