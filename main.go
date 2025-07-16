package main

import "fmt"

func quickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	pivotIndex := len(data) / 2
	pivot := data[pivotIndex]

	var left, right []int

	for i, v := range data {
		if i == pivotIndex {
			continue
		}
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{9, 3, 7, 4, 2, 8, 1, 6, 5}
	sorted := quickSort(arr)
	fmt.Println("Отсортированный массив:", sorted)
}
