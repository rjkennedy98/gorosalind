package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := "rosalind_test.txt"

	chartMap := GetChartMap()

	fmt.Printf("%v\n", chartMap)

}

func GetChartMap() map[string]string {

	chartMap := make(map[string]string)

	chartFilename := "dnacodons.txt"
	bytes, err := ioutil.ReadFile(chartFilename)
	if err != nil {
		fmt.Printf("error occured %s", err.Error())
		os.Exit(1)
	}
	chart := string(bytes)
	fmt.Printf("chart %v\n", chart)
	lines := strings.Split(chart, "\n")

	for i := 0; i < len(lines); i++ {
		tokens := strings.Split(lines[i], "|")
		if len(tokens) < 2 {
			continue
		}
		fmt.Printf("line is %s %s\n", tokens[0], tokens[1])
		protein := tokens[0]
		seq := strings.Split(tokens[1], ",")
		for k := 0; k < len(seq); k++ {
			chartMap[(seq[k])] = protein
		}

	}

	return chartMap
}
