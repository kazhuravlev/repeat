package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("not enough argument")
		os.Exit(1)
	}

	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("first argument should be repeat count")
		os.Exit(1)
	}

	if count < 0 {
		fmt.Println("count argument should be greater that zero")
		os.Exit(1)
	}

	tail := os.Args[2:]
	for i := 0; i < count; i++ {
		cmd := exec.Command(tail[0], tail[1:]...) //nolint:gosec
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("run command", err.Error())
			os.Exit(1)
		}
	}
}
