package study

import (
	"fmt"
	"math/rand"
	"testing"
)

import (
	. "github.com/packagewjx/algorithm-go/datastructure"
)

func TestEfficientClosestPair(t *testing.T) {
	len := 100
	p := make([]*Point, 0, len)

	for i := 0; i < len; i++ {
		p = append(p, &Point{X: rand.Float64(), Y: rand.Float64()})
		fmt.Printf("X: %.10f Y:%.10f\n", p[i].X, p[i].Y)
	}

	bp := bruteForceClosestPair(p)

	fmt.Println(bp)
	fmt.Println(EfficientClosestPair(p))

}
