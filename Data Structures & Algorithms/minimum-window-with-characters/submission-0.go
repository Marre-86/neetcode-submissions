func minWindow(s string, t string) string {
    if t == "" {
        return ""
    }

    shortest := ""
    shortestLength := math.MaxInt64

    needle := make (map[rune]int)
    window := make (map[rune]int)

    for _,letter := range t {
        needle[letter]++
    }

    need, have := len(needle), 0
    l := 0

    for r:=0; r < len(s); r++ {
        if _,found := needle[rune(s[r])]; found {
            window[rune(s[r])]++
            if window[rune(s[r])] == needle[rune(s[r])] {
                have++
            }
        }

        for need == have {
            if len(s[l:r+1]) < shortestLength {
                shortestLength = len(s[l:r+1])
                shortest = s[l:r+1]
            }
            if _,found := needle[rune(s[l])]; found {
                if window[rune(s[l])] == needle[rune(s[l])] {
                    have--
                }
                window[rune(s[l])]--
            }
            l++
        }

    }

    return shortest
}