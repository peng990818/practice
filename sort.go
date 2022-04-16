package main

import (
    "math/rand"
    "time"
)

// InsertionSort 插入排序(升序）
func InsertionSort(src []int) []int {
    if len(src) == 0 || len(src) == 1 {
        return src
    }
    for j := 1; j < len(src); j++ {
        // 1. 标记数组有序部分的最后一张
        key := src[j]
        // 2. 向前遍历有序部分
        i := j - 1
        // 将比第一个无序数大的向后移动
        for i >= 0 && src[i] > key {
            src[i+1] = src[i]
            i = i - 1
        }
        // 如果数组遍历到头或者找到第一个小于等于他的数则插入到这个数后面
        src[i+1] = key
    }
    return src
}

// InsertionSortDes 插入排序(降序）
func InsertionSortDes(src []int) []int {
    if len(src) == 0 || len(src) == 1 {
        return src
    }
    for j := 1; j < len(src); j++ {
        // 1. 标记数组有序部分的最后一张
        key := src[j]
        // 2. 向前遍历有序部分
        i := j - 1
        // 将比第一个无序数小的向后移动
        for i >= 0 && src[i] < key {
            src[i+1] = src[i]
            i = i - 1
        }
        // 如果数组遍历到头或者找到第一个小于等于他的数则插入到这个数后面
        src[i+1] = key
    }
    return src
}

// MergeSort 归并排序
func MergeSort(src []int, p, r int) []int {
    if p < r {
        q := (p + r) / 2
        // 左侧递归
        MergeSort(src, p, q)
        // 右侧递归
        MergeSort(src, q+1, r)
        // 合并
        Merge2(src, p, q, r)
    }
    return src
}

// Merge1 使用哨兵的merge
func Merge1(src []int, p, q, r int) {
    // 第一个数组的长度
    n1 := q - p + 1
    // 第二个数组的长度
    n2 := r - q
    // 创建两个数组，长度分别比原有的长度多一个，最后一个元素为哨兵牌
    L := make([]int, n1+1)
    R := make([]int, n2+1)
    // 为每一个数组赋值
    for i := 0; i < n1; i++ {
        L[i] = src[p+i]
    }
    for j := 0; j < n2; j++ {
        R[j] = src[q+1+j]
    }
    // 设置哨兵
    L[n1] = 1000000
    R[n2] = 1000000
    i, j := 0, 0
    // 重排数组
    for k := p; k <= r; k++ {
        if L[i] <= R[j] {
            src[k] = L[i]
            i++
        } else {
            src[k] = R[j]
            j++
        }
    }
}

