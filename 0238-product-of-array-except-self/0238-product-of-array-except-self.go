func productExceptSelf(nums []int) []int {
    
    resultArr := make([]int,len(nums))

    prefix := 1
    for i:=0;i<len(nums);i++ {
        resultArr[i] = prefix
        prefix *= nums[i]
    }
    suffix := 1
    for i:=len(nums)-1;i>=0;i-- {
       resultArr[i] *= suffix
       suffix *= nums[i]
    }

    return resultArr

}