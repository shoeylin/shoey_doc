package problem0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two []int
}

type ans struct {
	one float64
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
				one: []int{1, 3},
				two: []int{2},
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: []int{1, 3},
				two: []int{2, 4},
			},
			a: ans{
				one: 2.5,
			},
		},
		question{
			p: para{
				one: []int{1},
				two: []int{1},
			},
			a: ans{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findMedianSortedArrays(p.one, p.two), "輸入:%v", p)
	}
	ast.Panics(func() { findMedianSortedArrays([]int{}, []int{}) }, "對空切片求中位數 卻沒有panic")
}