func Merge2(src []int, p, q, r int) {
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

// 逆序对
func ReversePairs(src []int) int {
    return mergeSort(src, 0, len(src)-1)
}

func mergeSort(src []int, start, end int) int {
    if start >= end {
        return 0
    }
    // 1. 找到中间位置
    mid := start + (end-start)/2
    // 2. 递归左右两边
    cnt := mergeSort(src, start, mid) + mergeSort(src, mid+1, end)
    // 3. 申请临时数组
    tmp := []int{}
    i, j := start, mid+1
    for i <= mid && j <= end {
        if src[i] <= src[j] {
            tmp = append(tmp, src[i])
            cnt += j - mid - 1 // 右数组当前元素的左元素全部小于左数组当前元素，因此全为逆序对
            i++
        } else {
            tmp = append(tmp, src[j])
            j++
        }
    }
    for ; i <= mid; i++ {
        tmp = append(tmp, src[i])
        cnt += end - mid // 同上所述
    }
    for ; j <= end; j++ {
        tmp = append(tmp, src[j])
    }
    for i := start; i <= end; i++ {
        src[i] = tmp[i-start]
    }
    return cnt
}

// BubbleSort 冒泡排序
func BubbleSort(src []int) []int {
    if len(src) == 0 || len(src) == 1 {
        return src
    }
    /*for i:=0;i<len(src);i++{
        for j:=len(src)-1;j>=i+1;j--{
            if src[j] < src[j-1] {
                src[j] = src[j]^src[j-1]
                src[j-1] = src[j]^src[j-1]
                src[j] = src[j]^src[j-1]
            }
        }
    }*/
    for i := 0; i < len(src); i++ {
        for j := 0; j < len(src)-1-i; j++ {
            if src[j] > src[j+1] {
                src[j] = src[j] ^ src[j+1]
                src[j+1] = src[j] ^ src[j+1]
                src[j] = src[j] ^ src[j+1]
            }
        }
    }
    return src
}

// MaxHeapify1 非递归大根堆调整
func MaxHeapify1(src []int, i int, n int) []int {
    largest := 0
    for i < len(src) {
        l := 2*i + 1
        r := 2*i + 2
        if l < n && src[l] > src[i] {
            largest = l
        } else {
            largest = i
        }
        if r < n && src[r] > src[largest] {
            largest = r
        }
        if largest != i {
            src[i] = src[i] ^ src[largest]
            src[largest] = src[i] ^ src[largest]
            src[i] = src[i] ^ src[largest]
            i = largest
        } else {
            break
        }
    }
    return src
}

// 堆排序

// BuildMaxHeap 构建大根堆
func BuildMaxHeap(src []int) {
    // 从最后一个父节点建堆，并不断调整
    l := len(src)
    for i := l/2 - 1; i >= 0; i-- {
        MaxHeapify2(src, i, l)
    }
}

func MinHeapify1(src []int, n, i int) []int {
    // 1. 标记最小下标
    smallest := 0
    for i < len(src) {
        // 2. 左孩子
        l := 2*i + 1
        // 3. 右孩子
        r := 2*i + 2
        // 4. 通过判断找出最小下标
        if l < n && src[l] < src[i] {
            smallest = l
        } else {
            smallest = i
        }
        if r < n && src[r] < src[smallest] {
            smallest = r
        }
        // 5. 如果不是父节点，则交换，并将子节点变成新的父节点，继续调整
        // 否则跳出循环
        if smallest != i {
            src[i] = src[i] ^ src[smallest]
            src[smallest] = src[i] ^ src[smallest]
            src[i] = src[i] ^ src[smallest]
            i = smallest
        } else {
            break
        }
    }
    return src
}

func MinHeapify2(src []int, n, i int) []int {
    smallest := 0
    l := 2*i + 1
    r := 2*i + 2
    if l < n && src[l] < src[i] {
        smallest = l
    } else {
        smallest = i
    }
    if r < n && src[r] < src[smallest] {
        smallest = r
    }
    if smallest != i {
        src[i] = src[i] ^ src[smallest]
        src[smallest] = src[i] ^ src[smallest]
        src[i] = src[i] ^ src[smallest]
        MinHeapify2(src, n, smallest)
    }
    return src
}

func BuildMinHeap(src []int) []int {
    l := len(src)
    for i := l/2 - 1; i >= 0; i-- {
        MinHeapify1(src, l, i)
    }
    return src
}

// MaxHeapify2 递归大根堆调整
func MaxHeapify2(src []int, n, i int) {
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
        src[max], src[i] = src[i], src[max]
        MaxHeapify2(src, n, max)
    }
}

//arr 存储堆的数组
//n 数组长度
//i 待维护元素的下标
func heapify(arr []int, n int, i int) {
    max := i
    //左孩子节点
    lson := 2*i + 1
    //右孩子节点
    rson := 2*i + 2
    if lson < n && arr[lson] > arr[max] {
        max = lson
    }
    if rson < n && arr[rson] > arr[max] {
        max = rson
    }
    //把父节点换下去并向下调整
    if max != i {
        arr[max], arr[i] = arr[i], arr[max]
        heapify(arr, n, max)
    }
}

func HeapSort(arr []int) []int {
    //建堆
    n := len(arr)
    for i := n/2 - 1; i >= 0; i-- {
        //heapify(arr, n, i)
        MinHeapify1(arr, n, i)
    }
    //排序
    for i := n - 1; i > 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        //heapify(arr, i, 0)
        MinHeapify1(arr, i, 0)
    }
    return arr
}

