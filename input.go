package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

type Commit struct {
	Branch  string
	Files   []string
	Type    string
	Message string
}

func UserInput(options []string, branch string) (c Commit, err error) {

	c = Commit{
		Branch:  branch,
		Files:   selectFiles(options),
		Type:    selectPrefix(),
		Message: typeMessage(),
	}

	confirmMsg := fmt.Sprintf("Branch: %v\n  Files: %v\n  Commit message: %s: %s\n", c.Branch, c.Files, c.Type, c.Message)

	if confirmation(confirmMsg) {
		fmt.Println("Starting process...")
		return c, nil
	} else {
		return Commit{}, fmt.Errorf("Cancelled")
	}

}

func selectFiles(options []string) (selected []string) {

	prompt := &survey.MultiSelect{
		Message: "Which files do you want to add?",
		Options: options,
	}

	survey.AskOne(prompt, &selected)

	fmt.Printf("You choose to add: %v\n", selected)

	return selected
}

func selectPrefix() (prefix string) {

	prefixes := []string{"feat", "ci", "fix", "docs", "style", "refactor", "chore", "test", "deletion"}

	prompt := &survey.Select{
		Message: "Choose a prefix",
		Options: prefixes,
	}

	survey.AskOne(prompt, &prefix)

	return prefix
}

func typeMessage() (message string) {
	prompt := &survey.Input{
		Message: "Type your commit message",
	}

	survey.AskOne(prompt, &message)

	return message
}

func confirmation(message string) (confirmation bool) {
	prompt := &survey.Confirm{
		Message: message,
		Default: true,
	}

	survey.AskOne(prompt, &confirmation)

	return confirmation
}
