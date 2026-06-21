import "slices"

func findClosestElements(arr []int, k int, x int) []int {
	res := make([]int, k)

	if x < arr[0] {
		for i := 0; i < k; i++ {
			res[i] = arr[i]
		}

		return res
	}

	if x > arr[len(arr)-1] {
		for i := len(arr) - k; i < len(arr); i++ {
			res[i-(len(arr)-k)] = arr[i]
		}

		return res
	}

	l := 0
	r := 0

	for i := range arr {
		if i == len(arr)-1 {
			r = i
			l = i - 1
			break
		}

		if arr[i] < x && arr[i+1] >= x {
			r = i + 1
			l = i
			break
		}
	}

	for k > 0 {
		if r == len(arr) && l >= 0 {
			res[k-1] = arr[l]
			l -= 1
			k -= 1
			continue
		}

		if l < 0 && r < len(arr) {
			res[k-1] = arr[r]
			r += 1
			k -= 1
			continue
		}

		if absolute(arr[l]-x) <= absolute(arr[r]-x) {
			res[k-1] = arr[l]
			l -= 1
			k -= 1
		} else {
			res[k-1] = arr[r]
			r += 1
			k -= 1
		}
	}

	slices.Sort(res)

	return res
}

func absolute(a int) int {
	if a >= 0 {
		return a
	}

	return a * -1
}

