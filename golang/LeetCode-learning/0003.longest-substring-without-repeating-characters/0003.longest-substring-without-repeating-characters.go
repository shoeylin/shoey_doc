package problem0003

func lengthOfLongestSubstring(s string) int {
	// location[s[i]] == j 表示:
	// s中第i個字符串 上次出現在s的位置 所以 在s[j+1:i]中沒有s[i]
	// location[s[i]] == -1 表示: s[i] 在s中第一次出現
	location := [256]int{} // 只有256長是因為 假定輸入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1 // 先設置所有的字符都沒有見過
	}
	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		//說明s[i]已經在s[left:i+1]中重複了
		//並且s[i]上次出現的位置在location[s[i]]
		if location[s[i]] >= left {
			left = location[s[i]] + 1 //在s[left:i+1]中去除s[i]字符及其之前的部分
		} else if i+1-left > maxLen {
			//fmt.Pringln(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}
	return maxLen
}
