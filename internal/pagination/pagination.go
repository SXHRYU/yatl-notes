package pagination

const DefaultSize = 20

func TotalPages(count, size int) int {
	return (count + size - 1) / size
}
