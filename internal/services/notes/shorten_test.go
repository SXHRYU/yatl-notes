package notes

import (
	"strings"
	"testing"
)

func TestShortenNote(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantRunes    int // expected rune count of the result, EXCLUDING any "..." suffix
		wantEllipsis bool
	}{
		{
			name:         "empty string is left untouched",
			input:        "",
			wantRunes:    0,
			wantEllipsis: false,
		},
		{
			name:         "short ascii text is left untouched",
			input:        strings.Repeat("a", 10),
			wantRunes:    10,
			wantEllipsis: false,
		},
		{
			name:         "short cyrillic text is left untouched",
			input:        strings.Repeat("а", 5),
			wantRunes:    5,
			wantEllipsis: false,
		},
		{
			name:         "ascii text right at the boundary (392 runes) is untouched",
			input:        strings.Repeat("a", 392),
			wantRunes:    392,
			wantEllipsis: false,
		},
		{
			name:         "ascii text one rune past the boundary (393 runes) gets cut",
			input:        strings.Repeat("a", 393),
			wantRunes:    393,
			wantEllipsis: true,
		},
		{
			name:         "long ascii text gets cut to 393 runes + ellipsis",
			input:        strings.Repeat("a", 500),
			wantRunes:    393,
			wantEllipsis: true,
		},
		{
			name:         "long cyrillic (2-byte) text gets cut much earlier, at 197 runes",
			input:        strings.Repeat("а", 500),
			wantRunes:    197,
			wantEllipsis: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.input
			shortenNote(&s)

			got := s
			hasEllipsis := strings.HasSuffix(got, "...")
			if hasEllipsis != tt.wantEllipsis {
				t.Fatalf(
					"ellipsis present = %v, want %v (result: %q)",
					hasEllipsis,
					tt.wantEllipsis,
					got,
				)
			}

			core := got
			if hasEllipsis {
				core = strings.TrimSuffix(got, "...")
			}
			gotRunes := len([]rune(core))
			if gotRunes != tt.wantRunes {
				t.Errorf(
					"rune count (excluding ellipsis) = %d, want %d",
					gotRunes,
					tt.wantRunes,
				)
			}

			if !tt.wantEllipsis && len([]rune(tt.input)) != gotRunes {
				t.Errorf(
					"text was truncated (%d -> %d runes) without appending an ellipsis",
					len([]rune(tt.input)),
					gotRunes,
				)
			}
		})
	}
}
