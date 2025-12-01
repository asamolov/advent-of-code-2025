package utils

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInput() []string {
	var path string
	if len(os.Args) == 2 {
		path = os.Args[1]
	} else {
		path = "input.txt"
	}
	fmt.Printf("reading from input file: %s\n", path)

	fp, err := os.Open(path)
	check(err)
	defer fp.Close()
	scanner := bufio.NewScanner(fp)

	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	check(scanner.Err())
	return result
}
