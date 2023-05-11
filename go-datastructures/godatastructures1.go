package godatastructures

import (
	"sort"
	"strings"
)

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

// checks two strins for being anagrams like "toto" and "otto"
// or "lactoprotein" and "protectional"
// adds a rune to the map and increases its value per times rune found
// decreases the value of the rune when found in the other string
// if a rune is not found in the other string, its done
func isAnagramButGood(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var m = make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}
	return true
}

// quitting when the inputs cannot be an anagram
// sorting the strings and comparing them
// most slow and wasteful code possible, its literally javascript
func isAnagramButItsJavaScript(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ss := strings.Split(s, "")
	sort.Strings(ss)
	ts := strings.Split(t, "")
	sort.Strings(ts)
	s = strings.Join(ss, "")
	t = strings.Join(ts, "")
	if t == s {
		return true
	} else {
		return false
	}
}

// [1,2,3,4] - target 6 - [1,3]
// [3,3] - target 6 - [0,1]
func TwoSum(nums []int, target int) []int {
	var targetIndeces []int
	var m = make(map[int]int)

	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		j, k := m[target-v]
		if k {
			targetIndeces = append(targetIndeces, i)
			targetIndeces = append(targetIndeces, j)
		}
	}
	return targetIndeces
}
