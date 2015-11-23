package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ByOrder []string

var orderedCodes []string

func GetVal(str string) int {
	// get code value
	refCode := strings.Join(orderedCodes, "")
	refCodeSize := len(orderedCodes)

	size := len(str)
	base := len(orderedCodes) + 1
	codes := strings.Split(str, "")
	var total int
	for i := 0; i < size; i++ {
		exp := size - i
		weight := int(math.Pow(float64(base), float64(exp)))
		thisCode := codes[i]

		ind := refCodeSize - strings.Index(refCode, thisCode)
		total += int(weight * ind)
	}
	return (-1 * total)
}
func (s ByOrder) Len() int {

	return len(s)
}
func (s ByOrder) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByOrder) Less(i, j int) bool {

	return GetVal(s[i]) < GetVal(s[j])
}

func main() {
	filename := "rosalind_lexf.txt"
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error is %s", err.Error())
		os.Exit(1)
	}
	input := string(f)
	tokens := strings.Split(input, "\n")
	codeLine := tokens[0]
	numLine := tokens[1]

	iter64, _ := strconv.ParseInt(numLine, 10, 64)
	iter := int(iter64)
	codes := strings.Split(codeLine, " ")

	orderedCodes = codes

	var combos []string

	for i := 0; i < iter; i++ {
		combos = GenerateList(combos, codes)
	}

	sort.Sort(ByOrder(combos))
	for j := 0; j < len(combos); j++ {
		fmt.Printf("%s\n", combos[j])
	}

}

func GenerateList(combos []string, codes []string) []string {
	if len(combos) < 1 {
		return codes
	}
	var newcombos []string
	for i := 0; i < len(combos); i++ {
		for j := 0; j < len(codes); j++ {
			a := []string{combos[i], codes[j]}
			newcombos = append(newcombos, strings.Join(a, ""))
		}
	}
	return newcombos
}
