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
		if p[j]=='*'
	}
}
