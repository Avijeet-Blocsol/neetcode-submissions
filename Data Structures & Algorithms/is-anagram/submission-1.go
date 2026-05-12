func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    charMap := map[rune]int{
        rune('a') : 0,
        rune('b') : 1,
        rune('c') : 2,
        rune('d') : 3,
        rune('e') : 4,
        rune('f') : 5,
        rune('g') : 6,
        rune('h') : 7,
        rune('i') : 8,
        rune('j') : 9,
        rune('k') : 10,
        rune('l') : 11,
        rune('m') : 12,
        rune('n') : 13,
        rune('o') : 14,
        rune('p') : 15,
        rune('q') : 16,
        rune('r') : 17,
        rune('s') : 18,
        rune('t') : 19,
        rune('u') : 20,
        rune('v') : 21,
        rune('w') : 22,
        rune('x') : 23,
        rune('y') : 24,
        rune('z') : 25,
    }

    strArr := make([]int, 26)

    for _, val := range(s) {
        strArr[charMap[val]]++
    }

    for _, val := range(t) {
        strArr[charMap[val]]--
    }

    for _, val := range(strArr) {
        if val != 0 {
            return false
        }
    }

    return true
}
