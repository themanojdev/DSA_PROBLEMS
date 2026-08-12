func trap(height []int) int {

    start,water,end := 0,0,len(height)-1
    leftMax,rightMax := 0,0

    for start < end {
        if height[start] < height[end] {
            if height[start] < leftMax {
                water += leftMax - height[start]
            } else {
                leftMax = height[start]
            }
            start++
        } else {

            if height[end] < rightMax {
                water += rightMax - height[end]
            } else {
                rightMax = height[end]
            }

            end--
        }
    }

    return water
}