func missingElement(nums []int, k int) int {
    n := len(nums)
    missedInGap := 0

    for i:=1;i<n;i++{
        missedInGap = nums[i]-nums[i-1]-1

        if missedInGap >= k {
            return nums[i-1]+k
        } 

        k -= missedInGap
    }

    return nums[n-1] + k
}