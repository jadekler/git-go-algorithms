package sort_test

import (
    "testing"
    "git-go-algorithms/sort"
    "reflect"
)

func TestBubblesort(t *testing.T) {
    actual := sort.Bubblesort([]int{3, 1, 2})
    expected := []int{1, 2, 3}
    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v", expected, actual)
        t.Fail()
    }
}
