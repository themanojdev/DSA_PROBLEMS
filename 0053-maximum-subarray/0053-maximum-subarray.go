func maxSubArray(nums []int) int {
    
    maxSum := nums[0]

    sum := nums[0]

    if len(nums) == 1 {
        maxSum = nums[0]
        return maxSum
    }

    for i:=1;i<len(nums);i++ {
        if sum >= 0 {
            sum += nums[i]
        } else if sum < 0 {
            sum = nums[i]
        }

        if sum > maxSum {
            maxSum = sum
        }
    }

    return maxSum
}