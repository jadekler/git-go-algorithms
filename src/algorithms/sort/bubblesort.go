package sort

func Bubblesort(arr []int) []int {
    var swapped bool
    var hold int

    for i := 0; i < len(arr) - 1; i++ {
        if arr[i] > arr[i+1] {
            swapped = true
            hold = arr[i]
            arr[i] = arr[i+1]
            arr[i+1] = hold
        }
    }

    if swapped {
        return Bubblesort(arr)
    }

    return arr
}