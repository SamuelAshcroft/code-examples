package myutil

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func ReturnPositionMax(array []int) (int, int) {
	var max int = array[0]
	var position int = 0
	for i, value := range array {
		if max < value {
			max = value
			position = i
		}
	}
	return position, max
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func PopSlice(array []byte) (byte, []byte) {
	return array[len(array)-1], array[:len(array)-1]
}
