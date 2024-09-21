package piscine

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Function to run a quad executable and capture its output
func runQuadExecutable(quadName string, cols, rows string) (string, error) {
	exePath, err := filepath.Abs(quadName)
	if err != nil {
		return "", fmt.Errorf("error getting absolute path for %s: %v", quadName, err)
	}

	cmd := exec.Command(exePath, cols, rows)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running %s: %v", quadName, err)
	}
	return out.String(), nil
}

func main() {
	// Check if arguments are provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./quadchecker <cols> <rows>")
		return
	}

	cols := os.Args[1]
	rows := os.Args[2]

	// Reading from stdin (the piped input)
	var input bytes.Buffer
	_, err := input.ReadFrom(os.Stdin)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	pipedOutput := input.String()

	// List of quad executables to compare against
	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}

	// Compare outputs with the quads
	var matches []string
	for _, quad := range quads {
		output, err := runQuadExecutable(quad, cols, rows)
		if err != nil {
			fmt.Printf("Error running %s: %v\n", quad, err)
			continue
		}

		if strings.TrimSpace(output) == strings.TrimSpace(pipedOutput) {
			matches = append(matches, fmt.Sprintf("[%s] [%s] [%s]", quad, cols, rows))
		}
	}

	// Print the result in the required format
	if len(matches) > 0 {
		fmt.Println(strings.Join(matches, " || "))
	}
}
