package core

func IsType[T any](val any) bool {
	_, ok := val.(T)
	return ok
}
