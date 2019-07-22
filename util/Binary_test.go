package util

import (
	"github.com/packagewjx/algorithm-go/util/testutils"
	"testing"
)

func TestHighestOne(t *testing.T) {
	testutils.AssertEqual(t, -1, HighestOne(0))
	testutils.AssertEqual(t, 0, HighestOne(1))
	testutils.AssertEqual(t, 1, HighestOne(2))
	testutils.AssertEqual(t, 1, HighestOne(3))
	testutils.AssertEqual(t, 2, HighestOne(4))
	testutils.AssertEqual(t, 12, HighestOne(0x1FFF))
	testutils.AssertEqual(t, 31, HighestOne(0xFFFFFFFF))
	testutils.AssertEqual(t, 28, HighestOne(0x1FFFFFFF))
}
