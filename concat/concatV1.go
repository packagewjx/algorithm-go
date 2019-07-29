package concat

import (
	"fmt"
	"os"
)

type context struct {
	indices       [26]index
	minimumLength int
}

type index struct {
	sequences []*sequence
}

type sequence struct {
	sequence *string
	seqType  int
}

const CMP_EQ = 1
const CMP_NEQ = 2
const CMP_A_PREFIX_B = 3
const CMP_B_PREFIX_A = 4
const TYPE_STRING = 1
const TYPE_CONCAT = 2
const TYPE_SUFFIX = 3

// 比较a和b的关系，并得出是否相等和前后缀关系
func compare(a *string, b *string) int {
	if len(*a) == len(*b) {
		if *a == *b {
			return CMP_EQ
		} else {
			return CMP_NEQ
		}
	} else if len(*a) > len(*b) {
		c := (*a)[0:len(*b)]
		if *b == c {
			return CMP_B_PREFIX_A
		} else {
			return CMP_NEQ
		}
	} else {
		c := (*b)[0:len(*a)]
		if *a == c {
			return CMP_A_PREFIX_B
		} else {
			return CMP_NEQ
		}
	}
}

func compareSequence(seq1 *sequence, seq2 *sequence) int {
	return compare(seq1.sequence, seq2.sequence)
}

func getIndex(s *string) uint8 {
	return (*s)[0] - 'A'
}

func insertSequenceToIndex(seq *sequence, idx *index) *sequence {
	for i := 0; i < len(idx.sequences); i++ {
		seq2 := idx.sequences[i]
		cmp := compareSequence(seq, seq2)
		if cmp == CMP_NEQ {
			continue
		} else if cmp == CMP_EQ {
			if seq.seqType == TYPE_SUFFIX && seq2.seqType == TYPE_STRING {
				// 如果本来保存了一条字符串，是seq的后缀，就需要设置为后缀，方便后面去除。否则会有重复
				seq2.seqType = TYPE_SUFFIX
			}
			return seq2
		} else if cmp == CMP_A_PREFIX_B {
			return seq2
		} else if cmp == CMP_B_PREFIX_A {
			idx.sequences[i] = seq
			return seq
		} else {
			print("error")
			os.Exit(1)
		}
	}

	// 这里是找不到
	idx.sequences = append(idx.sequences, seq)
	return seq
}

func insertString(str *string, ctx *context) {
	seq := new(sequence)
	seq.sequence = str
	seq.seqType = TYPE_STRING
	if seq != insertSequenceToIndex(seq, &ctx.indices[getIndex(str)]) {
		seq = nil
	}

	for i := 1; i <= len(*str)-ctx.minimumLength; i++ {
		seq = new(sequence)
		subStr := (*str)[i:]
		seq.sequence = &subStr
		seq.seqType = TYPE_SUFFIX
		idxNum := getIndex(&subStr)
		if seq != insertSequenceToIndex(seq, &ctx.indices[idxNum]) {
			seq = nil
			break
		}
	}
}

func removeSuffix(idx *index) {
	s := make([]*sequence, 0, len(idx.sequences)>>5)
	for i := 0; i < len(idx.sequences); i++ {
		if idx.sequences[i].seqType == TYPE_SUFFIX {
			idx.sequences[i] = nil
		} else {
			s = append(s, idx.sequences[i])
		}
	}
	// 移动位置
	put := 0
	for i := 0; i < len(idx.sequences); i++ {
		if idx.sequences[i] == nil {
			continue
		}
		idx.sequences[put] = idx.sequences[i]
		put++
	}
	idx.sequences = idx.sequences[0:put]
}

func combine(idx *index, ctx *context) {
	for i := 0; i < len(idx.sequences); i++ {
		seq := idx.sequences[i]
		combined := false
		for j := 1; !combined && j < len(*seq.sequence); j++ {
			suffix := (*seq.sequence)[j:]
			suffixLen := len(*seq.sequence) - j
			targetIndex := &ctx.indices[getIndex(&suffix)]

			for k, seq2 := range targetIndex.sequences {
				if seq2 == seq {
					// 不能跟自己比较
					continue
				}
				cmp := compare(&suffix, seq2.sequence)
				if cmp == CMP_NEQ {
					continue
				} else if cmp == CMP_A_PREFIX_B {
					combineString := *seq.sequence + (*seq2.sequence)[suffixLen:]
					if seq2.seqType == TYPE_CONCAT {
						seq2.sequence = nil
					}
					seq2 = nil
					// 删除数组中的seq2
					targetIndex.sequences = append(targetIndex.sequences[:k], targetIndex.sequences[k+1:]...)
					if targetIndex == idx && k >= i {
						i--
					}

					seq.sequence = &combineString
					seq.seqType = TYPE_CONCAT

					// 结束循环
					combined = true
					break
				} else {
					print("contain sub sequence")
					os.Exit(1)
				}
			}
		}
	}
}

func Concat(sequences []string, minimumLength int) *string {
	var ctx context
	ctx.minimumLength = minimumLength
	for i := 0; i < 26; i++ {
		ctx.indices[i].sequences = make([]*sequence, 0, 100)
	}

	for i := 0; i < len(sequences); i++ {
		insertString(&sequences[i], &ctx)
	}

	for i := 0; i < len(ctx.indices); i++ {
		removeSuffix(&ctx.indices[i])
	}

	var ret int
	lastSum := -1
	for true {
		for i := 0; i < len(ctx.indices); i++ {
			combine(&ctx.indices[i], &ctx)
		}

		sum := 0
		for idxNum, idx := range ctx.indices {
			sum += len(idx.sequences)
			if len(idx.sequences) == 1 {
				ret = idxNum
			}
		}
		if sum == 1 {
			break
		}
		print(sum, "\n")
		if sum == lastSum {
			print("剩余", sum, "无法拼起来\n")
			printContext(&ctx)
			os.Exit(1)
		}
		lastSum = sum
	}

	return ctx.indices[ret].sequences[0].sequence
}

func printContext(ctx *context) {
	for i := 0; i < len(ctx.indices); i++ {
		fmt.Printf("=========%c=========\n", i+'A')
		for j := 0; j < len(ctx.indices[i].sequences); j++ {
			fmt.Printf("%s %d\n", *(ctx.indices[i].sequences[j].sequence), ctx.indices[i].sequences[j].seqType)
		}
	}
}
