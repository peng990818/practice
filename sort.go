package main

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

