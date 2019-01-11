package problem0010

func isMatch(s, p string) bool {
	sSize := len(s)
	pSize := len(p)

	dp := make([][]bool, sSize+1)
	for i := range dp {
		dp[i] = make([]bool, pSize+1)
	}

	//dp[i][j] 代表了 s[:i]能否與p[:j]匹配

	dp[0][0] = true
	/*
		根據題目的設定 "" 可以與 "a*b*c*" 相匹配
		所以 需要把相應的dp設置成 true
	*/
	for j := 1; j < pSize && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < sSize; i++ {
		for j := 0; j < pSize; j++ {
			if p[j] == '.' || p[j] == s[i] {
				//p[j]與 s[i]可以匹配上 所以 只要前面匹配 這裡就能匹配上
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				//此時 p[j]的匹配情況與 p[j-1]的內容相關
				if p[j-1] != s[i] && p[j-1] != '.' {
					//p[j]無法與 s[i]匹配上
					//p[j-1:j+1] 只能當作""
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					//p[j]與s[i]匹配上
					//p[j-1;j+1]作為"x*", 可以有三種解釋
					dp[i+1][j+1] = dp[i+1][j-1] || // "x*"解釋為 ""
						dp[i+1][j] || // "x*"解釋為 "x"
						dp[i][j+1] // "x*"解釋為 "xxx..."
				}
			}
		}
	}

	return dp[sSize][pSize]
}
