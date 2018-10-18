package datastructure

import "math"

type Point struct {
	X float64
	Y float64
}

func (p *Point) EuclideanDistance(p2 *Point) float64 {
	return math.Sqrt((p.X-p2.X)*(p.X-p2.X) + (p.Y-p2.Y)*(p.Y-p2.Y))
}

type Line struct {
	Slope    float64
	Constant float64
}

func NewLine(p1, p2 *Point) *Line {
	line := &Line{}
	line.Slope = (p1.Y - p2.Y) / (p1.X - p2.X)
	line.Constant = p2.Y - line.Slope*p2.X
	return line
}

func (l *Line) CalY(x float64) float64 {
	return l.Slope*x + l.Constant
}

func (l *Line) CalX(y float64) float64 {
	return (y - l.Constant) / l.Slope
}

func (l *Line) IsAbove(p1 *Point) (isAbove bool) {
	return l.CalY(p1.X) <= p1.Y
}

type Triangle struct {
	Points [3]*Point
}

func NewTriangle(p1, p2, p3 *Point) *Triangle {
	return &Triangle{[3]*Point{p1, p2, p3}}
}

func (t *Triangle) Area() float64 {
	p1 := t.Points[0]
	p2 := t.Points[1]
	p3 := t.Points[2]
	temp := p1.X*p2.Y + p3.X*p1.Y + p2.X*p3.Y - p3.X*p2.Y - p2.X*p1.Y - p1.X*p3.Y
	if temp > 0 {
		return temp / 2
	} else {
		return -temp / 2
	}
}
