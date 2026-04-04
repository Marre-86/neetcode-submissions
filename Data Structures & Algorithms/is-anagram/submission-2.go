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
