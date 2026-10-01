package main

// HasDuplicate reports whether arr contains at least one duplicated element.
// It runs in O(n) time and O(n) space on average, using a hash set.
func HasDuplicate(arr []int) bool {
	seen := make(map[int]struct{}, len(arr))
	for _, v := range arr {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}

func main() {
	println("HasDuplicate([]int{1, 2, 3})        =", HasDuplicate([]int{1, 2, 3}))
	println("HasDuplicate([]int{1, 2, 2, 3})    =", HasDuplicate([]int{1, 2, 2, 3}))
	println("HasDuplicate([]int{})              =", HasDuplicate([]int{}))
	println("HasDuplicate([]int{42})            =", HasDuplicate([]int{42}))
}
