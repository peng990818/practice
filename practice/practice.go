package main

import (
    "fmt"
    "math/rand"
)

func main() {
    test := []int{8, 6, 2, 9, 10, 12, 22, 13, 14, 11, 15, 18}
    //QuickSort1(test, 0, len(test)-1)
    //QuickSort2(test, 0, len(test)-1)
    //RandomQuickSort(test, 0, len(test)-1)
    //InsertionSort(test)
    //BubbleSort(test)
    //HeapSort(test)
    //MergeSort(test, 0, len(test)-1)
    //res := CountingSort(test, 22)
    //SelectSort(test)
    //BucketSort(test, 10)
    //ShellSort(test, 2)
    RadixSort(test)
    fmt.Println(test)
}

// 1. 快速排序
func partition(src []int, p, r int) int {
    x := src[r]
    i := p - 1
    for j := p; j < r; j++ {
        if src[j] <= x {
            i++
            src[i], src[j] = src[j], src[i]
        }
    }
    src[i+1], src[r] = src[r], src[i+1]
    return i + 1
}

// QuickSort1 尾递归
func QuickSort1(src []int, p, r int) {
    for p < r {
        q := partition(src, p, r)
        QuickSort1(src, p, q-1)
        p = q + 1
    }
}

// QuickSort2 常规递归
func QuickSort2(src []int, p, r int) {
    if p < r {
        q := partition(src, p, r)
        QuickSort2(src, p, q-1)
        QuickSort2(src, q+1, r)
    }
}

// 2. 随机快排
func randInt(max, min int) int {
    if max == 0 || min == 0 || min >= max {
        return max
    }
    return rand.Intn(max-min) + min
}

func RandomPartition(src []int, p, r int) int {
    i := randInt(r, p)
    src[i], src[r] = src[r], src[i]
    return partition(src, p, r)
}

func RandomQuickSort(src []int, p, r int) {
    /*if p < r {
        q := RandomPartition(src, p, r)
        RandomQuickSort(src, p, q-1)
        RandomQuickSort(src, q+1, r)
    }*/
    for p < r {
        q := RandomPartition(src, p, r)
        RandomQuickSort(src, p, q-1)
        p = q + 1
    }
}

// InsertionSort 3. 插入排序
func InsertionSort(src []int) {
    if len(src) == 0 {
        return
    }
    for j := 1; j < len(src); j++ {
        key := src[j]
        i := j - 1
        for i >= 0 && src[i] > key {
            src[i+1] = src[i]
            i--
        }
        src[i+1] = key
    }
}

// BubbleSort 4. 冒泡排序
func BubbleSort(src []int) {
    if len(src) == 0 {
        return
    }
    for i := 0; i < len(src); i++ {
        for j := 0; j < len(src)-1-i; j++ {
            if src[j] > src[j+1] {
                src[j], src[j+1] = src[j+1], src[j]
            }
        }
    }
}

// 5. 堆排序
func heapify(src []int, n, i int) {
    max := i
    l := 2*i + 1
    r := 2*i + 2
    if l < n && src[l] > src[max] {
        max = l
    }
    if r < n && src[r] > src[max] {
        max = r
    }
    if max != i {
        src[i], src[max] = src[max], src[i]
        heapify(src, n, max)
    }
}

func HeapSort(src []int) {
    // 1. 建堆
    n := len(src)
    for i := n/2 - 1; i >= 0; i-- {
        heapify(src, n, i)
    }
    // 2. 缩减堆大小，并调整堆
    for i := n - 1; i > 0; i-- {
        src[0], src[i] = src[i], src[0]
        heapify(src, i, 0)
    }
}

// 6. 归并排序
func merge(src []int, p, q, r int) {
    n1 := q - p + 1
    n2 := r - q
    L := make([]int, n1)
    R := make([]int, n2)
    for i := 0; i < n1; i++ {
        L[i] = src[p+i]
    }
    for j := 0; j < n2; j++ {
        R[j] = src[q+1+j]
    }
    i, j := 0, 0
    for k := p; k <= r; k++ {
        if i == n1 {
            src[k] = R[j]
            j++
        } else if j == n2 {
            src[k] = L[i]
            i++
        } else if L[i] <= R[j] {
            src[k] = L[i]
            i++
        } else {
            src[k] = R[j]
            j++
        }
    }
}

func MergeSort(src []int, p, r int) {
    if p < r {
        q := (p + r) / 2
        MergeSort(src, p, q)
        MergeSort(src, q+1, r)
        merge(src, p, q, r)
    }
}

// CountingSort 7. 计数排序 k 表示数据中的最大值
func CountingSort(src []int, k int) []int {
    res := make([]int, len(src))
    tmp := make([]int, k+1)
    for j:=0;j<len(src);j++ {
        tmp[src[j]]++
    }
    for i:=1;i<k+1;i++ {
        tmp[i] = tmp[i]+ tmp[i-1]
    }
    for j:=len(src)-1;j>=0;j-- {
        res[tmp[src[j]]-1] = src[j]
        tmp[src[j]]--
    }
    return res
}

// SelectSort 8. 选择排序
func SelectSort(src []int) {
    l := len(src)
    if l == 0 {
        return
    }
    for i:=0;i<l;i++ {
        min := i
        for j:=l-1;j>i;j-- {
            if src[min] > src[j] {
                min = j
            }
        }
        src[min], src[i] = src[i], src[min]
    }
}

// BucketSort 8. 桶排序
func BucketSort(src []int, bucketSize int) {
    if len(src) == 0 {
        return
    }
    // 1. 找到最大值和最小值
    min := src[0]
    max := src[0]
    for i:=1;i<len(src);i++ {
        if src[i] < min {
            min = src[i]
        }
        if src[i] > max {
            max = src[i]
        }
    }
    // 2.桶切片初始化
    bucketCount := make([][]int, (max-min)/bucketSize+1)
    // 3. 数据入桶
    for i:=0;i<len(src);i++{
        bucketCount[(src[i]-min)/bucketSize] = append(bucketCount[(src[i]-min)/bucketSize], src[i])
    }
    // 4. 将每个桶排序，写回原数组
    key := 0
    for _, bucket := range bucketCount {
        if len(bucket) == 0 {
            continue
        }
        InsertionSort(bucket)
        for _, value := range bucket {
            src[key] = value
            key++
        }
    }
}

// 9. 希尔排序
func InsertionSortByStep(src []int, step int) {
    for j:=step;j<len(src);j+=step {
        key := src[j]
        i := j-step
        for i>=0&&src[i]>key {
            src[i+step] = src[i]
            i-=step
        }
        src[i+step] = key
    }
}

func ShellSort(src []int, gap int) {
    l := len(src)
    for step:=l/gap;step>=1;step/=gap {
        for i:=0;i<step;i++ {
            InsertionSortByStep(src, step)
        }
    }
}

// 10. 基数排序
func RadixSort(src []int) {
    if len(src) == 0 {
        return
    }
    // 找到最大值
    max := src[0]
    for i:=1;i<len(src);i++ {
        if src[i] > max {
            max = src[i]
        }
    }
    // 找到最大位数
    c := 1
    for max>=10 {
        max/=10
        c++
    }
    d := 1
    for i:=0;i<c;i++ {
        InsertionSortByDigit(src, d)
        d*=10
    }
}

func InsertionSortByDigit(src []int, d int) {
    if len(src) == 0 {
        return
    }
    for j:=1;j<len(src);j++ {
        key := src[j]
        i := j-1
        for i>=0&&src[i]/d%10>key/d%10 {
            src[i+1] = src[i]
            i--
        }
        src[i+1] = key
    }
}


