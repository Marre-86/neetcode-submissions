func isPalindrome(s string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	clean := re.ReplaceAllString(s, "")
	b := 0
	e := len(clean)-1
	for b < e {
		if !strings.EqualFold(string(clean[b]), string(clean[e])) {
			return false
		}
		b++
		e--
	}
	return true
}
