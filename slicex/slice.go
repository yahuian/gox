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

func Map[T any, U any](list []T, f func(index int, value T) U) []U {
	res := make([]U, 0, len(list))
	for i, v := range list {
		res = append(res, f(i, v))
	}
	return res
}

func Filter[T any](list []T, f func(index int, value T) bool) []T {
	res := make([]T, 0)
	for i, v := range list {
		if f(i, v) {
			res = append(res, v)
		}
	}
	return res
}

func Reduce[T any](list []T, initial T, f func(index int, result, value T) T) T {
	res := initial
	for i, v := range list {
		res = f(i, res, v)
	}
	return res
}

func Reverse[T any](list []T) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}
