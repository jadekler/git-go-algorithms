package sort_test

import (
    "testing"
)

func TestBubblesort(t *testing.T) {
    if sort.Bubblesort([]int{3, 1, 2}) != []int{1, 2, 3} {
        t.Fail()
    }
}
