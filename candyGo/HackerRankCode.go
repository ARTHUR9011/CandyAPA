func candies(n int32, arr []int32) int64 {
    if n <= 1 {
        return int64(n)
    }

    candies := make([]int64, n)
    candies[0] = 1

    for i := int32(1); i < n; i++ {
        if arr[i] > arr[i-1] {
            candies[i] = candies[i-1] + 1
        } else {
            candies[i] = 1
        }
    }

    total := candies[n-1]

    for i := n - 2; i >= 0; i-- {
        if arr[i] > arr[i+1] {
            candies[i] = max64(candies[i], candies[i+1]+1)
        }
        total += candies[i]
    }

    return total
}

func max64(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}
