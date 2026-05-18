func reverseString(s []byte) {
	
	n := len(s)

	for i := range n / 2 {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}

}
