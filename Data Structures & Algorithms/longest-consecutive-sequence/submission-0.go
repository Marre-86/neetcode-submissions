func longestConsecutive(nums []int) int {
	hashmap := make(map[int]struct{})

	for _,v := range nums {
		hashmap[v] = struct{}{}
	}

	max := 0

	for num,_ := range hashmap {
		cur := 1
		if _,ok := hashmap[num-1]; !ok {
			fmt.Println(num)
			for {
				if cur > max {
					max = cur
				}
				if _,ok := hashmap[num+cur]; !ok {
					break
				}
				cur++
			}
		}
	}


	return max
}
