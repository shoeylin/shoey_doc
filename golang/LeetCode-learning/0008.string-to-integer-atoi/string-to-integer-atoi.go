package problem0008

import (
	"strconv"
	"strings"
)

func myAtoi(str string) int {

	res, _ := strconv.Atoi(strings.TrimSpace(str))

	return res
}
