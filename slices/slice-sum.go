package slices


func SliceSum(numbers []int) int{
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}


func SumAll(numberSlice ...[]int) []int {
	var sub_array_sum []int
	for _, el := range numberSlice {
		sub_array_sum = append(sub_array_sum, SliceSum(el))
	}
	return sub_array_sum
}

func SumAllTails(numberSlice ...[]int) []int{
	var sub_array_sum []int
	for _, el := range numberSlice {
		if len(el) == 0 {
			sub_array_sum = append(sub_array_sum, 0)
		}else {
		sub_array_sum = append(sub_array_sum, SliceSum(el[1:]))
		}
	}
	return sub_array_sum
}