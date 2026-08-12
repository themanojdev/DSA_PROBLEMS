func majorityElement(nums []int) int {
    sort.Ints(nums)

    count,result := 1,nums[0]

    if len(nums) == 1 {
        return result
    }

    for i:=1;i<len(nums);i++ {
        if nums[i] == nums[i-1] {
            count++
        } else {
            count = 1
            result = nums[i]
        }

        if count > len(nums)/2 {
            return result
        }
    }

    return -1
}