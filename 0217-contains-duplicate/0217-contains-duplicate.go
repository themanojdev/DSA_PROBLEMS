func containsDuplicate(nums []int) bool {
    result := make(map[int]bool)

    for i:=0;i<len(nums);i++ {
        if result[nums[i]] {
            return true
        }
        result[nums[i]] = true
    }

    return false
}