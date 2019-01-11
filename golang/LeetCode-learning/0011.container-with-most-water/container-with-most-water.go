package problem0011

func maxArea(height []int) int {
	//從兩端開始尋找 至少保證了寬度是最大值
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)

		area := h * (j - i)
		if max < area {
			max = area
		}

		//朝著area具有變大的可能性方向變化
		if a < b {
			i++
		} else {
			j--
		}
	}

	return max
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
