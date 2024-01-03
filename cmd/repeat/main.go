package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
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
	var delay time.Duration
	var idxShift int
	if dur, err := time.ParseDuration(tail[0]); err == nil {
		delay = dur
		idxShift = 1
	}

	for i := 0; i < count; i++ {
		if i != 0 {
			time.Sleep(delay)
		}

		cmd := exec.Command(tail[0+idxShift], tail[1+idxShift:]...) //nolint:gosec
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err.Error())
		}
	}
}
