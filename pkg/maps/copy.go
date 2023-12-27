package maps

func Copied[K comparable, V any](orig map[K]V) map[K]V {
	copied := make(map[K]V, len(orig))

	for k, v := range orig {
		copied[k] = v
	}

	return copied
}
