package main

func OddEvenSum(arr []int) (int, int) {
	var oddSum, evenSum int

	for _, n := range arr {
		if n%2 == 0 {
			evenSum += n
		} else {
			oddSum += n
		}
	}
	return oddSum, evenSum
}
