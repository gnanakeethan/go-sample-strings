package gosample

import (
	"errors"
	"reflect"
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

func Test_averageNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 float64
	}{
		
		{name: "test1", args: args{s: "2-abc-34-abc"}, want: nil, want1: 18},
		{name: "test2", args: args{s: "abc-34-abc"}, want: errors.New("Invalid input string : " + "abc-34-abc"), want1: 0},
		{name: "test3", args: args{s: "-34-abc"}, want: errors.New("Invalid input string : " + "-34-abc"), want1: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := averageNumber(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageNumber() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("averageNumber() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
