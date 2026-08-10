type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res := strings.Builder{}

	if len(strs) == 0 {
		return "%2D"
	}

	for i := range strs {
		c := strs[i]

		if len(c) > 0 {
			for j := range c {
				switch c[j] {
				case '+':
					res.WriteString("%2B")
				case '%':
					res.WriteString("%2A")
				default:
					res.WriteString(string(c[j]))
				}
			}
		} else {
			res.WriteString("%2C")
		}

		res.WriteString("+")
	}

	r := res.String()

	return r[:len(r)-1]
}

func (s *Solution) Decode(encoded string) []string {

	res := []string{}
	i := 0

	if encoded == "%2D" {
		return res
	}
	
	for i < len(encoded) {
		var str strings.Builder
		j := i
		c := ""

		for j < len(encoded) {
			c = string(encoded[j])

			if c == "+" {
				break
			}

			if c == "%" {
				char := encoded[j : j+3]
				j += 3

				if char == "%2C" {
					break
				}

				if char == "%2A" {
					str.WriteString("%")
				} else {
					str.WriteString("+")
				}

			} else {
				str.WriteString(c)
				j += 1
			}
		}

		if str.Len() == 0 {
			res = append(res, "")
		} else {
			res = append(res, str.String())
		}

		i = j + 1
	}

	return res
}
