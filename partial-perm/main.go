package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var min int64

func Factorial(val *big.Int) *big.Int {

	if val.Cmp(big.NewInt(min)) < 0 {
		return big.NewInt(1)
	}
	lessOne := big.NewInt(0)
	lessOne.Sub(val, big.NewInt(1))
	toReturn := big.NewInt(0)
	return toReturn.Mul(val, Factorial(lessOne))
}

func main() {
	filename := "rosalind_pper.txt"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error is %s", err.Error())
		os.Exit(1)
	}

	input := string(b)
	lines := strings.Split(input, "\n")[0]

	tokens := strings.Split(lines, " ")
	fmt.Println(tokens)

	totalVals, _ := strconv.ParseInt(tokens[0], 10, 64)
	numSelections, _ := strconv.ParseInt(tokens[1], 10, 64)

	min = totalVals - numSelections + 1

	result := Factorial(big.NewInt(totalVals))
	//fmt.Println(result)

	var mod big.Int

	mod.Mod(result, big.NewInt(1000000))

	fmt.Println(&mod)
}
