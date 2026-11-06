package main

import "fmt"

type Point struct{ X, Y int }

var global *Point

func f1() int         { x := 42; return x }
func f2() *int        { x := 42; return &x }
func f3(p *Point) int { return p.X + p.Y }
func f4()             { p := Point{1, 2}; fmt.Println(p) }
func f5() []int       { s := make([]int, 10); return s }
func f6() int         { s := make([]int, 10); return len(s) }
func f7(n int) int    { s := make([]int, n); return len(s) }
func f8() func() int  { x := 0; return func() int { x++; return x } }
func f9(p *Point)     { global = p }
func f10() byte       { var big [200_000]byte; return big[0] }

func main() {
	p := Point{3, 4}
	f9(&p) 

	_ = f1()
	_ = f2()
	_ = f3(&p)
	f4()
	_ = f5()
	_ = f6()
	_ = f7(10)
	_ = f8()
	_ = f10()
}
