package problem0006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two int
}

type ans struct {
	one string
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
				one: "PAYPALISHIRING",
				two: 3,
			},
			a: ans{
				one: "PAHNAPLSIIGYIR",
			},
		},
		question{
			p: para{
				one: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
				two: 4,
			},
			a: ans{
				one: "AGMSYBFHLNRTXZCEIKOQUWDJPV",
			},
		},
		question{
			p: para{
				one: "A",
				two: 3,
			},
			a: ans{
				one: "Aa",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, convert(p.one, p.two), "輸入:%v", p)
	}
}
