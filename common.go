package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetLineFromStdin(prompt string) string {
    // fmt.Scanln breaks on any whitespace, so use a bufio.Scanner instead
    // This correctly handles backspace so user can edit inline before hitting enter
    fmt.Printf(prompt)
    scanner := bufio.NewScanner(os.Stdin)
    var line string
    if scanner.Scan() {
        line = scanner.Text()
        // fmt.Printf("Input was: %q\n", line)
    }
    return line
}

func getMasterPassphrase() string {
    passPhrase := os.Getenv("SECRET_BOX")
    if passPhrase == "" {
        fmt.Println("missing SECRET_BOX env var")
        passPhrase = GetLineFromStdin("Enter secret box passphrase -> ")
        fmt.Println()
    }
	return passPhrase
}
