func isValid(s string) bool {
	
	if len(s)%2 != 0 {
		return false
	}

    stack := []byte{}

	c := func(b byte) bool {
		return b == '(' || b == '[' || b == '{'
	}

	c2 := func(a, b byte) bool {

		if a == '(' {
			return b == ')'
		}

		if a == '[' {
			return b == ']'
		}

		return b == '}'
	}

	for i := range len(s) {
		if c(s[i]) {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if !c2(stack[len(stack)-1], s[i]) {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
