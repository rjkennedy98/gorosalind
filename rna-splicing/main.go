package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := "rosalind_splc.txt"

	chartMap := GetChartMap()

	//fmt.Printf("%v\n", chartMap)

	f, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("error occured %s", err.Error())
		os.Exit(1)
	}

	input := string(f)

	lines := strings.Split(input, "\n")

	var proteins []string

	buffer := *(bytes.NewBufferString(""))

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, ">") {
			protein := buffer.String()
			if protein != "" {
				proteins = append(proteins, protein)
			}
			buffer = *(bytes.NewBufferString(""))

		} else {
			buffer.WriteString(line)
		}
	}
	proteins = append(proteins, buffer.String())

	//fmt.Printf("list = %v\n", proteins)

	mainprotein := proteins[0]

	//fmt.Printf("mainprotein %s\n", mainprotein)
	for k := 1; k < len(proteins); k++ {
		mainprotein = strings.Join(strings.Split(mainprotein, proteins[k]), "")
		//fmt.Printf("mainprotein %s\n", mainprotein)
	}

	var finalProtArray []string

	mainProtArray := strings.Split(mainprotein, "")
	for j := 0; j < len(mainProtArray); j += 3 {
		code := strings.Join(mainProtArray[j:(j+3)], "")
		prot := chartMap[code]
		//fmt.Printf("code = %s, prot=%s\n", code, prot)
		if prot == "Stop" {
			break
		}
		finalProtArray = append(finalProtArray, prot)
	}

	fmt.Printf("\n%s\n\n", strings.Join(finalProtArray, ""))
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
	//fmt.Printf("chart %v\n", chart)
	lines := strings.Split(chart, "\n")

	for i := 0; i < len(lines); i++ {
		tokens := strings.Split(lines[i], "|")
		if len(tokens) < 2 {
			continue
		}
		//fmt.Printf("line is %s %s\n", tokens[0], tokens[1])
		protein := tokens[0]
		seq := strings.Split(tokens[1], ",")
		for k := 0; k < len(seq); k++ {
			chartMap[(seq[k])] = protein
		}

	}

	return chartMap
}
