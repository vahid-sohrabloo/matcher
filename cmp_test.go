package matcher_test

import (
	"testing"

	"github.com/vahid-sohrabloo/matcher"
)

func TestCmp(t *testing.T) {
	tests := []struct {
		name  string
		x     interface{}
		y     interface{}
		match bool
	}{
		{
			x: map[int]int{
				1: 1,
				2: 2,
			},
			y: map[int]int{
				1: 1,
				2: 2,
			},
			match: true,
		},
		{
			x: map[int]int{
				1: 1,
				2: 2,
			},
			y: map[int]int{
				1: 1,
			},
			match: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matcher.Cmp(t, tt.x); !got.Matches(tt.y) == tt.match {
				t.Errorf("matcher.Cmp(%v).Matches(%v)", tt.x, tt.y)
			}
		})
	}
}

func Test_cmpMatcher_String(t *testing.T) {
	tests := []struct {
		name string
		x    interface{}
		want string
	}{
		{
			x: map[int]int{
				1: 1,
			},
			want: "is equal to map[1:1] with Options{}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matcher.Cmp(t, tt.x).String(); got != tt.want {
				t.Errorf("cmpMatcher.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
