func hasDuplicate(nums []int) bool {
    numMap := make(map[int]bool)

    for _, val := range(nums) {
        if _, exists := numMap[val]; exists {
            return true
        }

        numMap[val] = true
    }

    return false
}
