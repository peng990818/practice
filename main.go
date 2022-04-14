package main

import (
    "fmt"
)

func main() {
    test := []int{8, 6, 2, 9, 10, 12, 13, 14, 11, 15, 18}
    //test1 := []int{1, 1, 1, 1, 1, 1, 1, 1}
    //res := InsertionSort(test)
    //res := InsertionSortDes(test)
    //res := MergeSort(test, 0, len(test)-1)
    //res := BubbleSort(test)
    //res := ReversePairs(test)
    //res := MaxHeapify1(test, 1)
    //res := BuildMaxHeap(test)
    //res := BuildMinHeap(test)
    //res := HeapSort(test)
    //res := QuickSort(test, 0, len(test)-1)
    //res := RandomQuickSort(test, 0, len(test)-1)
    //res := QuickSort2(test, 0, len(test)-1)
    //res := CountingSort2(test, 18)
    //res := BucketSort(test, 10)
    res := RadixSort(test)
    fmt.Println(res)
}
