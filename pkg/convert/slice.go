package convert

func MapSliceFunc[T, U any](s []T, f func(T) (U, error)) ([]U, error) {
	out := make([]U, len(s))
	for i, v := range s {
		u, err := f(v)
		if err != nil {
			return nil, err
		}
		out[i] = u
	}
	return out, nil
}
