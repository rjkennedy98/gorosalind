package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

func Factorial(val *big.Int) *big.Int {

	if val.Cmp(big.NewInt(2)) < 0 {
		return big.NewInt(1)
	}
	lessOne := big.NewInt(0)
	lessOne.Sub(val, big.NewInt(1))
	toReturn := big.NewInt(0)
	return toReturn.Mul(val, Factorial(lessOne))
}

func main() {
	filename := "rosalind_pmch.txt"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error is %s", err.Error())
		os.Exit(1)
	}

	input := string(b)
	lines := strings.Split(input, "\n")[1:]

	seq := strings.Join(lines, "")
	fmt.Println(seq)

	aCount := int64(0)
	gCount := int64(0)

	seqArray := strings.Split(seq, "")

	for i := 0; i < len(seqArray); i++ {

		if seqArray[i] == "A" {
			aCount++
		}
		if seqArray[i] == "G" {
			gCount++
		}
	}

	//fmt.Printf("A count=%d\nG count=%d\n", aCount, gCount)

	result := big.NewInt(0)

	result.Mul(Factorial(big.NewInt(aCount)), Factorial(big.NewInt(gCount)))
	fmt.Println(result)
}
