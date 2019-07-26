package concat

import (
	"bufio"
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
	length := 15000
	size := 15000
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
	str := ConcatV3(sequences, min)
	fmt.Println(*str)

	//for i := 0; i < len(sequences); i++ {
	//	if strings.Index(*str, sequences[i]) == -1 {
	//		t.Error("不包含", sequences[i])
	//	}
	//}
}

func TestTTTT(t *testing.T) {
	buf := make([]byte, 10000000)
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(rand.Intn(128))
	}
	str := string(buf)
	s := str
	if s == str {
		print(true)
	}
	sp := &str
	if sp == &str {

	}
}

func TestReadMap(t *testing.T) {
	readFile := func(fileName string) map[string][]string {
		file, _ := os.Open(fileName)
		reader := bufio.NewReader(file)
		line, e := reader.ReadString('\n')
		line = line[:len(line)-1]
		prefixMap := make(map[string][]string)
		for e == nil {
			colon := strings.Index(line, ":")
			prefix := line[:colon]
			split := strings.Split(line[colon+1:], " ")
			prefixMap[prefix] = split
			line, e = reader.ReadString('\n')
			line = line[:len(line)-1]
		}
		return prefixMap
	}

	map1 := readFile("context3-9.txt")
	map2 := readFile("context3-3230.txt")
	nullCount := 0
	diffCount := 0
	totalCount := 0
	for prefix, seq1 := range map1 {
		totalCount++
		seq2, ok := map2[prefix]
		if !ok {
			fmt.Printf("%s not found\n", prefix)
			nullCount++
			continue
		}
		if len(seq2) != len(seq1) {
			diffCount++
		} else {
			for i := 0; i < len(seq2); i++ {
				found := false
				for j := 0; j < len(seq1); j++ {
					if seq1[j] == seq2[i] {
						found = true
						break
					}
				}
				if !found {
					diffCount++
					break
				}
			}
		}
	}
	fmt.Printf("total: %d null: %d diff: %d\n", totalCount, nullCount, diffCount)

}
