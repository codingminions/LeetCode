func isAlienSorted(words []string, order string) bool {
    hMap := make(map[rune]int,0)

    for i,str := range order {
        hMap[str] = i+1
    }

    checkLexicography := func(word1, word2 string)bool{
        parseLength := len(word1)

        if(len(word1)>len(word2)){
            parseLength = len(word2)
        } 

        for i:=0;i<parseLength;i++{
            if(hMap[rune(word1[i])]>hMap[rune(word2[i])]){
                return false
            } else if (hMap[rune(word1[i])]<hMap[rune(word2[i])]){
                return true
            }
        }

        if(len(word1)>len(word2)){
            return false
        }

        return true
    }

    for i:=0;i<len(words)-1;i++ {
        if checkLexicography(words[i], words[i+1])==false{
            return false
        }
    }

    return true
}