package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
)

func generate(originalLength, sequencesLength, minSequenceLength, maxSequenceLength int) {
	buf := make([]byte, originalLength)
	for i := 0; i < originalLength; i++ {
		buf[i] = byte(rand.Intn(26) + 'A')
	}
	longString := string(buf)

	file, e := os.Create("originalSequence.txt")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
	file.WriteString(longString)
	file.Close()

	file, e = os.Create("sequences.txt")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
	for i := 0; i < sequencesLength; i++ {
		len := rand.Intn(maxSequenceLength-minSequenceLength+1) + minSequenceLength
		begin := rand.Intn(originalLength - len + 1)
		seq := longString[begin : begin+len]
		file.WriteString(seq + "\n")
	}
	file.Close()
}

func main() {
	fmt.Printf("%d\n", runtime.NumCPU())

	argv := os.Args
	if len(argv) != 5 {
		fmt.Printf("usage: %s originalLength sequencesLength minSequenceLength maxSequenceLength\n", argv[0])
		os.Exit(1)
	}
	originalLength, err := strconv.Atoi(argv[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	sequencesLength, err := strconv.Atoi(argv[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	minSequenceLength, err := strconv.Atoi(argv[3])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	maxSequenceLength, err := strconv.Atoi(argv[4])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	generate(originalLength, sequencesLength, minSequenceLength, maxSequenceLength)
}
