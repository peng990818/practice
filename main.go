package main

import (
    "fmt"
)

func main() {
    //test := []int{8, 6, 2, 9, 10, 12, 13, 14, 11, 15, 18}
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
    //res := RadixSort(test)
    //max, min := GetMaxAndMin(test)
    //SelectSort(test)
    //ShellSort(test, 2)
    res := RandomIntSlice(15, 1000)
    fmt.Println(res)
    temp := RandomSelect(res, 0, len(res)-1, 3)
    //ShellSort(res, 2)
    fmt.Println(temp)
}
