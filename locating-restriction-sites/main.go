package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Print("Filename: ")

	var buffer bytes.Buffer

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	buffer.WriteString(scanner.Text())

	filename := buffer.String()

	byteSequence, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("Error openning file %s \n", err.Error())
		os.Exit(1)
	}
	fullSequence := string(byteSequence)
	inputArray := strings.Split(fullSequence, "\n")[1:]

	var buffer2 bytes.Buffer
	for k := 0; k < len(inputArray); k++ {
		buffer2.WriteString(inputArray[k])
	}
	sequence := buffer2.String()

	//fmt.Printf("using seq %s", sequence)

	seqArray := strings.Split(sequence, "")
	//fmt.Printf("Checking sequence %s\n", seqArray)
	for i := 0; i < len(seqArray); i++ {
		subSeq := seqArray[i:]
		start := i + 1
		max := 12
		if len(subSeq) < 12 {
			max = len(subSeq)
		}
		//fmt.Printf("max is %d\n", max)
		for j := 4; j <= max; j += 2 {
			thisSeq := subSeq[:j]
			//fmt.Printf("checking %s\n", thisSeq)
			seqLength := len(thisSeq)
			if isReversable(thisSeq) {
				fmt.Printf("%d %d\n", start, seqLength)
			}

		}
	}
}

func isReversable(array []string) bool {
	length := len(array)
	for i := 0; i < length/2; i++ {
		a := array[i]
		b := array[length-i-1]
		//fmt.Printf("checking %s and %\n", a, b)
		if a == "A" && b != "T" {
			return false
		}
		if a == "T" && b != "A" {
			return false
		}
		if a == "C" && b != "G" {
			return false
		}
		if a == "G" && b != "C" {
			return false
		}

	}
	return true
}
