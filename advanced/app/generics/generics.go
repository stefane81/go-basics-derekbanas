package generics

type MyConstraint interface {
	int | float64
}

func GetSumGen[T MyConstraint](x T, y T) T {
	return x + y
}
