func convertToTitle(columnNumber int) string {
	columnName := ""

	for columnNumber != 0 {
		columnNumber -= 1
		columnName = getSingleLetter(columnNumber % 26) + columnName
		columnNumber = columnNumber / 26
	}

	return columnName
}

func getSingleLetter(n int) string {
    if n < 0 || n > 25 {
        return ""
    }
    return string(rune(65 + n))
}