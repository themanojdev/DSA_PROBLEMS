func findDisappearedNumbers(nums []int) []int {
    for i:=0;i<len(nums);i++ {
        idx := abs(nums[i]) -1

        if nums[idx] > 0 {
            nums[idx] = -nums[idx]
        }
    }

    result := []int{}
    for i:=0;i<len(nums);i++ {
        if nums[i]>0 {
            result = append(result,i+1)
        }
    }
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}