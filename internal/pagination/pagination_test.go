package pagination

import "testing"

func TestTotalPages(t *testing.T) {
	tests := []struct {
		name  string
		count int
		size  int
		want  int
	}{
		{name: "exact multiple", count: 40, size: 20, want: 2},
		{name: "remainder rounds up", count: 41, size: 20, want: 3},
		{name: "single item", count: 1, size: 20, want: 1},
		{name: "no items", count: 0, size: 20, want: 0},
		{name: "count smaller than size", count: 5, size: 20, want: 1},
		{name: "default page size", count: 100, size: DefaultSize, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TotalPages(tt.count, tt.size)
			if got != tt.want {
				t.Errorf(
					"TotalPages(%d, %d) = %d, want %d",
					tt.count,
					tt.size,
					got,
					tt.want,
				)
			}
		})
	}
}
