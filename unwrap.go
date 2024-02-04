package unwrap

// Unwrap returns the most underlying value of w.
func Unwrap[T any](w T) T {
	var t any = w
	for {
		u, ok := (t).(interface{ Unwrap() T })
		if !ok {
			return t.(T)
		}
		t = u.Unwrap()
	}
}
