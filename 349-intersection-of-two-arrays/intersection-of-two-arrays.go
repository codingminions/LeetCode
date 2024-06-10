func intersection(nums1 []int, nums2 []int) []int {
    hMap := make(map[int]bool,0)
    result := []int{}

    for _,ele := range nums1{
        hMap[ele] = true
    }

    for _, ele := range nums2{
        if(hMap[ele]==true){
            result = append(result, ele)
            hMap[ele]=false
        }
    }

    return result
}