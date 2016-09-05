package sort_test

import "testing"

func TestMergesort(t *testing.T) {
    if sort.Mergesort([]int{3, 1, 2}) != []int{1, 2, 3} {
        t.Fail()
    }
}
