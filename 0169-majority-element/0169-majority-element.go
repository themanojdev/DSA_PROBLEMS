func majorityElement(nums []int) int {
    
    count,result := 0,0

    for i:=0;i<len(nums);i++ {
        if count == 0 {
            result = nums[i]
        }
        if result == nums[i] {
            count++
        } else {
            count--
        }
    }
    return result
}