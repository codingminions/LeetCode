func lengthOfLongestSubstring(s string) int {
    hMap := make(map[byte]int, 0)
    l, lg  := 0,0

    for r:=0 ; r<len(s);r++ {
        hMap[s[r]]++
        for hMap[s[r]] > 1 {
            hMap[s[l]]--
            l++
        }
        lg = int(math.Max(float64(lg), float64(r-l+1)))
    }

    return lg
}