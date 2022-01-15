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

func Test_wholeStory(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 string
	}{
		{name: "test1", args: args{s: "2-abc-34-abc"}, want: nil, want1: "abc abc"},
		{name: "test2", args: args{s: "abc-34-abc"}, want: errors.New("Invalid input string : " + "abc-34-abc"), want1: ""},
		{name: "test3", args: args{s: "-34-abc"}, want: errors.New("Invalid input string : " + "-34-abc"), want1: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := wholeStory(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wholeStory() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("wholeStory() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_storyStats(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 string
		want2 string
		want3 float64
		want4 []string
	}{
		{name: "test1", args: args{s: "2-abc-34-abc-23-ab"}, want: nil, want1: "ab", want2: "abc", want3: 8.0 / 3.0, want4: []string{"abc", "abc", "ab"}},
		{name: "test2", args: args{s: "abc-34-abc"}, want: errors.New("Invalid input string : " + "abc-34-abc"), want1: "", want2: "", want3: 0, want4: []string{}},
		{name: "test3", args: args{s: "-34-abc"}, want: errors.New("Invalid input string : " + "-34-abc"), want1: "", want2: "", want3: 0, want4: []string{}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, got4 := storyStats(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("storyStats() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("storyStats() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("storyStats() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("storyStats() got3 = %v, want %v", got3, tt.want3)
			}
			if !reflect.DeepEqual(got4, tt.want4) {
				t.Errorf("storyStats() got4 = %v, want %v", got4, tt.want4)
			}
		})
	}
}
