package utils

type Comparator func(a, b interface{}) int

func IntComparator(a, b interface{}) int {
	aVal := a.(int)
	bVal := b.(int)
	if aVal > bVal {
		return 1
	} else if aVal < bVal {
		return -1
	} else {
		return 0
	}
}
