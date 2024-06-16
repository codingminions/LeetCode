func lengthOfLongestSubstring(s string) int {
    hMap := make(map[byte]int, 0)
    l, lg  := 0,0

    for r:=0 ; r<len(s);r++ {
        if hMap[s[r]] > 0 {
            l = int(math.Max(float64(l), float64(hMap[s[r]])))
        }
        lg = int(math.Max(float64(lg), float64(r-l+1)))
        hMap[s[r]] = r+1
    }

    return lg
}