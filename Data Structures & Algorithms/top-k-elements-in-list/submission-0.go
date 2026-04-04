func topKFrequent(nums []int, k int) []int {
    hmap := make(map[int]int)

    for _,v := range nums {
        hmap[v]++
    }

    type kv struct {
        Key int
        Value int
    }

    var pairs []kv
    for k,v := range hmap {
        pairs =  append(pairs, kv{Key:k,Value:v})
    }

    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].Value > pairs[j].Value
    })

    result := make([]int, 0)

    for i:=0;i<k;i++ {
        result = append(result, pairs[i].Key)
    }

    return result
}