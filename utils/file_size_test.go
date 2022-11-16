package utils

import (
	"testing"
)

func TestHumanFileSize(t *testing.T) {
	type args struct {
		size float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test 1", args{size: 8437092352}, "7.86 GB"},
		{"Test 2", args{size: 1527754752}, "1.42 GB"},
		{"Test 2", args{size: 6909337600}, "6.43 GB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HumanFileSize(tt.args.size); got != tt.want {
				t.Errorf("HumanFileSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
