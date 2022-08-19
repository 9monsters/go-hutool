package util

func IsNil[T any](obj T) bool {
	return obj == nil
}

func DefaultIfNil[T any](obj T, defaultObj T) T {
	if IsNil(obj) {
		return obj
	}
	return defaultObj
}
