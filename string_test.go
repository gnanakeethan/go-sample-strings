package gosample

import (
	"testing"
)

func Test_testValidity(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{s: "2-abc-34-abc"}, want: true},
		{name: "test2", args: args{s: "abc-34-abc"}, want: false},
		{name: "test3", args: args{s: "-34-abc"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := testValidity(tt.args.s); got != tt.want {
				t.Errorf("testValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}
