func characterReplacement(s string, k int) int {
    maxLength := 0
    for i:=0; i<len(s); i++ {
        if i>0 && s[i] == s[i-1] {
            continue
        }
        char := s[i]
        replacements := k
        j := i + 1
        for replacements > 0 && j < len(s) {
            if s[j] != char {
                replacements--
            }
            j++
        }
        if j == len(s) && replacements > 0 {
                j += replacements
        }
        for j < len(s) && s[j] == char {
            j++
        }
        length := min(j - i, len(s))
        if length > maxLength {
            maxLength = length
        }

    }

    return maxLength
}