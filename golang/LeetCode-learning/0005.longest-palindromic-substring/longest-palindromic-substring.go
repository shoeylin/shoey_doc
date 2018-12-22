package problem0005

func longestPalindrome(s string) string {
	if len(s) < 2 { //肯定是回文 直接返回
		return s
	}
	// 最長回文的首字符索引 和最長回文的長度
	begin, maxLen := 0, 1

	// 在for 循環中
	// b 代表回文的**首**字符索引號
	// e 代表回文的**尾**字符索引號
	// i 代表回文`正中間段`首字符的索引號
	// 在每一次for循環中
	// 先從i開始 利用`正中間段`所有字符相同的特性 讓b,e分別指向`正中間段`的首尾
	// 再從`正中間段`向兩邊擴張 讓b e分別指向此正中間段為中心的最長回文的首尾
	for i := 0; i < len(s); { //以s[i]為正中間段首字符開始尋找最長回文
		if len(s)-i <= maxLen/2 {
			// 因為i是回文`正中間段`首字符的索引號
			// 假設此時能找到最長回文的長度為l, 則 l <= (len(s)-i)*2 -1
			// 如果 len(s)-i <= maxLen/2 成立
			// 則 l <= maxLen - 1
			// 則 l < maxLen
			// 所以 無須再找下去
			break
		}

		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
			//循環結果後 s[b:e+1]是一串相同的字符串
		}

		//下一個回文的`正中間段`的首字符只會是s[e+1]
		//為下次循環做準備
		i = e + 1

		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
			//循環結束後 s[b:e+1]是這次能找到的最長回文
		}
		newLen := e + 1 - b
		//創建新紀錄的話 就更新紀錄
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}

	}

	return s[begin : begin+maxLen]
}
