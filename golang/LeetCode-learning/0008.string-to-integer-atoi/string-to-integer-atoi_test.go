package problem0008

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one int
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
				one: "123",
			},
			a: ans{
				one: 123,
			},
		},
		question{
			p: para{
				one: "123 show",
			},
			a: ans{
				one: 123,
			},
		},
		question{
			p: para{
				one: "-123",
			},
			a: ans{
				one: -123,
			},
		},
		question{
			p: para{
				one: "2147483648",
			},
			a: ans{
				one: math.MaxInt32,
			},
		},
		question{
			p: para{
				one: "-2147483649",
			},
			a: ans{
				one: math.MinInt32,
			},
		},
		question{
			p: para{
				one: " 1234a6789",
			},
			a: ans{
				one: 1234,
			},
		},
		question{
			p: para{
				one: "  -0012a42 ",
			},
			a: ans{
				one: -12,
			},
		},
		question{
			p: para{
				one: " asdfdfs ",
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: "",
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: " +1   ",
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: "-",
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459",
			},
			a: ans{
				one: math.MaxInt32,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, myAtoi(p.one), "輸入:%v", p)
	}
}
