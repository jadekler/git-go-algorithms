package problems

import "fmt"

func SpiralPrint(arr [][]int) []int {
    var res []int

    var rowStart, rowEnd, columnStart, columnEnd, index int
    rowEnd = len(arr[0])
    columnEnd = len(arr)

    fmt.Println(rowStart, rowEnd, columnStart, columnEnd, index)

    for rowStart < rowEnd && columnStart < columnEnd {
        for index = rowStart; index < rowEnd; index++ {
            res = append(res, arr[columnStart][index])
        }

        for index = columnStart + 1; index < columnEnd; index++ {
            res = append(res, arr[index][rowEnd - 1])
        }

        if columnEnd - columnStart > 1 {
            for index = rowEnd - 2; index >= rowStart; index-- {
                res = append(res, arr[columnEnd - 1][index])
            }
        }

        if rowEnd - rowStart > 1 {
            for index = columnEnd - 2; index >= columnStart + 1; index-- {
                res = append(res, arr[index][rowStart])
            }
        }

        rowStart++
        rowEnd--
        columnStart++
        columnEnd--
    }

    return res
}