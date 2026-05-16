type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	var result strings.Builder

	for _, str := range strs {
		fmt.Fprintf(&result, "%d#%s", len(str), str)
	}

	return result.String()
}

func (s *Solution) Decode(encoded string) []string {

	fmt.Printf("encoded is %s\n", encoded)

	result := []string{}

	j := 0

	for j < len(encoded) {
		if encoded[j] == '1' || encoded[j] == '2' || encoded[j] == '3' || encoded[j] == '4' || encoded[j] == '5' || encoded[j] == '6' || encoded[j] == '7' || encoded[j] == '8' || encoded[j] == '9' || encoded[j] == '0' {

			number := ""

			for encoded[j] != '#' {
				number += string(encoded[j])
				j += 1
			}

			num, err := strconv.Atoi(number)

			if err != nil {
				fmt.Println("Error during conversion:", err)
			}

			result = append(result, encoded[j+1:j+num+1])

			j = j + num + 1
		}
	}

	return result
}

