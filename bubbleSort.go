package main

import(
	"fmt"
	"strconv"
)

func BubbleSort(nums []int)  {
	numLength := len(nums)
		for i := 0; i < numLength-1; i++ {
			for j :=0; j < numLength-1; j++ {
				if nums[j] > nums[j+1] {
					Swap(nums, j)
				}
			}
		}
}

func Swap(arr []int, j int) {
	arr[j], arr[j+1] = arr[j+1], arr[j]
}

func main() {
	var inputNumbers []int
	var a string
	for integerCount := 0; integerCount < 10; integerCount++ {
		fmt.Println("Enter number: ", integerCount+1, "/10")
		fmt.Scanln(&a)
		inputNum, err := strconv.Atoi(a)
		if err != nil  || a == "" {
			fmt.Println("Error converting entered input to number:", err)
			return
		}
		inputNumbers = append(inputNumbers,inputNum)

	}

	BubbleSort(inputNumbers)
	fmt.Println("The input numbers sorted are:")
	fmt.Println(inputNumbers)
}