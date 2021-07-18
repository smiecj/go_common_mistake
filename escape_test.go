package go_common_mistake

import (
	"fmt"
	"testing"
)

func sum() int {
	nums := make([]int, 100)
	for i := 1; i <= len(nums); i++ {
		nums[i-1] = i
	}

	s := 0
	for i := range nums {
		s += i
	}
	return s
}

//func TestNotEscape(t *testing.T) {
//	s := sum()
//	log.Println(s)
//}

type Point struct {
	X int
	Y int
}

func center(p *Point) {
	p.X = p.X / 2
	p.Y = p.Y / 2
}

//go:noinline
func makePoint() *Point {
	p := new(Point)
	p.X = 10
	p.Y = 6
	return p
}

func TestNotEscape(t *testing.T) {
	//p := new(Point)
	p := makePoint()
	//p.X = 10
	//p.Y = 6
	center(p)
	fmt.Println(p.X, p.Y)
}
