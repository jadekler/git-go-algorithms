package sort

import "math"

func Selectionsort(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        var smallest int = math.MaxInt32
        smallestIndex := 0

        for j := i; j < len(arr); j++ {
            if arr[j] < smallest {
                smallest = arr[j]
                smallestIndex = j
            }
        }

        arr[smallestIndex] = arr[i]
        arr[i] = smallest
    }

    return arr
}