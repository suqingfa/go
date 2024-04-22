package slices

func MapFunc[S ~[]E, E any, T any](s S, f func(E) T) []T {
	res := make([]T, len(s))
	for i, e := range s {
		res[i] = f(e)
	}
	return res
}

func Count[S ~[]E, E comparable](s S, e E) int {
	return CountFunc(s, func(v E) bool {
		return e == v
	})
}

func CountFunc[S ~[]E, E any](s S, check func(E) bool) int {
	res := 0
	for _, e := range s {
		if check(e) {
			res++
		}
	}
	return res
}

func Indexes[S ~[]E, E comparable](s S, e E) []int {
	return IndexesFunc(s, func(v E) bool {
		return v == e
	})
}

func IndexesFunc[S ~[]E, E any](s S, check func(E) bool) []int {
	res := make([]int, 0)
	for i, v := range s {
		if check(v) {
			res = append(res, i)
		}
	}
	return res
}
