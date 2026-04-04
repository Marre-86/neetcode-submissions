type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return "⏏"
	}
	str := ""
	for i,v := range strs {
		if i == 0 {
			str = v
		} else {
			str = str + "⌘" + v
		}
	}
	return str
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "⏏" {
		return []string{}
	}
	decoded := make([]string, 0)
	prevI := 0
	for i,v := range encoded {
		if (v == '⌘') {
			decoded = append(decoded, encoded[prevI:i])
			prevI = i+3
		}
	}
	decoded = append(decoded, encoded[prevI:len(encoded)])
	return decoded 
}
