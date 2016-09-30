package problems

import "fmt"

func SpiralPrint(arr [][]int) []int {
    var res []int

    var rowStart, rowEnd, columnStart, columnEnd, index int
    rowEnd = len(arr[0]) - 1
    columnEnd = len(arr) - 1

    fmt.Println(rowStart, rowEnd, columnStart, columnEnd, index)

    for rowStart <= rowEnd && columnStart <= columnEnd {
        for index = rowStart; index < rowEnd; index++ {
            res = append(res, arr[columnStart][index])
        }

        for index = columnStart; index < columnEnd; index++ {
            res = append(res, arr[index][rowEnd ])
        }

        for index = rowEnd; index >= rowStart; index-- {
            res = append(res, arr[columnEnd][index])
        }

        for index = columnEnd - 1; index >= columnStart + 1; index-- {
            res = append(res, arr[index][rowStart])
        }

        rowStart++
        columnStart++
        rowEnd--
        columnEnd--
    }

    return res
}