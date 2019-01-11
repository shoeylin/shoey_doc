package problem0010

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "aa",
				two: "a",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "aa",
				two: ".*",
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isMatch(p.one, p.two), "輸入:%v", p)
	}
}
