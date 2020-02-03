package problem0015

import "sort"

func threeSum(nums []int) [][]int {
	// 排序後 可以按規律查找
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		if nums[i] > 0 {
			break
		}
		// 避免添加重複的結果
		// i > 0 是為了防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				// 較小的 l 需要變大
				l++
			case s > 0:
				// 較大的 r 需要變小
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				//為避免重複添加, l 和 r 都需要移動不同的元素上
				l, r = next(nums, l, r)
			}
		}

	}

	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}
