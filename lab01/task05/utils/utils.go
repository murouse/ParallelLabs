package utils

type Predicate[T any] func(x *T) bool

type Predicates[T any] []Predicate[T]

func Remove[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func RemoveIf[T any](s []T, predicate Predicate[T]) []T {
	for i := 0; i < len(s); i++ {
		if predicate(&s[i]) {
			s = Remove(s, i)
			i--
		}
	}
	return s
}
