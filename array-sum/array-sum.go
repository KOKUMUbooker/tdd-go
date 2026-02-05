package main

func ArraySum(n [5]int) int {
	res := 0
	for _, v := range n {
		res += v
	}

	return res
}
