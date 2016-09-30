package problems

func SpiralPrint(arr [][]int) []int {
    var res []int

    width := len(arr[0])
    height := len(arr)
    totalElems := width * height

    for i := 0; i < totalElems; i++ {
        row := i / width
        column := i - (row * width)

        res = append(res, arr[row][column])
    }

    return res
}