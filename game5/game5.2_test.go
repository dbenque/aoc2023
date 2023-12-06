package game5

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInterseptRange(t *testing.T) {
	tests := []struct {
		name    string
		r       *RangeItem
		spliter *RangeItem
		want    *RangeItem
	}{
		{
			name:    "1",
			r:       &RangeItem{Start: 10, End: 20},
			spliter: &RangeItem{Start: 1, End: 5},
			want:    nil,
		},
		{
			name:    "2",
			r:       &RangeItem{Start: 10, End: 20},
			spliter: &RangeItem{Start: 1, End: 15},
			want:    &RangeItem{Start: 10, End: 15},
		},
		{
			name:    "3",
			r:       &RangeItem{Start: 10, End: 20},
			spliter: &RangeItem{Start: 15, End: 25},
			want:    &RangeItem{Start: 15, End: 20},
		},
		{
			name:    "4",
			r:       &RangeItem{Start: 10, End: 20},
			spliter: &RangeItem{Start: 25, End: 35},
			want:    nil,
		},
		{
			name:    "5",
			r:       &RangeItem{Start: 10, End: 20},
			spliter: &RangeItem{Start: 2, End: 35},
			want:    &RangeItem{Start: 10, End: 20},
		},
		{
			name:    "6",
			r:       &RangeItem{Start: 10, End: 10},
			spliter: &RangeItem{Start: 10, End: 10},
			want:    &RangeItem{Start: 10, End: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterceptRange(tt.r, tt.spliter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transformRange(t *testing.T) {
	tests := []struct {
		name          string
		r             *RangeItem
		mappersRanges []*RangeItem
		want          []*RangeItem
	}{
		{
			name:          "1",
			r:             &RangeItem{Start: 10, End: 20},
			mappersRanges: []*RangeItem{{Start: 12, End: 15}},
			want:          []*RangeItem{{Start: 10, End: 12}, {Start: 12, End: 15}, {Start: 16, End: 20}},
		},
		{
			name:          "2",
			r:             &RangeItem{Start: 10, End: 20},
			mappersRanges: []*RangeItem{{Start: 1, End: 5}},
			want:          []*RangeItem{{Start: 10, End: 10}},
		},
		{
			name:          "3",
			r:             &RangeItem{Start: 10, End: 20},
			mappersRanges: []*RangeItem{{Start: 1, End: 15}},
			want:          []*RangeItem{{Start: 10, End: 15}, {Start: 16, End: 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transformRange(tt.r, tt.mappersRanges); !reflect.DeepEqual(got, tt.want) {
				fmt.Printf("%s - Got : %s\n", tt.name, PrintRangeItems(got))
				fmt.Printf("%s - Want: %s\n", tt.name, PrintRangeItems(tt.want))
				//t.Errorf("transformRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func PrintRangeItem(r *RangeItem) string {
	return fmt.Sprintf("{%d,%d}", r.Start, r.End)
}
func PrintRangeItems(rs []*RangeItem) (s string) {
	for _, r := range rs {
		s = s + PrintRangeItem(r)
	}
	return
}
