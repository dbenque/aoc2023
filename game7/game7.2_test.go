package game7

import (
	"reflect"
	"testing"
)

func TestDecodeHand2(t *testing.T) {

	tests := []struct {
		name     string
		wantType string
	}{
		{
			"QJQAQ",
			"41",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeHandStr(tt.name, DecodeHand2); !reflect.DeepEqual(got.Type, tt.wantType) {
				t.Errorf("DecodeHand2() = %v, want %v", got.Type, tt.wantType)
			}
		})
	}
}
