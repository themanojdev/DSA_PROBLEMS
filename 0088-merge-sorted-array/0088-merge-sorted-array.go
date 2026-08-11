func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    result := make([]int,m+n)

    i,j,start := 0,0,0

    if n == 0 {
        return
    }

    for i < m && j < n {
        if nums1[i] < nums2[j] {
            result[start] = nums1[i]
            i++
        } else {
            result[start] = nums2[j]
            j++
        }
        start++  
    }

    for i < m {
        result[start] = nums1[i]
        i++
        start++
    }

    for j < n {
        result[start] = nums2[j]
        j++
        start++
    }

    

    copy(nums1,result)
}