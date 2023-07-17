package utility

func Map[T, R any](target []T, transformer func(T) R) []R {
	value := make([]R, len(target))
	for i, v := range target {
		value[i] = transformer(v)
	}
	return value
}

func Filter[T any](target []T, checker func(T) bool) []T {
	var value []T
	for _, v := range target {
		if checker(v) {
			value = append(value, v)
		}
	}
	return value
}

func Reduce[T, R any](target []T, reducer func(R, T) R, initValue R) R {
	acc := initValue
	for _, v := range target {
		acc = reducer(acc, v)
	}
	return acc
}
