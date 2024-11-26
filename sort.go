package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func ceilDiv(a int, b int) int {
	return (a + b - 1) / b
}

func divideToChunks(numChunks int, arr []int) [][]int {
	var returnArr [][]int
	chunkLength := ceilDiv(len(arr), numChunks)
	for i := 0; i < len(arr); i += chunkLength {
		endIdx := i + chunkLength
		if endIdx > len(arr) {
			endIdx = len(arr)
		}
		returnArr = append(returnArr, arr[i:endIdx])
	}
	return returnArr
}

func sortInts(arr []int, c chan []int, wg *sync.WaitGroup) {
	fmt.Println("Sorting subarray: ", arr)
	sort.Ints(arr)
	c <- arr
	wg.Done()
}

func mergeSortTwoSlices(arr1 []int, arr2 []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}
	return result
}

func recursiveMergeSort(slices [][]int) []int {
	if len(slices) == 1 {
		return slices[0]
	}

	var mergedSlices [][]int
	for i := 0; i < len(slices); i += 2 {
		if i+1 < len(slices) {
			merged := mergeSortTwoSlices(slices[i], slices[i+1])
			mergedSlices = append(mergedSlices, merged)
		} else {
			mergedSlices = append(mergedSlices, slices[i])
		}
	}

	return recursiveMergeSort(mergedSlices)
}

func handleInputError() {
	fmt.Println("Invalid input, you must enter 4 integers separated by spaces. Exting program.")
	os.Exit(1)
}

func getUserInput() []int {
	fmt.Println("Input at least 4 numbers separated by spaces: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()
	userInputArr := strings.Fields(userInput)
	if len(userInputArr) < 4 {
		handleInputError()
	}
	var inputNumbers []int
	for _, str := range userInputArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			handleInputError()
		}
		inputNumbers = append(inputNumbers, num)

	}

	return inputNumbers
}

func main() {
	var wg sync.WaitGroup
	userInputSlice := getUserInput()
	chunks := divideToChunks(4, userInputSlice)
	c := make(chan []int, 4)

	for _, chunk := range chunks {
		wg.Add(1)
		go sortInts(chunk, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	var sortedArrays [][]int
	for subArray := range c {
		sortedArrays = append(sortedArrays, subArray)
	}
	sortedAll := recursiveMergeSort(sortedArrays)
	fmt.Println(sortedAll)
}
