func lengthOfLongestSubstring(s string) int {
    hashmap := make(map[byte]int)
    maxLength := 0
    activeSegmentStartIndex := 0

    for i:=0; i<len(s); i++ {
        if i - activeSegmentStartIndex > maxLength {
            maxLength = i - activeSegmentStartIndex
            
        }
        if prevI,ok := hashmap[s[i]]; ok {
            activeSegmentStartIndex = max(activeSegmentStartIndex, prevI + 1)
        }
        hashmap[s[i]] = i
    }

    if len(s) - activeSegmentStartIndex > maxLength {
        maxLength = len(s) - activeSegmentStartIndex
        
    }

    return maxLength

}