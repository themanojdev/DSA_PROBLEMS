func twoSum(nums []int, target int) []int {
    
    if len(nums) == 0 {
        return []int{}
    }

    resultMap := make(map[int]int)
    for i:=0 ; i<len(nums);i++ {

        b := target-nums[i]


        if value,ok := resultMap[b] ; ok {
            return []int{i,value}
        }

        resultMap[nums[i]]=i
        
    }

    return []int{}
}