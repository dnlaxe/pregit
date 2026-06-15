package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "status")

	fmt.Println(cmd)
}
