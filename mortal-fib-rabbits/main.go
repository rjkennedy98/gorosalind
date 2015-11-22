package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("New Fib: ")
	var buffer bytes.Buffer

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	buffer.WriteString(scanner.Text())

	input := buffer.String()
	fmt.Printf("Received Input %s\n", input)

	first := strings.Split(input, " ")[0]
	second := strings.Split(input, " ")[1]

	iterations, _ := strconv.ParseInt(first, 10, 64)
	months, _ := strconv.ParseInt(second, 10, 64)

	fmt.Printf("Using iterations:%d and months:%d \n", iterations, months)

	fibArray := make([]big.Int, int(months))
	one := big.NewInt(1)
	fibArray[0] = *one
	for i := 1; i <= int(iterations); i++ {
		prettyPrint(fibArray)
		printTotal(fibArray)

		shift(&fibArray)
	}
}

func printTotal(fibArray []big.Int) {
	total := big.NewInt(0)
	for i := 0; i < len(fibArray); i++ {
		(*total).Add(total, &(fibArray[i]))
	}
	fmt.Printf(" total: %s\n", (*total).String())
}

func prettyPrint(fibArray []big.Int) {
	fmt.Printf("[ ")
	for i := 0; i < len(fibArray); i++ {
		fmt.Printf(" %s ", fibArray[i].String())
	}
	fmt.Printf(" ]")
}
func shift(fibArray *[]big.Int) {
	totalChildren := big.NewInt(0)

	for i := 1; i < len(*fibArray); i++ {
		*totalChildren = *totalChildren.Add(totalChildren, &((*fibArray)[i]))
	}

	newFibArray := make([]big.Int, len(*fibArray))
	newFibArray[0] = (*totalChildren)

	for j := 0; j < len(*fibArray)-1; j++ {
		newFibArray[j+1] = (*fibArray)[j]
	}

	(*fibArray) = newFibArray

	//	fmt.Printf("total child: %s\n", (*totalChildren).String())

}
