package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := "rosalind_long.txt"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error is %s", err.Error())
		os.Exit(1)
	}

	input := string(b)
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

	// fmt.Println(proteins)

	for len(proteins) > 1 {
		// fmt.Println(proteins)
		found := false
		first := proteins[0]

		for j := 1; j < len(proteins); j++ {

			compare := proteins[j]
			half := GetOverlap(first, compare)
			if half != "" {
				// fmt.Printf("first=%s overlap=%s\n", first, half)
				if strings.HasPrefix(first, half) {
					// fmt.Print("hasPrefix\n")
					pre := strings.Split(compare, half)[0]
					newProtein := fmt.Sprint(pre, first)
					CreateNewList(newProtein, j, &proteins)
				}
				if strings.HasSuffix(first, half) {
					// fmt.Print("hasSuffix\n")
					post := strings.Split(compare, half)[1]
					newProtein := fmt.Sprint(first, post)
					CreateNewList(newProtein, j, &proteins)
				}

				found = true

				//fmt.Print("Should not reach here\n")
			}
			if found {
				break
			}
		}

	}

	fmt.Println(proteins[0])
}
func GetOverlap(first string, compare string) string {
	tokens := strings.Split(compare, "")
	mid := int((len(tokens) / 2))
	firsthalf := strings.Join(tokens[:mid], "")
	laterhalf := strings.Join(tokens[mid:], "")
	//fmt.Printf("major protein: %s\n", first)
	//fmt.Printf("firsthalf=%s laterhalf=%s\n", firsthalf, laterhalf)
	if strings.Contains(first, firsthalf) {

		for i := mid; i < len(tokens); i++ {
			half := strings.Join(tokens[:i], "")
			if strings.HasSuffix(first, half) {
				return half
			}
		}

	}
	if strings.Contains(first, laterhalf) {
		for i := 0; i < mid; i++ {
			half := strings.Join(tokens[i:], "")
			if strings.HasPrefix(first, half) {
				return half
			}
		}

	}
	return ""
}
func CreateNewList(first string, skipIndex int, array *[]string) {
	(*array)[0] = first
	(*array)[skipIndex] = (*array)[len(*array)-1]
	(*array) = (*array)[:len(*array)-1]
	// for j := 1; j < len(thisArray); j++ {
	// 	if j != skipIndex {
	// 		newArray = append(newArray, thisArray[j])
	// 	}
	// }
	// *array = newArray
	// fmt.Printf("newArray %v\n", array)
}
