package study

import "strings"

type StringMatcher interface {
	// 在text中寻找matchString首次出现的位置
	// 规定若传入空字符串或者模式长度比text要长，则返回-1
	// 找不到返回-1
	Match(matchString, text string) int
}

// ========================================Horspool匹配算法=============================================================

type shiftTable struct {
	table map[byte]int
	// 匹配串的长度
	length int
}

// 创建移动表。
// 若matchString长度为0，返回没有内容的表，所有操作都返回0
func newHorspoolTable(matchString string) *shiftTable {
	if len(matchString) == 0 {
		return &shiftTable{map[byte]int{}, 0}
	}
	t := &shiftTable{make(map[byte]int), len(matchString)}
	for i := len(matchString) - 2; i >= 0; i-- {
		char := matchString[i]
		if _, ok := t.table[char]; !ok {
			t.table[char] = len(matchString) - 1 - i
		}
	}
	return t
}

// 查找某个字符的移动.若结果为-1，代表没有char对应的移动值，此时应该移动整个匹配串
func (t *shiftTable) get(char byte) int {
	shift, ok := t.table[char]
	if ok {
		return shift
	} else {
		return t.length
	}
}

type HorspoolMatcher struct {
}

func (HorspoolMatcher) Match(matchString, text string) int {
	if matchString == "" || len(matchString) > len(text) {
		return -1
	}

	table := newHorspoolTable(matchString)

	for rightMost := len(matchString) - 1; rightMost < len(text); {
		for currentInMatchString := len(matchString) - 1; currentInMatchString >= 0; currentInMatchString-- {
			currentCharInText := text[rightMost-len(matchString)+currentInMatchString+1]
			if matchString[currentInMatchString] == currentCharInText {
				// 从后面开始一个一个匹配，若匹配成功，则继续匹配

				if currentInMatchString == 0 {
					// 若已经是第一个字符了，说明这个模式匹配成功，返回值
					return rightMost - len(matchString) + 1
				}
			} else {
				// 否则退出，并移动rightMost

				// 算出当前字符到模式最右端的距离
				charInMatchStringToRightMost := len(matchString) - currentInMatchString - 1
				// 用移动表的值，减去当前字符到模式最右端的距离
				rightMost += table.get(currentCharInText) - charInMatchStringToRightMost
				break
			}
		}
	}

	return -1
}

// ==============================================Boyer-Moore匹配算法====================================================

type boyermooreShiftTable struct {
	badSymbol  map[byte]int
	goodSuffix []int
	length     int
}

func newBoyerMooreTable(pattern string) *boyermooreShiftTable {
	badTable := newHorspoolTable(pattern)
	table := &boyermooreShiftTable{}
	table.length = badTable.length
	table.badSymbol = badTable.table

	// 计算好后缀
	table.goodSuffix = make([]int, len(pattern))
	// 定义好后缀0为整条的长度
	table.goodSuffix[0] = len(pattern)
	//将字符串转过来
	reversePattern := reverseString(pattern)
	for i := 1; i < len(pattern); i++ {
		// 如果刚好i-1移动的位置，在这个好后缀里面
		// 此时我们需要比较，新加入的字符是不是匹配的
		if i > 1 && len(pattern)-table.goodSuffix[i-1]-(i-1)-1 >= 0 && pattern[len(pattern)-(i-1)-1] == pattern[len(pattern)-table.goodSuffix[i-1]-(i-1)-1] {
			// 如果匹配，则我们可以用上一个的值
			table.goodSuffix[i] = table.goodSuffix[i-1]

		} else if goodIndex := strings.Index(reversePattern[i:], reversePattern[:i]); goodIndex != -1 {
			// 借用库函数来寻找第一个好后缀。index是找到的这个好后缀在整条模式字符串中，从后面开始数过来的下标值
			table.goodSuffix[i] = goodIndex + i
		} else {
			// 我们只能寄希望字符串的长度为i-1的前缀，
			// 等于i-1时候的好后缀，如果不是，则彻底没有匹配，接下来都需要整条所有字符串
			if pattern[:i-1] == pattern[len(pattern)-(i-1):] {
				table.goodSuffix[i] = len(reversePattern) - (i - 1)
			} else {
				table.goodSuffix[i] = len(reversePattern)
			}

			// 这里只会进来一次，找不到一次，后面更不会找到，用i-1的值填入即可
			// 这个包含了i-1时在最前面有匹配的情况，此时也是用i-1的值填充接下来的所有值
			for i += 1; i < len(pattern); i++ {
				table.goodSuffix[i] = table.goodSuffix[i-1]
			}
			break
		}
	}

	return table
}

func (t *boyermooreShiftTable) getBadSymbol(char byte) int {
	if val, ok := t.badSymbol[char]; ok {
		return val
	} else {
		return t.length
	}
}

// 反转一条字符串
func reverseString(text string) string {
	builder := strings.Builder{}
	for i := len(text) - 1; i >= 0; i-- {
		builder.WriteByte(text[i])
	}
	return builder.String()
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}

type BoyerMooreMatcher struct {
}

func (BoyerMooreMatcher) Match(matchString, text string) int {
	if matchString == "" || len(matchString) > len(text) {
		return -1
	}

	table := newBoyerMooreTable(matchString)

	for rightMost := len(matchString) - 1; rightMost < len(text); {
		for currentInPattern := len(matchString) - 1; currentInPattern >= 0; currentInPattern-- {
			charCompared := len(matchString) - 1 - currentInPattern
			currentCharInText := text[rightMost-charCompared]
			if matchString[currentInPattern] != currentCharInText {
				if charCompared == 0 {
					// 此时，一个字符都没有匹配，则移动坏后缀表中指定的长度
					rightMost += table.getBadSymbol(currentCharInText)
					break
				} else {
					badSymbol := table.getBadSymbol(currentCharInText) - charCompared
					goodSuffix := table.goodSuffix[charCompared]
					rightMost += max(badSymbol, goodSuffix)
				}
			} else if currentInPattern == 0 {
				return rightMost - len(matchString) + 1
			}
		}
	}
	return -1
}
