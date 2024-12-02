package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    // Check if file path is provided
    if len(os.Args) < 2 {
        fmt.Println("Please provide a file path")
        os.Exit(1)
    }

    // Open file
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()

    var left, right []int
    scanner := bufio.NewScanner(file)

    // Read and parse file lines
    for scanner.Scan() {
        // Split line by three spaces
        parts := strings.Split(scanner.Text(), "   ")
        if len(parts) != 2 {
            fmt.Printf("Invalid line format: %s\n", scanner.Text())
            continue
        }

        nl, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
        nr, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

        if err1 != nil || err2 != nil {
            fmt.Printf("Error converting numbers: %v, %v\n", err1, err2)
            continue
        }

        left = append(left, nl)
        right = append(right, nr)
    }

    // Check for scanning errors
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        os.Exit(1)
    }

    // Part 1: Calculate absolute differences
    sortedLeft := make([]int, len(left))
    sortedRight := make([]int, len(right))
    copy(sortedLeft, left)
    copy(sortedRight, right)
    
    sort.Ints(sortedLeft)
    sort.Ints(sortedRight)

    var p1 []int
    for i := range sortedLeft {
        p1 = append(p1, abs(sortedLeft[i] - sortedRight[i]))
    }
    fmt.Printf("Part 1: %d\n", sum(p1))

    // Part 2: Count frequencies
    c := make(map[int]int)
    for _, n := range right {
        c[n]++
    }

    var p2 []int
    for _, n := range left {
        p2 = append(p2, n * c[n])
    }
    fmt.Printf("Part 2: %d\n", sum(p2))
}

// Helper function for absolute value
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

// Helper function to sum slice
func sum(slice []int) int {
    total := 0
    for _, v := range slice {
        total += v
    }
    return total
}