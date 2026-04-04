type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    begin := ""
    end := ""
    for _,v := range strs {
        begin += strconv.Itoa(len(v)) + ","
        end += v
    }
    return begin + "#" + end
}

func (s *Solution) Decode(encoded string) []string {
    var result []string
    fmt.Println(encoded)
    begin, end, found := strings.Cut(encoded, "#")
    if found {
        for _,length := range strings.Split(begin,",") {
            if length, err := strconv.Atoi(length); err == nil {
                word := end[:length]
                end = end[length:]
                result = append(result, word)
            }
        }
    }
    return result
}
