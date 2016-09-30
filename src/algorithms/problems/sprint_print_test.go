package problems_test

import (
    "testing"
    "git-go-algorithms/problems"
    "reflect"
)

// 1
func TestSpiralPrintSingleElem(t *testing.T) {
    expected := []int{1}
    actual := problems.SpiralPrint([][]int{[]int{1}})

    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v\n", expected, actual)
        t.Fail()
    }
}

// 3 2 1
func TestSpiralPrintSingleRow(t *testing.T) {
    expected := []int{3, 2, 1}
    actual := problems.SpiralPrint([][]int{[]int{3, 2, 1}})

    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v\n", expected, actual)
        t.Fail()
    }
}

// 3
// 2
// 1
func TestSpiralPrintSingleColumn(t *testing.T) {
    expected := []int{3, 2, 1}
    actual := problems.SpiralPrint([][]int{[]int{3}, []int{2}, []int{1}})

    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v\n", expected, actual)
        t.Fail()
    }
}

// [
//  [ 3 2 1 ]
//  [ 6 5 4 ]
//           ]
func TestSpiralPrintTwoRow(t *testing.T) {
    expected := []int{3, 2, 1, 4, 5, 6}
    actual := problems.SpiralPrint([][]int{[]int{3, 2, 1}, []int{6, 5, 4}})

    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v\n", expected, actual)
        t.Fail()
    }
}

// [
//  [ 3 6 ]
//  [ 2 5 ]
//  [ 1 4 ]
//         ]
func TestSpiralPrintTwoColumn(t *testing.T) {
    expected := []int{3, 6, 5, 4, 1, 2}
    actual := problems.SpiralPrint([][]int{[]int{3, 6}, []int{2, 5}, []int{1, 4}})

    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v\n", expected, actual)
        t.Fail()
    }
}

// [
//  [ 3 2 1 ]
//  [ 4 5 6 ]
//  [ 7 8 9 ]
//           ]
func TestSpiralPrintSquare(t *testing.T) {
    expected := []int{3, 2, 1, 6, 9, 8, 7, 4, 5}
    actual := problems.SpiralPrint([][]int{[]int{3, 2, 1}, []int{4, 5, 6}, []int{7, 8, 9}})

    if !reflect.DeepEqual(actual, expected) {
        t.Logf("Expected %v, got %v\n", expected, actual)
        t.Fail()
    }
}
