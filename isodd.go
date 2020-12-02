package isodd

func isodd(x int) bool {
	result:= true
	if x % 2  == 0 {
		result = false
	}
	return result
}