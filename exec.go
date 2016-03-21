package main

import (
	"os/exec"
	"fmt"
	"io"
	"bufio"
	"os"
)

func f(r io.Reader){
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func main() {
	cmd := exec.Command("ls", "-la")
	outReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Print(err)
		return
	}

	errReader, err := cmd.StderrPipe()
	if err != nil {
		fmt.Print(err)
		return
	}

	go f(outReader)
	go f(errReader)
	if err = cmd.Run(); err != nil {
		fmt.Print(err)
		return
	}
}