// 分割
func partition1(src []int, p, r int) int {
    // 以最后一位作为分割点
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

// QuickSort1 快速排序
func QuickSort1(src []int, p, r int) []int {
    if p < r {
        q := partition1(src, p, r)
        QuickSort1(src, p, q-1)
        QuickSort1(src, q+1, r)
    }
    return src
}

func partition2(src []int, p, r int) int {
    x := src[r]
    i := p - 1
    for j := p; j < r; j++ {
        if src[j] < x {
            i++
            src[i], src[j] = src[j], src[i]
        }
        if src[j] == x {
            i++
        }
    }
    if i == r-1 {
        return (p + r) / 2
    }
    src[i+1], src[r] = src[r], src[i+1]
    return i + 1
}

func RandomQuickSort(src []int, p, r int) []int {
    if p < r {
        q := randomPartition(src, p, r)
        RandomQuickSort(src, p, q-1)
        RandomQuickSort(src, q+1, r)
    }
    return src
}

func randomPartition(src []int, p, r int) int {
    i := RandInt(r, p)
    src[i], src[r] = src[r], src[i]
    return partition1(src, p, r)
}

func RandInt(max, min int) int {
    if min >= max || min == 0 || max == 0 {
        return max
    }
    return rand.Intn(max-min) + min
}

// QuickSort2 尾递归快速排序
func QuickSort2(src []int, p, r int) []int {
    for p < r {
        q := partition1(src, p , r)
        QuickSort2(src, p, q-1)
        p = q+1
    }
    return src
}

// CountingSort1 稳定 计数排序
func CountingSort1(src []int, k int) []int {
    res := make([]int, len(src))
    tmp := make([]int, k+1)
    for i:=0;i<k+1;i++ {
        tmp[i]=0
    }
    for j:=0;j<len(src);j++ {
        tmp[src[j]]++
    }
    for i:=1;i<k+1;i++ {
        tmp[i] = tmp[i] + tmp[i-1]
    }
    //fmt.Println(tmp)
    for j:=len(src)-1;j>=0;j-- {
        res[tmp[src[j]]-1] = src[j]
        tmp[src[j]]--
    }
    return res
}

// CountingSort2 不稳定 计数排序
func CountingSort2(src []int, k int) []int {
    res := make([]int, 0, len(src))
    tmp := make([]int, k+1)
    for i:=0;i<k+1;i++ {
        tmp[i]=0
    }
    for j:=0;j<len(src);j++ {
        tmp[src[j]]++
    }
    for i:=0;i<k+1;i++ {
        if tmp[i] == 0{
            continue
        }
        for j:=0;j<tmp[i];j++ {
            res = append(res, i)
        }
    }
    return res
}

// BucketSort 桶排序
func BucketSort(src []int, bucketSize int) []int {
    // 1. 找到最大值和最小值
    minValue := src[0]
    maxValue := src[0]

    for i := 0; i<len(src); i++ {
        if minValue > src[i] {
            minValue = src[i]
        }
        if maxValue < src[i] {
            maxValue = src[i]
        }
    }

    // 2. 桶切片初始化
    // 桶数 = 最大值-最小值 / 桶的大小 + 1
    bucketCount := make([][]int,(maxValue - minValue) / bucketSize +1)

    // 3. 数据入桶
    // 找到对应的桶的索引，将数据写入
    for i := 0; i<len(src); i++ {
        bucketCount[(src[i] - minValue) / bucketSize] = append(bucketCount[(src[i] - minValue) / bucketSize] , src[i])
    }
    // 4. 将非零的桶中的数据进行排序，写回原有的数组
    key := 0
    for _, bucket := range bucketCount {
        if len(bucket) == 0 {
            continue
        }
        bucket = InsertionSort(bucket)
        for _, value := range bucket {
            src[key] = value
            key++
        }
    }
    return src
}

func InsertionByDigit(src []int, d int) {
    if len(src) == 0 {
        return
    }
    for i:=1;i<len(src);i++ {
        key := src[i]
        j := i-1
        for j >=0 && src[j]/d > key/d {
            src[j+1] = src[j]
            j--
        }
        src[j+1] = key
    }
}

// RadixSort 基数排序
func RadixSort(src []int) []int {
    if len(src) == 0 {
        return nil
    }
    // 1. 找到最大值
    max := src[0]
    for i:=1;i<len(src);i++ {
        if src[i]>max {
            max = src[i]
        }
    }
    // 2. 找到最大位数
    d := 1
    for max >= 10 {
        max /= 10
        d*=10
    }

    for c:=1;c!=d;c*=10{
        InsertionByDigit(src, c)
    }
    return src
}

// SelectSort 选择排序
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
        src[i], src[min] = src[min], src[i]
    }
}

// ShellSort 希尔排序
func ShellSort(src []int, gap int) {
    l := len(src)
    for step:=l/gap; step>=1; step/=gap {
        for i:=0;i<step;i++ {
            InsertionSortByStep(src, step)
        }
    }
}

func InsertionSortByStep(src []int, step int) {
    for j:=step;j<len(src);j+=step {
        key := src[j]
        i := j-step
        for i>=0 && src[i] > key {
            src[i+step] = src[i]
            i-=step
        }
        src[i+step] = key
    }
}

// RandomIntSlice 随机一个整数切片
func RandomIntSlice(length int, limit int) []int {
    if length==0 || limit == 0 {
        return nil
    }
    res := make([]int, 0, length)
    rand.Seed(time.Now().Unix())
    for i:=0;i<length;i++ {
        res = append(res, rand.Intn(limit))
    }
    return res
}


