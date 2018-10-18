package study

import (
	. "github.com/packagewjx/algorithm-go/datastructure"
	"math"
	"sort"
)

// 存储点的最小堆
type pointArray struct {
	points   []*Point
	compareX bool
}

func (m *pointArray) Len() int {
	return len(m.points)
}

func (m *pointArray) Less(i, j int) bool {
	if m.compareX {
		return m.points[i].X < m.points[j].X
	} else {
		return m.points[i].Y < m.points[j].Y
	}
}

func (m *pointArray) Swap(i, j int) {
	temp := m.points[i]
	m.points[i] = m.points[j]
	m.points[j] = temp
}

func EfficientClosestPair(P []*Point) float64 {
	if len(P) <= 3 {
		return bruteForceClosestPair(P)
	}

	Q := make([]*Point, len(P), len(P))
	copy(Q, P)

	p := &pointArray{P, true}
	sort.Sort(p)

	q := &pointArray{Q, false}
	sort.Sort(q)

	return efficientClosestPair(p.points, q.points)
}

// P：以x升序排列的点集
// Q：以y升序排列的点集
// 返回最近的距离
func efficientClosestPair(P, Q []*Point) float64 {
	if len(P) <= 3 {
		return bruteForceClosestPair(P)
	}

	leftP, leftQ, rightP, rightQ := splitPointsWithNonDecreaseX(P)
	leftD := efficientClosestPair(leftP, leftQ)
	rightD := efficientClosestPair(rightP, rightQ)
	minD := leftD
	if rightD < minD {
		minD = rightD
	}

	// 取len/2-1的上界，那个点的x
	middleX := P[len(P)>>1+len(P)&1-1].X

	pointsInD := xInDPoints(Q, middleX, minD)
	minDSq := minD * minD

	for i := 0; i < len(pointsInD); i++ {
		for j := i + 1; j < len(pointsInD); j++ {
			p1 := pointsInD[i]
			p2 := pointsInD[j]
			if d := euclideanDistanceNoSqrt(p1, p2); d < minDSq {
				minDSq = d
			}
		}
	}

	return math.Sqrt(minDSq)
}

// 蛮力计算最近点对
func bruteForceClosestPair(points []*Point) float64 {
	closest := math.MaxFloat64
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if d := euclideanDistanceNoSqrt(points[i], points[j]); closest > d {
				closest = d
			}
		}
	}
	return math.Sqrt(closest)
}

// 算出两点距离，但是不算平方根，以节省成本
func euclideanDistanceNoSqrt(p1, p2 *Point) float64 {
	return (p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)
}

func splitPointsWithNonDecreaseX(points []*Point) (leftP, leftQ, rightP, rightQ []*Point) {
	p := &pointArray{points: points, compareX: true}
	sort.Sort(p)

	lQ := &pointArray{points: make([]*Point, 0, len(p.points)/2+1), compareX: false}
	rQ := &pointArray{points: make([]*Point, 0, len(p.points)/2+1), compareX: false}

	for i := 0; i < len(p.points)/2; i++ {
		lQ.points = append(lQ.points, p.points[i])
	}

	for i := len(p.points) / 2; i < len(p.points); i++ {
		rQ.points = append(rQ.points, p.points[i])
	}

	sort.Sort(lQ)
	sort.Sort(rQ)

	return p.points[:len(p.points)/2], lQ.points, p.points[len(p.points)/2:], rQ.points
}

// Q：按Y升序排列的点集
func xInDPoints(Q []*Point, middleX float64, xDistance float64) []*Point {
	result := make([]*Point, 0, len(Q)>>2)

	for i := 0; i < len(Q); i++ {
		d := Q[i].X - middleX
		if d < 0 {
			d = -d
		}

		if d < xDistance {
			result = append(result, Q[i])
		}
	}

	return result
}
