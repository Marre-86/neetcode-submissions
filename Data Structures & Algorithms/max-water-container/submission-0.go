func maxArea(heights []int) int {
	maxArea := 0
	for i,_ := range heights {
		for e:= i+1; e<len(heights); e++ {
			area := (e - i) * min(heights[i], heights[e])
			if maxArea < area {
				maxArea = area
			}
		}
	}

	return maxArea
}
