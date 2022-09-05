package util

func IsNil[T any](obj T) bool {
	return (interface{})(obj) == (interface{})(nil)
}
func IsEmpty[T any](obj T) bool {
	return (interface{})(obj) == (interface{})("")
}

func DefaultIfNil[T any](obj T, defaultObj T) T {
	if !IsNil(obj) {
		return obj
	}
	return defaultObj
}
func DefaultIfEmpty[T any](obj T, defaultObj T) T {
	if IsNil(obj) || IsEmpty(obj) {
		return defaultObj
	}
	return obj
}
