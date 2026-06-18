package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	_ "github.com/AlecAivazis/survey/v2"
)

const Version = "dev"

func main() {

	out, err := PendingFiles()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	branch, err := GitBranch()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	options := parseOptions(out)

	c, err := UserInput(options, branch)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.Execute()
	if err != nil {
		fmt.Println("Process failed:", err)
	}

}

func parseOptions(input []byte) (files []string) {

	scanner := bufio.NewScanner(bytes.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		filename := line[3:]
		files = append(files, filename)
	}

	return files
}
