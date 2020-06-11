package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

var cp ColoredPoint

func main() {
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// method expression
func init() {
	p := Point{1, 2}
	q := Point{4, 6}
	distance := Point.Distance  // method expression
	fmt.Println(distance(p, q)) // 5
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
	fmt.Println("init1 ----------------------------")
}

// method value
func init() {
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance   // method value
	fmt.Println(distanceFromP(q)) // 5
	var origin Point
	fmt.Println(distanceFromP(origin)) // 2.23...

	scaleP := p.ScaleBy // method value
	scaleP(2)
	scaleP(3)
	scaleP(10)
	fmt.Println(p)
	fmt.Println("init 2----------------------------")
}

// ColoredPoint methods are promoted indirectly from *Point
func init() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	type ColoredPoint struct {
		*Point
		Color color.RGBA
	}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point)) // "5"
	q.Point = p.Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // {2 2} {2 2}
	fmt.Println("init3 ----------------------------")
}
