package problem0001

func twoSum(nums []int, target int) []int {
	//index 負責保存map[整數]整數的序列號
	index := make(map[int]int, len(nums))

	//通過 for 循環 獲取b的序列號
	for i, b := range nums {
		//通過查詢map 獲取a = target - b 的序列號
		if j, ok := index[target-b]; ok {
			//ok 為true
			//說明在i之前 存在 nums[j] == a
			return []int{j, i}
			// 注意 順序是 j i
		}

		//把 b 和 i 的值 存入map
		index[b] = i
	}

	return nil
}

func test(sum []int, target int) []int{

	index := make(map[int]int, len(sum))

	for i, b := range sum {
		if j, ok := index[target-b]; ok {
			return []int{j, i}
		}
		index[b] = i
	}
}
