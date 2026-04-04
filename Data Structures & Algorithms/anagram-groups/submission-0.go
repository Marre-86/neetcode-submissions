func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	for _,word := range strs {
		found := false
		for i := range result {
			if isAnagram(word,result[i][0]) {
				result[i] = append(result[i], word)
				found = true
				break;
			}
		}
		if !found {
			result = append(result, []string{word})
		}
	}
	return result
}

func isAnagram(s string, t string) bool {
	if (len(s) != len(t)) {
		return false
	}

	hashmap := make(map[byte]int)

	for i:=0;i<len(s);i++ {
		if _,found := hashmap[s[i]]; found {
			hashmap[s[i]] += 1
		} else {
			hashmap[s[i]] = 1
		}
	}

	for i:=0;i<len(t);i++ {
		if _,found := hashmap[t[i]]; found {
			hashmap[t[i]] -= 1
			if hashmap[t[i]] == 0 {
				delete(hashmap, t[i])
			}
		} else {
			return false
		}
	}

	return true
}
