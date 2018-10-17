package study

import (
	. "github.com/packagewjx/algorithm-go/datastructure"
)

var convexHullPoints []*Point

func QuickConvexHull(points []*Point) []*Point {
	convexHullPoints = []*Point{}
	if len(points) == 0 {
		return []*Point{}
	}

	// 找出最左和最右
	leftMost := points[0]
	rightMost := points[0]
	for i := 1; i < len(points); i++ {
		if points[i].X < leftMost.X {
			leftMost = points[i]
		} else if points[i].X > rightMost.X {
			rightMost = points[i]
		}
	}

	convexHullPoints = append(convexHullPoints, leftMost, rightMost)

	upperPoints, lowerPoints := separatePoints(leftMost, rightMost, points)
	qCHRecursive(leftMost, rightMost, upperPoints, true)
	qCHRecursive(leftMost, rightMost, lowerPoints, false)
	return convexHullPoints
}

// isAbove指定是算上方的还是下方的最大点
// Points应该都是在leftMost和rightMost上方或者是下方的，看isAbove的值
func qCHRecursive(leftMost, rightMost *Point, points []*Point, isAbove bool) {
	if len(points) == 0 {
		return
	}

	// 找出最大面积的点
	var areaMost *Point
	largestArea := float64(0)
	for i := 0; i < len(points); i++ {
		area := NewTriangle(leftMost, rightMost, points[i]).Area()
		if area > largestArea {
			areaMost = points[i]
			largestArea = area
		}
	}

	convexHullPoints = append(convexHullPoints, areaMost)

	leftUpperPoints, leftLowerPoints := separatePoints(leftMost, areaMost, points)
	rightUpperPoints, rightLowerPoints := separatePoints(areaMost, rightMost, points)
	if isAbove {
		qCHRecursive(leftMost, areaMost, leftUpperPoints, true)
		qCHRecursive(areaMost, rightMost, rightUpperPoints, true)
	} else {
		qCHRecursive(leftMost, areaMost, leftLowerPoints, false)
		qCHRecursive(areaMost, rightMost, rightLowerPoints, false)
	}
}

func separatePoints(leftMost, rightMost *Point, points []*Point) (upperPoints, lowerPoints []*Point) {
	if len(points) == 0 {
		return []*Point{}, []*Point{}
	}
	line := NewLine(leftMost, rightMost)
	upperPoints = make([]*Point, 0, len(points)/2)
	lowerPoints = make([]*Point, 0, len(points)/2)
	for i := 0; i < len(points); i++ {
		if points[i] == leftMost || points[i] == rightMost {
			continue
		}
		if line.IsAbove(points[i]) {
			upperPoints = append(upperPoints, points[i])
		} else {
			lowerPoints = append(lowerPoints, points[i])
		}
	}
	return
}
