func maxArea(height []int) int {
    l, r := 0, len(height)-1
    area := 0

    for l<r {
        newArea := min(height[l],height[r])*(r-l)
        area = max(area, newArea)
        if(height[l]>height[r]){
            r--
        } else{
            l++
        }
    }

    return area
}

func min(i, j int)int{
    if(i>j){
        return j
    }
    return i
}

func max(i, j int)int{
    if(i>j){
        return i
    } 
    return j
}