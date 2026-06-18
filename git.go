package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func (c *Commit) Execute() error {
	if err := gitAdd(c.Files); err != nil {
		return err
	}
	fmt.Println("Files added successfully!")
	if err := gitCommitMessage(c.Type, c.Message); err != nil {
		return err
	}

	fmt.Println("Message commited!")
	if err := gitPush(); err != nil {
		return err
	}

	fmt.Println("Files pushed!")
	return nil
}

func gitAdd(selected []string) error {

	fmt.Println("Adding files...")

	args := append([]string{"add"}, selected...)
	cmd := exec.Command("git", args...)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("Git adding fail! Here is why: %s\n", string(out))
	}

	return nil
}

func gitCommitMessage(prefix, message string) error {

	fmt.Println("Commiting message...")
	commitMsg := fmt.Sprintf("%s: %s", prefix, message)
	cmd := exec.Command("git", "commit", "-m", commitMsg)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("Git adding commit message fail! Here is why: %s\n", string(out))
	}

	return nil

}

func gitPush() error {

	fmt.Println("Pushing files...")
	cmd := exec.Command("git", "push")
	out, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("Git push fail! Here is why: %s\n", string(out))
	}

	return nil

}

func PendingFiles() (out []byte, err error) {
	cmd := exec.Command("git", "status", "--porcelain")

	out, err = cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Command fail")
	}

	return out, nil
}

func GitBranch() (string, error) {
	fmt.Println("Finding branch...")
	cmd := exec.Command("git", "branch", "--show-current")
	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("Branch find fail! Here is why: %s\n", string(out))
	}

	return string(bytes.TrimSpace(out)), nil
}
