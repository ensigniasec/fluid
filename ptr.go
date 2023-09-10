package fluid

func ptr[T any](s T) *T {
	return &s
}
