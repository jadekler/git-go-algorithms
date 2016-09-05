package sort_test

import (
    "testing"
)

func TestHeapsort(t *testing.T) {
    if sort.Heapsort([]int{3, 1, 2}) != []int{1, 2, 3} {
        t.Fail()
    }
}
