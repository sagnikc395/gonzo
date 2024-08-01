package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Builtin command functions
func gonzoCd(args []string) int {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "lsh: expected argument to \"cd\"")
	} else {
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "lsh:", err)
		}
	}
	return 1
}

func gonzoHelp(args []string) int {
	fmt.Println("gonzo shell 😤")
	fmt.Println("Type program names and arguments, and hit enter.")
	fmt.Println("The following are built in:")
	for _, cmd := range builtinStr {
		fmt.Println("  ", cmd)
	}
	fmt.Println("Use the man command for information on other programs.")
	return 1
}

func gonzoExit(args []string) int {
	return 0
}

// Builtin command maps
var builtinStr = []string{
	"cd",
	"help",
	"exit",
}

var builtinFunc = map[string]func([]string) int{
	"cd":   gonzoCd,
	"help": gonzoHelp,
	"exit": gonzoExit,
}

func gonzoLaunch(args []string) int {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "lsh:", err)
	}
	return 1
}

func gonzoExecute(args []string) int {
	if len(args) == 0 {
		return 1
	}

	if fn, ok := builtinFunc[args[0]]; ok {
		return fn(args)
	}

	return gonzoLaunch(args)
}

func gonzoReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "lsh: error reading input:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(line)
}

func gonzoSplitLine(line string) []string {
	return strings.Fields(line)
}

func gonzoLoop() {
	for {
		fmt.Print("> ")
		line := gonzoReadLine()
		args := gonzoSplitLine(line)
		status := gonzoExecute(args)

		if status == 0 {
			break
		}
	}
}

func main() {
	// Load config files, if any.

	// Run command loop.
	gonzoLoop()

	// Perform any shutdown/cleanup.
}
