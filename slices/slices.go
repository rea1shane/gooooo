package slices

// Map 将 slice 转换成 map，判断 contains 效率更高。
func Map[T comparable](items []T) (result map[T]struct{}) {
	if length := len(items); length > 0 {
		result = make(map[T]struct{}, length)

		for _, item := range items {
			result[item] = struct{}{}
		}
	}
	return
}
