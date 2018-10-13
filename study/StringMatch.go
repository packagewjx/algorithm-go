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
	//首先将字符串转过来，然后使用后缀搜索除去后缀之后的字符串即可，若搜索到，则放入goodSuffix(i)
	// 若没有，则使用goodSuffix(i - 1)作为goodSuffix(i)
	reversePattern := reverseString(pattern)
	for i := 1; i < len(pattern); i++ {
		// 借用库函数来寻找第一个好后缀。index是找到的这个好后缀在整条模式字符串中，从后面开始数过来的下标值
		if goodIndex := strings.Index(reversePattern[i:], reversePattern[:i]); goodIndex == -1 {
			// 若没有找到，需要寻找短一点的好后缀的，在最前面是否出现
			// 因为如果好后缀中的后面一部分在最前面出现了，就可能会有好后缀的另外一部分出现在整条模式字符串的前面
			// 比如在好后缀长度i=3时，ABCBAB，好后缀BAB，但是其中AB在最前面出现了，那么AB前面的B也可能继续出现
			// 因此这个好后缀移动是4。

			// 首先我可以排除掉i=i-1时候的长度，因为i-1的好后缀，没有让i时候的好后缀继续在该位置匹配
			// 比如ABCBAB这里，好后缀为B时，结果为2，但是好后缀为AB时，由于移动2到达的B前面是C，因此不匹配AB，那么肯定得
			// 在前缀ABCB中看看是否有好后缀的后缀部分了，又因为运行到这里的时候，我已经知道我这条好后缀没有在除去好后缀
			// 的字符串找到，所以，可以去掉i-1时候的找到的好后缀后面的所有字符串，即，CB的B和后面的字符串。

			// 看看上一个好后缀的位置，是否在最前面，如果是的话，说明这条i的好后缀没有匹配，但是有小一点的好后缀，在最前面
			// 有匹配，我可以使用。具体是判断i-1与这个最前面的匹配好后缀的前缀的长度哪个大，如果i-1大于等于，则说明这个
			// 前缀的确是短与i-1的，i-1的结果我可以直接拿来用。如果不是，则说明这个前缀其实不是前缀，是在中间的，
			// 但是就是因为在中间，i-1的匹配，i的时候
			// 不匹配，说明这个子串，多了好后缀一个字符之后，就不匹配了，比如BAOBAB，i为1的时候，找到下标为3的B，但是i为
			// 2的时候，前面没有AB的匹配项，找i-1的时候，那个B的前面是O，我就不应该直接取i-1的值了作为i时候的值。此时应该
			// 继续寻找，在i-1这个匹配的位置之前，寻找i-1那条好后缀是否有另外一个匹配，比如例子中最前面的B，就有匹配了，
			// 所以此时的i的结果是5。
			// 但是，这个中间，指的是除去i时候的后缀的前面的字符串的中间，没有包含一种情况，这个i-1时候的字符串，
			// 其移动的位置，恰好到i时候多加的那个字符的位置，比如BABABA，i为2是移动2，i为3是，也应该移动2，因为刚好匹配
			// 到了倒数第三个A，这个A包含在i的后缀中，并且我们没有拿去比较，并且其前面还是刚好就是ABA。

			if len(reversePattern)-table.goodSuffix[i-1] > i-1 {
				// 这个i-1时候的匹配串在中间，且i时候没有匹配串，
				// 如果刚好i-1移动的位置，在这个好后缀里面
				// 此时我们需要比较，新加入的字符是不是匹配的
				if pattern[len(pattern)-(i-1)-1] == pattern[len(pattern)-table.goodSuffix[i-1]-(i-1)-1] {
					// 如果匹配，则我们可以用上一个的值
					table.goodSuffix[i] = table.goodSuffix[i-1]
					continue
				}

				// 如果不匹配，我们只能寄希望字符串的长度为i-1的前缀，
				// 等于i-1时候的字符串，如果不是，则彻底没有匹配，接下来都需要整条所有字符串
				if pattern[:i-1] == pattern[len(pattern)-(i-1):] {
					table.goodSuffix[i] = len(reversePattern) - (i - 1)
				} else {
					table.goodSuffix[i] = len(reversePattern)
				}
			} else {
				table.goodSuffix[i] = table.goodSuffix[i-1]
			}

			// 这里只会进来一次，找不到一次，后面更不会找到，用i-1的值填入即可
			// 这个包含了i-1时在最前面有匹配的情况，此时也是用i-1的值填充接下来的所有值
			for i += 1; i < len(pattern); i++ {
				table.goodSuffix[i] = table.goodSuffix[i-1]
			}
			break

		} else {
			table.goodSuffix[i] = goodIndex + i
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
