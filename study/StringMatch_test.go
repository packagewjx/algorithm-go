package study

import (
	"fmt"
	"github.com/packagewjx/algorithm-go/util/testutils"
	"testing"
)

func TestShiftTable(t *testing.T) {
	str := "BARBER"
	table := newHorspoolTable(str)
	testutils.AssertTrue(table.get('E') == 1, "", t)
	testutils.AssertTrue(table.get('B') == 2, "", t)
	testutils.AssertTrue(table.get('A') == 4, "", t)
	testutils.AssertTrue(table.get('D') == 6, "", t)
	testutils.AssertTrue(table.get('R') == 3, "", t)

	str = ""
	table = newHorspoolTable(str)
	testutils.AssertTrue(table != nil, "", t)

	str = "A"
	table = newHorspoolTable(str)
	testutils.AssertTrue(table.get('A') == 1, "", t)
	testutils.AssertTrue(table.get('B') == 1, "", t)
}

func TestHorspoolMatcher(t *testing.T) {
	matcher := HorspoolMatcher{}

	testMatcher(matcher, t)
}

func testMatcher(matcher StringMatcher, t *testing.T) {
	fun := func(matchString, text string, expect int) {
		result := matcher.Match(matchString, text)
		testutils.AssertTrue(result == expect,
			fmt.Sprintf("模式串：%s，总文本：%s，结果应该为%d，但是为%d", matchString, text, expect, result), t)
	}

	fun("BARBER", "JIM SAW ME IN A BARBERSHOP", 16)
	fun("A", "JIM SAW ME IN A BARBERSHOP", 5)
	fun("J", "JIM SAW ME IN A BARBERSHOP", 0)
	fun("SHOP", "JIM SAW ME IN A BARBERSHOP", 22)
	fun(" SAW", "JIM SAW ME IN A BARBERSHOP", 3)
	fun("", "JIM SAW ME IN A BARBERSHOP", -1)
	fun("BABABA", "CBCBVVBABACDSFBABABAFKDSFE", 14)
}

func TestBoyerMooreMatcher_Match(t *testing.T) {
	testMatcher(BoyerMooreMatcher{}, t)
}

func TestReverseString(t *testing.T) {
	fun := func(text, expected string) {
		result := reverseString(text)
		testutils.AssertTrue(result == expected, fmt.Sprintf("原串：%s，输出：%s，应该为：%s"), t)
	}

	fun("ABC", "CBA")
	fun("", "")
	fun("A", "A")
}
