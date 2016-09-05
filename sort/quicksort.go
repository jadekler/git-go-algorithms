package sort

import "fmt"

func Quicksort(arr []int) []int {
    var hold int
    pivotPoint := pivotPointFn(arr)

    fmt.Println(arr)

    for i := 0; i < pivotPoint; i++ {
        if arr[i] > arr[pivotPoint] {
            hold = arr[i]
            pivotPoint = i
            arr[i] = arr[pivotPoint]
            arr[pivotPoint] = hold
        }
    }

    if len(arr) > 0 {
        left := Quicksort(arr[0:pivotPoint])
        right := Quicksort(arr[pivotPoint:])

        return append(left, right...)
    } else {
        return arr
    }
}
