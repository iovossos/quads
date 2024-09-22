package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Step 1: Read the input from stdin (the piped output from one of the quad executables)
	input := new(bytes.Buffer)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Extract the dimensions dynamically from the input (piped input in this case)
	// You can assume the dimensions are captured from the arguments of the quad executable
	args := strings.Fields(os.Getenv("CMDLINE")) // Fetches the original command arguments passed to the quad executables
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: Unable to extract dimensions from the input.")
		os.Exit(1)
	}

	dimX := args[1] // Extracted width
	dimY := args[2] // Extracted height

	// Step 2: Define the quad executables to compare against
	quadExecutables := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	matches := []string{} // To store the matches

	// Step 3: Loop through each quad executable and compare outputs
	for _, quad := range quadExecutables {
		// Call the quad executable with the same dimensions that were piped in
		cmd := exec.Command(quad, dimX, dimY) // Use the dimensions extracted dynamically
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to run %s: %v\n", quad, err)
			continue
		}

		// Compare the output of the quad executable with the piped input
		if out.String() == input.String() {
			// Add the match in the required format
			matches = append(matches, fmt.Sprintf("[%s] [%s] [%s]", quad, dimX, dimY))
		}
	}

	// Step 4: Output the matches in the desired format
	if len(matches) > 0 {
		fmt.Println(strings.Join(matches, " || "))
	} else {
		fmt.Println("No match found with any quad.")
	}
}
