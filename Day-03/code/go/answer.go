package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func processMultiplicationsBasic(text string) int {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := pattern.FindAllStringSubmatch(text, -1)
	sum := 0

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		if x >= 1 && x <= 999 && y >= 1 && y <= 999 {
			sum += x * y
		}
	}

	return sum
}

func processMultiplicationsWithControl(text string) int {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	sum := 0
	chars := strings.Split(text, "")
	multiplyEnabled := true
	i := 0

	for i < len(chars) {
		// Check for "don't" first - must be checked before "do"
		if i+4 < len(chars) && strings.Join(chars[i:i+5], "") == "don't" {
			multiplyEnabled = false
			i += 5
			continue
		}

		// Check for standalone "do" (not part of "don't")
		if i+1 < len(chars) && strings.Join(chars[i:i+2], "") == "do" {
			// Make sure this "do" is not part of "don't"
			if i == 0 || strings.Join(chars[i-3:i], "") != "don" {
				multiplyEnabled = true
			}
			i += 2
			continue
		}

		// Only process multiplication if enabled
		if multiplyEnabled && chars[i] == "m" {
			end := i + 15
			if end > len(chars) {
				end = len(chars)
			}
			substring := strings.Join(chars[i:end], "")
			matches := pattern.FindStringSubmatch(substring)

			if matches != nil {
				x, _ := strconv.Atoi(matches[1])
				y, _ := strconv.Atoi(matches[2])
				if x >= 1 && x <= 999 && y >= 1 && y <= 999 {
					sum += x * y
					i += len(matches[0]) - 1
				}
			}
		}

		i++
	}

	return sum
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run script.go <input_file>")
		os.Exit(1)
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	text := string(content)
	basicSum := processMultiplicationsBasic(text)
	controlledSum := processMultiplicationsWithControl(text)

	fmt.Printf("Basic sum (without do/don't): %d\n", basicSum)
	fmt.Printf("Controlled sum (with do/don't): %d\n", controlledSum)
}
