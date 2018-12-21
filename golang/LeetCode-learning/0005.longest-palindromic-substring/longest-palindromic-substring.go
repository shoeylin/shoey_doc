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
			//因為
		}

	}
	return ""
}
