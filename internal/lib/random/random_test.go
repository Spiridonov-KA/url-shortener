package random

import "testing"

func TestNewRandomString_Length(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		size int
	}{
		{name: "zero", size: 0},
		{name: "one", size: 1},
		{name: "six", size: 6},
		{name: "sixty_four", size: 64},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := NewRandomString(tt.size)
			if len(got) != tt.size {
				t.Fatalf("expected length %d, got %d (%q)", tt.size, len(got), got)
			}
		})
	}
}