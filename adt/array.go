package adt

/*
判断该字符串是否在数组中
*/
func InArray[T comparable](value T, array []T) bool {
	for _, i := range array {
		if i == value {
			return true
		}
	}
	return false
}

/*
判断两个字符串数组是否相等
长度相同
内容相同
不要求内容相同
*/
func IsSameArray[T comparable](array1 []T, array2 []T) bool {
	if len(array1) != len(array2) {
		return false
	}
	for _, i := range array1 {
		if !InArray(i, array2) {
			return false
		}
	}
	return true
}
