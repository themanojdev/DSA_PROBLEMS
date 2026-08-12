func sortColors(nums []int)  {

    low,mid,high := 0,0,len(nums)-1

    for mid <= high {

        if nums[mid] == 0 {
            swap(nums,low,mid)
            low++
            mid++
        } else if nums[mid] == 1 {
            mid++
        } else {
            swap(nums,mid,high)
            high--
        }
    }
    
}

func swap(num []int,i,j int) {
    num[i],num[j] = num[j],num[i]
}