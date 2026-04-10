func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	fmt.Println(nums)

	for i:=0; i < len(nums)-1; i++ {
		if i>0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]>0 {
			break
		}

		l, r := i + 1, len(nums) - 1

		for l < r {
			threeSum := nums[i] + nums[l] + nums[r]
			if threeSum < 0 {
				l++
			} else if threeSum > 0 {
				r--
			} else if threeSum == 0 {
				result = append(result, []int {nums[i], nums[l], nums[r]})
				l++
				r--
				for l<r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return result
}
