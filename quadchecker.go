package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Step 1: Read the input from stdin (the piped output from one of the quad executables)
	fmt.Println("Reading piped input...")
	input := new(bytes.Buffer)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Step 2: Define the quad executables to compare against
	quadExecutables := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	matchFound := false

	// Step 3: Loop through each quad executable and compare outputs
	for _, quad := range quadExecutables {
		// Call the quad executable with the same dimensions that were piped in
		fmt.Printf("Checking %s...\n", quad)
		cmd := exec.Command(quad, "1", "1") // Assuming dimensions are 1x1, adjust as needed
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to run %s: %v\n", quad, err)
			continue
		}

		// Compare the output of the quad executable with the piped input
		if out.String() == input.String() {
			fmt.Printf("Match found with %s!\n", quad)
			matchFound = true
		}
	}

	// Step 4: If no match is found
	if !matchFound {
		fmt.Println("No match found with any quad.")
	}
}
