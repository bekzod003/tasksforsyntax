package main

func hadDuplicate(arr []int) bool {
	for i1, n1 := range arr {
		for i2, n2 := range arr {
			if n1 == n2 && i1 != i2 {
				return false
			}
		}
	}
	return true
}
