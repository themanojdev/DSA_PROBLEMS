func rotate(nums []int, k int)  {
    start,end := 0 , len(nums)-1

    k = k%len(nums)

    reverse(start,end,nums)
    reverse(start,k-1,nums)
    reverse(k,end,nums)
}

func reverse(start,end int,arr []int) {
    for start < end {
        arr[start],arr[end] = arr[end],arr[start]
        start++
        end--
    }
    
}