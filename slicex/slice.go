package slicex

func Unique[T comparable](list []T) []T {
	m := make(map[T]struct{}, len(list))

	for _, v := range list {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
	}

	res := make([]T, 0, len(m))
	for k := range m {
		res = append(res, k)
	}

	return res
}

func Contain[T comparable](list []T, key T) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}
