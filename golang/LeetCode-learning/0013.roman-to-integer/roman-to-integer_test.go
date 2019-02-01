package problem0013

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参數
// one 代表第一個参數
type para struct {
	one string
}

// ans 是答案
// one 代表第一個答案
type ans struct {
	one int
}

func Test_Problem0013(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{"XXXIX"},
			ans{39},
		},
		question{
			para{"MDCCCLXXXVIII"},
			ans{1888},
		},
		question{
			para{"MMI"},
			ans{2001},
		},
		question{
			para{"MCMLXXVI"},
			ans{1976},
		},
		question{
			para{"MMMCMXCIX"},
			ans{3999},
		},
	}
	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, romanToInt(p.one), "輸入:%v", p)
	}
}
