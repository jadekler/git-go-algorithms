package sort_test

import (
    "testing"
)

func TestInsertionsort(t *testing.T) {
    if sort.Insertionsort([]int{3, 1, 2}) != []int{1, 2, 3} {
        t.Fail()
    }
}
