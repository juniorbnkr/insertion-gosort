# Insertion Sort Algorithm

This repository contains an implementation of the Insertion Sort algorithm in Go. The Insertion Sort algorithm efficiently sorts an array of integers in ascending order. This README provides an overview of the algorithm, instructions for running the code, and additional details for understanding and using the implementation.

## Algorithm Overview

The Insertion Sort algorithm follows the following steps:

1. Start with the second element in the array and iterate over all elements from left to right.
2. For each element, compare it with elements to its left and shift them to the right until the correct position is found.
3. Repeat this process until the entire array is sorted.

## Code Structure

The code is organized as follows:

- `insertion_sort.go`: Contains the implementation of the Insertion Sort algorithm in Go.
- `main.go`: Demonstrates the usage of the Insertion Sort algorithm on a sample array.

## Usage

To use the Insertion Sort algorithm, follow these steps:

1. Clone the repository to your local machine.
2. Open a terminal and navigate to the directory containing the code.
3. Run the following command to build and execute the code:

`go run main.go`

4. The program will output the unsorted array and the sorted array.

## Example

Here is an example usage of the Insertion Sort algorithm:

```go
package main

import (
 "fmt"
)

func main() {
 items := []int{5, 2, 8, 9, 1}
 insertionsort(items)
 fmt.Println("Sorted array:", items)
}

```

## Complexity Analysis
The time complexity of the Insertion Sort algorithm is O(n^2), where n is the number of elements in the array. This is because, in the worst-case scenario, each element needs to be compared with all previous elements. The space complexity of the algorithm is O(1) since the sorting is done in-place without requiring any additional memory.
