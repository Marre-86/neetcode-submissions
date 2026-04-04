func hasDuplicate(nums []int) bool {
    hashmap := make(map[int]bool)
	for _,v := range nums {
		if _,exists := hashmap[v]; exists {
			return true
		}
		hashmap[v] = true
	}
	return false
}
