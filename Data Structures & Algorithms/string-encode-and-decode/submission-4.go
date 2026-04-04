type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    before := ""
    after := ""
    for _,v := range strs {
        before += strconv.Itoa(len(v)) + ",";
        after += v
    }

    return before + "#" + after
}

func (s *Solution) Decode(encoded string) []string {
    result := make([]string,0)
    before, after, found := strings.Cut(encoded, "#")
    if found {
        lengths := strings.Split(before,",")
        for i:=0; i<len(lengths); i++ {
            startIndex,err := strconv.Atoi(lengths[i])
            if err == nil {
                result = append(result, after[:startIndex]) 
                after = after[startIndex:]
            }
        }
    }
    return result
}
