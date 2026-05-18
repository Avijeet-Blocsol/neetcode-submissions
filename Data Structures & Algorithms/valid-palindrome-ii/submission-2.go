func validPalindrome(s string) bool {
	
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			if checkValidPalindrome(s, left+1, right+1) {
				return true
			}

			return checkValidPalindrome(s, left, right)
		} else {
			left++
			right--
		}
	}

	return true
}

func checkValidPalindrome(s string, l, r int) bool {

	if !(l < r) {
		return false
	}

	val := s[l:r]

	n := len(val)

	for i := range n {
		if val[i] != val[n-1-i] {
			return false
		}
	}

	return true
}

