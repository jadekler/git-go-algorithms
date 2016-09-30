package sort_test

import (
    "testing"
    "git-go-algorithms/sort"
    "reflect"
)

func TestQuicksort(t *testing.T) {
    actual := sort.Quicksort([]int{3, 1, 2})
    expected := []int{1, 2, 3}
    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v", expected, actual)
        t.Fail()
    }
}

func TestQuicksortSame(t *testing.T) {
    actual := sort.Quicksort([]int{1, 2, 3})
    expected := []int{1, 2, 3}
    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v", expected, actual)
        t.Fail()
    }
}

func TestQuicksortComplex(t *testing.T) {
    actual := sort.Quicksort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
    expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v", expected, actual)
        t.Fail()
    }
}