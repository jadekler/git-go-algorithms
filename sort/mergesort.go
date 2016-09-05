package sort

import "fmt"

func Mergesort(arr []int) []int {
    return mergesortActual(arr, pivotPointFn(arr))
}

func mergesortActual(arr []int, pivotPoint int) []int {
    if pivotPoint > 0 {
        left := arr[0:pivotPoint]
        right := arr[pivotPoint:]

        left = mergesortActual(left, pivotPointFn(left))
        right = mergesortActual(right, pivotPointFn(right))

        arr = merge(left, right)
    } else {
        return arr
    }

    return arr
}

func merge(left, right []int) []int {
    var newArr []int

    leftIndex := 0
    rightIndex := 0

    for leftIndex < len(left) || rightIndex < len(right) {
        if leftIndex >= len(left) {
            newArr = append(newArr, right[rightIndex])
            rightIndex++
        } else if rightIndex >= len(right) {
            newArr = append(newArr, left[leftIndex])
            leftIndex++
        } else if left[leftIndex] < right[rightIndex] {
            newArr = append(newArr, left[leftIndex])
            leftIndex++
        } else {
            newArr = append(newArr, right[rightIndex])
            rightIndex++
        }
    }

    fmt.Println(leftIndex, left, rightIndex, right, newArr)

    return newArr
}