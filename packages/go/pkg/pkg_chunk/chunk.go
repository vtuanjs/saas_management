package pkgchunk

func Split[T any](items []T, limit int) [][]T {
	if limit <= 0 {
		return [][]T{items}
	}

	var chunks [][]T
	for i := 0; i < len(items); i += limit {
		end := i + limit
		if end > len(items) {
			end = len(items)
		}
		chunks = append(chunks, items[i:end])
	}
	return chunks
}
