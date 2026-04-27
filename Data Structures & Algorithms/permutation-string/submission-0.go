func checkInclusion(s1 string, s2 string) bool {
    var needle [26]int
    for i:=0; i<len(s1); i++ {
        needle[s1[i] - 'a'] += 1
    }

    for l:=0;l<len(s2); l++ {
        r  := l + len(s1)
        if r > len(s2) {
            break
        }

        var haystack [26]int
        for i:=l; i< r; i++ {
            haystack[s2[i] - 'a'] += 1
        }

        if needle == haystack {
            return true
        }

    }
    return false
}
