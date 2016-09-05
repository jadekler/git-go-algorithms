package sort_test

import "testing"

func TestQuicksort(t *testing.T) {
    if sort.Quicksort([]int{3, 1, 2}) != []int{1, 2, 3} {
        t.Fail()
    }
}
