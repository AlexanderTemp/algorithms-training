// @leet start
// Uso de Hash table y demás 
import "sort"

func maxKDistinct(nums []int, k int) []int {
	m := make(map[int]bool)

	for _, x := range(nums) {
		m[x] = true
	}

	keys := make([]int, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})


	lent := len(keys)

	if lent > k {
		lent = k
	}

	return keys[0:lent]
}
// @leet end


// Lo mejor no es necesario el hash map solo ordenar y rebotar hasta completar el array
func maxKDistinct(nums []int, k int) []int {
	sort.Ints(nums)
	tam := len(nums)

	res := make([]int, 1, k)

	res[0] = nums[tam - 1]
	for i := N - 2; i >= 0; i -- {
		if k > 1 && nums[i] != nums[i+1] {
			res = append(res, nums[i])
			k--
		}
	}

	return res
}
