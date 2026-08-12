func missingNumber(nums []int) int {
    
    sum,n := 0,len(nums)

    for i:=0;i<len(nums);i++ {
        sum += nums[i]
    }

    return n * (n+1)/2 - sum
}