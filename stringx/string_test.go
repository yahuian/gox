package stringx_test

import (
	"testing"

	"github.com/yahuian/gox/stringx"
)

func TestLimitRune(t *testing.T) {
	type args struct {
		s      string
		n      int
		suffix []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "s > n",
			args: args{
				s: "北国风光，千里冰封。hi",
				n: 4,
			},
			want: "北国风光...",
		},
		{
			name: "s = n",
			args: args{
				s: "北国风光，千里冰封。hi",
				n: 12,
			},
			want: "北国风光，千里冰封。hi",
		},
		{
			name: "s < n",
			args: args{
				s: "北国风光，千里冰封。hi",
				n: 15,
			},
			want: "北国风光，千里冰封。hi",
		},
		{
			name: "s < n with suffix",
			args: args{
				s:      "北国风光，千里冰封。hi",
				n:      10,
				suffix: []string{"***"},
			},
			want: "北国风光，千里冰封。***",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringx.LimitRune(tt.args.s, tt.args.n, tt.args.suffix...); got != tt.want {
				t.Errorf("LimitRune() = %v, want %v", got, tt.want)
			}
		})
	}
}
