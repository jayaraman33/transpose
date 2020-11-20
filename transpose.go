package transpose

func Transpose(input []string) []string {


	if len(input) == 0 {
		return []string{}
	}

	for i := 0; i < len(input)-1; i++ {
		for len(input[i]) < len(input[i+1]) {
			input[i] += " "
		}
	}

	output := make([]string, len(input[0]))
	for i := range output {
		for _, row := range input {
			if i < len(row) {
				output[i] += string(row[i])
			}
		}
	}
	return output
}
