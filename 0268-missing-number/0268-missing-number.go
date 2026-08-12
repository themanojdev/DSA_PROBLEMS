func missingNumber(nums []int) int {
    
    xor := len(nums)

    for i:=0;i<len(nums);i++ {
        xor = xor ^ i ^ nums[i]
    }
    return xor
}