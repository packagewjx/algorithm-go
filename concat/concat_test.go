package concat

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	a := "wujun"
	b := "wujunxian"
	fmt.Println(compare(&a, &b))

	a = "www"
	b = "www"
	fmt.Println(compare(&a, &b))
}

func contains(str string, sequences []string) bool {
	for i := 0; i < len(sequences); i++ {
		if !strings.Contains(str, sequences[i]) {
			return false
		}
	}
	return true
}

func TestSmallData(t *testing.T) {
	str := []string{"WEFED", "DNORE", "EFEDNORE", "NIROEPJ", "RETPRNI", "PJGFVM", "GFVMWRD", "WRDNWEN", "WRDNWE",
		"ENERDKN", "CVPENLWXR", "ERDKNNXW", "NXWIURE", "EMBCREM", "IUREMBCRE", "BJLKCVP", "PENL", "XRSIW", "ENLWX",
		"RETPRNIROE", "EPJGFVMW", "WRD", "RDKNN", "UREMBC", "BJLKCV", "PENLWX", "ENLWXR", "REMBJLK", "IUREMB", "REMBJLKC"}
	result := ConcatV3(str, 3)
	print(*result, "\n")
	print(contains(*result, str), "\n")
}

func TestBigData(t *testing.T) {
	length := 20000
	size := 20000
	min := 80
	max := 150

	// 总字符串
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = byte(rand.Intn(26) + 'A')
	}
	longString := string(buf)
	file, _ := os.Create("testData.txt")
	sequences := make([]string, 0, size)
	for i := 0; i < size; i++ {
		len := rand.Intn(max-min) + min
		begin := rand.Intn(length - len)
		seq := longString[begin : begin+len]
		sequences = append(sequences, seq)
		file.WriteString(seq + "\n")
	}
	file.Close()

	file, _ = os.Create("answer.txt")
	file.WriteString(longString)
	file.Close()

	print("start\n")
	str := ConcatV4(sequences)
	contains(*str, sequences)
	println(*str)
}
