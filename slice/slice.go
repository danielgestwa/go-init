package main

func Sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func SumElements(arr_one []int, arr_two []int) []int {
	var sum []int = []int{}

	arr_one_t := arr_one
	arr_two_t := arr_two
	if len(arr_two) > len(arr_one) {
		arr_one_t = arr_two
		arr_two_t = arr_one
	}

	for key, val := range arr_one_t {
		if len(arr_two_t) > key {
			sum = append(sum, val+arr_two_t[key])
		} else {
			sum = append(sum, val)
		}
	}
	return sum
}
