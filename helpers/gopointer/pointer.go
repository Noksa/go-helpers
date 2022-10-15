package gopointer

func NewOf[T any](val T) *T {
	t := new(T)
	*t = val
	return t
}
