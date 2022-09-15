package main

import (
    "fmt"
    "practice/tree"
)

func main() {
    /*test := []int{8, 4, 2, 7, 6, 9, 10, 12, 13, 14, 11, 15, 18}
      // test1 := []int{1, 1, 1, 1, 1, 1, 1, 1}
      // res := InsertionSort(test)
      // res := InsertionSortDes(test)
      // res := MergeSort(test, 0, len(test)-1)
      // res := BubbleSort(test)
      // res := ReversePairs(test)
      // res := MaxHeapify1(test, 1)
      // res := BuildMaxHeap(test)
      // res := BuildMinHeap(test)
      // res := HeapSort(test)
      // res := QuickSort(test, 0, len(test)-1)
      // res := RandomQuickSort(test, 0, len(test)-1)
      // res := QuickSort2(test, 0, len(test)-1)
      // res := CountingSort1(test, 18)
      // res := BucketSort(test, 10)
      // res := RadixSort(test)
      // max, min := GetMaxAndMin(test)
      // SelectSort(test)
      // ShellSort(test, 2)
      // res := RandomIntSlice(15, 1000)
      // fmt.Println(res)
      // temp := RandomSelect(res, 0, len(res)-1, 3)
      // ShellSort(res, 2)
      // fmt.Println(temp)
      res := tree.CreateBinarySearchTree(test)
      // Inorder(res)
      tree.Delete(res, 10)
      // remove(res, 10)
      tree.Inorder(res)*/

    src := []int{11, 2, 14, 1, 7, 15, 5, 8}
    tree.CreateRBT(src, len(src))
    tree.Traversal(tree.PRBT)
    fmt.Println("===================")
    tree.DeleteNode(tree.PRBT, 1)
    tree.Traversal(tree.PRBT)
    fmt.Println("===================")

    tree.DeleteNode(tree.PRBT, 7)
    tree.Traversal(tree.PRBT)
    fmt.Println("===================")

    tree.DeleteNode(tree.PRBT, 14)
    tree.Traversal(tree.PRBT)
    fmt.Println("===================")

    tree.DeleteNode(tree.PRBT, 15)
    tree.Traversal(tree.PRBT)
    fmt.Println("===================")
}
