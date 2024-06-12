func isMonotonic(nums []int) bool {
    monotonicIncreasing := true
    monotonicDecreasing := true

    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            monotonicDecreasing = false
        } else if nums[i] < nums[i-1] {
            monotonicIncreasing = false
        }
        if !monotonicIncreasing && !monotonicDecreasing {
            return false
        }
    }
    return true 
}