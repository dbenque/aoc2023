package game3

import "testing"

func TestSearchSymbolOnLine(t *testing.T) {

	tests := []struct {
		line string
		b    int
		e    int
		want bool
	}{
		{
			line: "............",
			b:    1,
			e:    5,
			want: false,
		},
		{
			line: "%...........",
			b:    1,
			e:    5,
			want: true,
		},
		{
			line: ".45567464356...",
			b:    1,
			e:    5,
			want: false,
		},
		{
			line: ".1234%",
			b:    1,
			e:    5,
			want: true,
		},
		{
			line: ".12345",
			b:    1,
			e:    5,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			if got := SearchSymbolOnLine(tt.line, tt.b, tt.e); got != tt.want {
				t.Errorf("SearchSymbolOnLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
