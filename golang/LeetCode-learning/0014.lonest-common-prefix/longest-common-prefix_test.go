package problem0014 // Easy
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是參數
// one 代表第一個參數
type para struct {
	one []string
}

// ans 是答案
// one 代表第一個答案
type ans struct {
	one string
}

func Test_Problem0014(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{
				[]string{"abcd", "abcde", "ab"},
			},
			ans{"ab"},
		},
		question{
			para{
				[]string{},
			},
			ans{""},
		},
		question{
			para{
				[]string{"abcd", "abcde", "ab"},
			},
			ans{"ab"},
		},
		//如需多個測試 可以複製上方元素
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, longestCommonPrefix(p.one), "輸入:%v", p)
	}
}
