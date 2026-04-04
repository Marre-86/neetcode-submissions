func productExceptSelf(nums []int) []int {
    prefixes := make([]int, len(nums))
    suffixes := make([]int, len(nums))
    result := make([]int, len(nums))

    prefixes[0] = 1;
    for i:=1; i<(len(nums)); i++ {
        prefixes[i] = prefixes[i-1] * nums[i-1]
    }
    fmt.Println(prefixes)

    suffixes[len(nums)-1] = 1;

    for j:=len(nums)-2; j>=0; j-- {
        suffixes[j] = suffixes[j+1] * nums[j+1]

    }
    fmt.Println(suffixes)

    for i:=0; i<(len(nums)); i++ {
        result[i] = prefixes[i] * suffixes[i]
    }

    return result
}
