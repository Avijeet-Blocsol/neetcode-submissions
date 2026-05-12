func twoSum(nums []int, target int) []int {

    numMap := make(map[int]int)

    for i, val := range(nums) {
        numMap[val] = i
    }

    for j, val := range(nums) {
        if l, exists := numMap[target-val]; exists && j != l {
            if j < l {
                return []int{j,l}
            } else {
                return []int{l,j}
            }
        }
    }

    return []int{}    
}
