package main

// GetMaxAndMin 获取最大值。最小值
func GetMaxAndMin(src []int) (int, int) {
    if len(src) == 0 {
        return 0, 0
    }
    min, max, num := src[0], src[0], len(src)
    for i:=1;i<num;i+=2 {
        if i == num-1 {
            if src[i] > max {
                max = src[num-1]
            }
            if src[i] < min {
                min = src[num-1]
            }
            break
        }
        tmpMin, tmpMax := 0, 0
        if src[i] > src[i+1] {
            tmpMin = src[i+1]
            tmpMax = src[i]
        } else {
            tmpMin = src[i]
            tmpMax = src[i+1]
        }
        if tmpMin < min {
            min = tmpMin
        }
        if tmpMax > max {
            max = tmpMax
        }
    }
    return max, min
}

// RandomSelect 随机选择第i小的数
func RandomSelect(src []int, p, r, i int) int {
    if p==r {
        return src[p]
    }
    q := randomPartition(src, p, r)
    k := q-p+1
    if i==k{
        return src[q]
    } else if i<k {
        return RandomSelect(src, p, q-1, i)
    } else {
        return RandomSelect(src, q+1, r, i-k)
    }
}
