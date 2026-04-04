func twoSum(nums []int, target int) []int {
    hashmap := map[int]int {}
    for i,v := range nums {
        complement := target - v
        if _,ok := hashmap[complement]; ok {
            if (i < hashmap[complement]) {
                return []int {i,hashmap[complement]}
            } else {
                return []int {hashmap[complement],i}
            }
        }
        hashmap[v] = i
    }

    return nil
}
