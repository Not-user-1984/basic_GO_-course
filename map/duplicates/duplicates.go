package duplicates



func RemoveDulicates(input []string) []string {
	output := make([]string, 0)
	inputSet := make(map[string]struct{},len(input))
	for _, v := range input{
		if _, ok := inputSet[v]; !ok {
			output = append(output, v)
		}
		inputSet[v]= struct{}{}
	}
	return output
}