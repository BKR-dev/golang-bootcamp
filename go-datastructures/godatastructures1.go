package godatastructures

// looks up a slice of integers for duplicate entries
// in a new map, if an entry is not found, add it to the map
// because that happens on the range through the slice
// its very fast but not the best for memory allocation
func containsDuplicate(nums []int) bool {
	var m = make(map[int]bool)
	for _, v := range nums {
		_, k := m[v]
		if k {
			return true
		} else {
			m[v] = true
		}
	}
	return false
}
