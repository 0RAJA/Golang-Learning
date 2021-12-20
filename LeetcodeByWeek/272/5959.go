package main

func kIncreasing(arr []int, k int) (ret int) {
	for i := 1; i <= k; i++ {
		for j := i - 1; j < len(arr); j += k {
			if j-k >= 0 && arr[j] < arr[j-k] {
				ret++
			}
		}
	}
	return
}
