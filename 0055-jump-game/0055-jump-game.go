func canJump(nums []int) bool {
    reachable := 0
    for i:=0;i<len(nums);i++ {
        
        if i > reachable {
            return false
        }

        newReach := i+nums[i]

        if newReach > reachable {
            reachable = newReach
        }

        if reachable >= len(nums)-1 {
            return true
        }
    }

    return true
}