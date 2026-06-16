import (
	"slices"
)


func numRescueBoats(people []int, limit int) int {
	n := len(people)

	slices.Sort(people)

	numBoats := 0

	left := 0

	right := n - 1

	for left <= right {

		t := limit - people[left]

		for people[right] > t && left < right {
			right--
			numBoats += 1
		}

		numBoats += 1
		left += 1
		right -= 1

	}

	return numBoats
}
