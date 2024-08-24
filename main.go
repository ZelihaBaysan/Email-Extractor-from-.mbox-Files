package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// ExtractEmails scans through an mbox file, extracts all email addresses, and writes them to an output.txt file.
//
// This function is useful for retrieving email addresses from a backup .mbox file. You can use this to see where
// your emails have been sent or to track where you have registered with your email address.
//
// Parameters:
//
// - mboxFilePath: the path to the mbox file you want to scan for email addresses (e.g., "/your/path/emails.mbox").
//
// - outputFilePath: the path where the output file containing the email addresses will be saved (e.g., "output.txt").
//
// Example usage:
//
// ExtractEmails("/your/path/emails.mbox", "output.txt")
func main() {
	// Open the mbox file
	mboxFile, err := os.Open("/your/path/emails.mbox")
	if err != nil {
		log.Fatalf("Mbox file could not be opened: %v", err)
	}
	defer mboxFile.Close()

	// Create the output.txt file
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Output file could not be created: %v", err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(mboxFile)

	// Regex pattern to match email addresses
	emailPattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	for scanner.Scan() {
		line := scanner.Text()

		// Find all email addresses in the current line
		emails := emailPattern.FindAllString(line, -1)

		// Write each email address to the output.txt file
		for _, email := range emails {
			fmt.Fprintln(outputFile, email)
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading mbox file: %v", err)
	}

	fmt.Println("Email addresses have been saved to output.txt")
}
