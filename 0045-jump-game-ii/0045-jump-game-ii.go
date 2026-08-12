func jump(nums []int) int {
    jump := 0
    far := 0
    currentEnd := 0

    for i := 0 ; i < len(nums) -1 ; i++ {

        if i+nums[i] > far {
            far = i + nums[i]
        }

        if i == currentEnd {
            jump++
            currentEnd = far
        }
    }
    return jump
}