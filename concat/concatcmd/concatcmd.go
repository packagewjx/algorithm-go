package main

import (
	"bufio"
	"github.com/packagewjx/algorithm-go/concat"
	"math"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		print("usage: ", os.Args[0], " fileName")
		os.Exit(1)
	}

	file, e := os.Open(os.Args[1])
	if e != nil {
		print(e.Error())
		os.Exit(1)
	}

	sequences := make([]string, 0, 100)
	reader := bufio.NewReader(file)
	s, e := reader.ReadString('\n')
	// 去掉\n
	minLen := math.MaxInt64
	for e == nil {
		s = s[:len(s)-1]
		if len(s) < minLen {
			minLen = len(s)
		}
		sequences = append(sequences, s)
		s, e = reader.ReadString('\n')
	}
	file.Close()

	result := concat.ConcatV4(sequences)

	file, _ = os.Create("output.txt")
	file.WriteString(*result)
	file.Close()
}
