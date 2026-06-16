package main

import (
	"bufio"
	"bytes"
	"fmt"

	_ "github.com/AlecAivazis/survey/v2"
)

func main() {

	out, err := PendingFiles()
	if err != nil {
		fmt.Println(err)
	}

	options := parseString(out)

	selected, prefix, message, err := UserInput(options)
	if err != nil {
		fmt.Println(err)
		return
	}

	GitProcess(selected, prefix, message)

}

func parseString(input []byte) (files []string) {

	scanner := bufio.NewScanner(bytes.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		filename := line[3:]
		files = append(files, filename)
	}

	return files
}
