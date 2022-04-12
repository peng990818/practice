package main

// HeapMaximum 最大优先级队列
// 返回src中具有最大关键字的元素
func HeapMaximum(src []int) int {
    if len(src) == 0 {
        return 0
    }
    return src[0]
}

// HeapExtractMax 去掉并返回src中具有最大关键字的元素
func HeapExtractMax(src []int, n int) int {
    if n == 0 {
        return 0
    }
    max := src[0]
    src[0] = src[n-1]
    MaxHeapify2(src, n-1, 1)
    return max
}

// HeapIncreaseKey 将其中一个元素的值增加到key
func HeapIncreaseKey(src []int, i, key int) {
    if key < src[i] {
        return
    }
    src[i] = key
    for i>0 && src[(i-1)/2] < src[i] {
        src[(i-1)/2], src[i] = src[i], src[(i-1)/2]
        i = (i-1)/2
    }
}
