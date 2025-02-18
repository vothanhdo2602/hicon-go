package ptr

func ToPtr[T any](v T) *T {
	return &v
}

func GetValue[T any](v *T) (output *T) {
	if v == nil {
		return output
	}
	return v
}
