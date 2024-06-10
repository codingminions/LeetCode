func intersection(nums1 []int, nums2 []int) []int {
    hMap := make([]int,1000)
    result := []int{}

    for _,ele := range nums1{
        hMap[ele]++
    }

    for _, ele := range nums2{
        if(hMap[ele]>0){
            result = append(result, ele)
            hMap[ele]=0
        }
    }

    return result
}