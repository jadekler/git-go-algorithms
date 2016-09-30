package sort_test

import (
    "testing"
    "algorithms/sort"
    "reflect"
)

func TestInsertionsort(t *testing.T) {
    actual := sort.Insertionsort([]int{3, 1, 2})
    expected := []int{1, 2, 3}
    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v", expected, actual)
        t.Fail()
    }
}

func TestInsertionsortSame(t *testing.T) {
    actual := sort.Insertionsort([]int{1, 2, 3})
    expected := []int{1, 2, 3}
    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v", expected, actual)
        t.Fail()
    }
}

func TestInsertionsortComplex(t *testing.T) {
    actual := sort.Insertionsort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
    expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v", expected, actual)
        t.Fail()
    }
}
