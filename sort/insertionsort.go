package sort

func Insertionsort(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        var j int
        elem := arr[i]

        for j = i; j > 0 && elem < arr[j - 1]; j-- {

        }

        arr = shiftup(arr, j, i)
        arr[j] = elem
    }

    return arr
}

func shiftup(arr []int, start, end int) []int {
    for i := end; i > start; i-- {
        arr[i] = arr[i - 1]
    }

    return arr
}