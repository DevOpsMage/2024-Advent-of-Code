package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "log"
)

func isSafe(row []int) bool {
    /**
     * Determines if a row is 'safe'.
     * A row is considered safe if the differences between each pair of
     * successive numbers are all in the range [1, 3] (increasing)
     * or all in the range [-3, -1] (decreasing).
     */
    if len(row) < 2 {
        // Rows with less than 2 elements are considered safe
        return true
    }

    // Calculate the first difference to determine the direction
    firstDiff := row[1] - row[0]

    if 1 <= firstDiff && firstDiff <= 3 {
        // Sequence should be increasing
        for i := 1; i < len(row)-1; i++ {
            diff := row[i+1] - row[i]
            if !(1 <= diff && diff <= 3) {
                return false
            }
        }
        return true
    } else if -3 <= firstDiff && firstDiff <= -1 {
        // Sequence should be decreasing
        for i := 1; i < len(row)-1; i++ {
            diff := row[i+1] - row[i]
            if !(-3 <= diff && diff <= -1) {
                return false
            }
        }
        return true
    } else {
        // Not increasing or decreasing within allowed increments
        return false
    }
}

func readData(filename string) [][]int {
    /**
     * Reads the input data from the specified file.
     * Each line in the file is expected to be a sequence of integers separated by spaces.
     * Returns a slice of slices of integers.
     */
    file, err := os.Open(filename)
    if err != nil {
        log.Fatalf("Failed to open file: %s", err)
    }
    defer file.Close()

    var data [][]int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            continue
        }
        fields := strings.Fields(line)
        var row []int
        for _, field := range fields {
            num, err := strconv.Atoi(field)
            if err != nil {
                log.Fatalf("Failed to parse integer: %s", err)
            }
            row = append(row, num)
        }
        data = append(data, row)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error reading file: %s", err)
    }

    return data
}

func main() {
    // Check if input filename is provided
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run script.go <input_file>")
        os.Exit(1)
    }
    filename := os.Args[1]

    // Read data from the input file
    data := readData(filename)

    // Count the number of safe rows
    safeCount := 0
    for _, row := range data {
        if isSafe(row) {
            safeCount++
        }
    }
    fmt.Println(safeCount)

    // For each row, check if removing any one element makes it safe
    safeCountAfterRemoval := 0
    for _, row := range data {
        isAnySafe := false
        for i := 0; i < len(row); i++ {
            // Create a new row excluding the current element
            newRow := append([]int{}, row[:i]...)
            newRow = append(newRow, row[i+1:]...)
            if isSafe(newRow) {
                isAnySafe = true
                break
            }
        }
        if isAnySafe {
            safeCountAfterRemoval++
        }
    }
    fmt.Println(safeCountAfterRemoval)
}
