package leetcode

import "testing"

func Test609(t *testing.T) {
	findDuplicate([]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"})
}
