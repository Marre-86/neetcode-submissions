func trap(height []int) int {
	l, r := 0, len(height) - 1
	maxL, maxR := 0, 0

	water := 0
	for l < r {
		if height[l] > maxL {
			maxL = height[l]
		}
		if height[r] > maxR {
			maxR = height[r]
		}
        surface := min(maxL, maxR)
		if maxL < maxR {
            water += max(0, surface - height[l])
			l++
		} else {
            water += max(0, surface - height[r])
			r--
		}
	}

	return water
}