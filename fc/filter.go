package fc

func Filter[T any](array []T, callback func(T) bool) []T {
	var result []T
	for _, value := range array {
		if callback(value) {
			result = append(result, value)
		}
	}
	return result
}
