package main

import (
	"reflect"
	"testing"
)

func Test_runLengthEncoding(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []runLength
	}{
		{
			name: "success",
			args: args{
				s: "aaabbbccc",
			},
			want: []runLength{
				{
					c: 'a',
					l: 3,
				},
				{
					c: 'b',
					l: 3,
				},
				{
					c: 'c',
					l: 3,
				},
			},
		},
		{
			name: "一文字",
			args: args{
				s: "a",
			},
			want: []runLength{
				{
					c: 'a',
					l: 1,
				},
			},
		},
		{
			name: "末尾が一文字違う",
			args: args{
				s: "aaabbbcccd",
			},
			want: []runLength{
				{
					c: 'a',
					l: 3,
				},
				{
					c: 'b',
					l: 3,
				},
				{
					c: 'c',
					l: 3,
				},
				{
					c: 'd',
					l: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runLengthEncoding(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runLengthEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runLengthDecoding(t *testing.T) {
	type args struct {
		rl []runLength
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				rl: []runLength{
					{
						c: 'a',
						l: 3,
					},
					{
						c: 'b',
						l: 3,
					},
					{
						c: 'c',
						l: 3,
					},
				},
			},
			want: "aaabbbccc",
		},
		{
			name: "一文字",
			args: args{
				rl: []runLength{
					{
						c: 'a',
						l: 1,
					},
				},
			},
			want: "a",
		},
		{
			name: "末尾が一文字違う",
			args: args{
				rl: []runLength{
					{
						c: 'a',
						l: 3,
					},
					{
						c: 'b',
						l: 3,
					},
					{
						c: 'c',
						l: 3,
					},
					{
						c: 'd',
						l: 1,
					},
				},
			},
			want: "aaabbbcccd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runLengthDecoding(tt.args.rl); got != tt.want {
				t.Errorf("runLengthDecoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zAlgorithm(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				s: "ababaababaabababc",
			},
			want: []int{17, 0, 3, 0, 1, 10, 0, 3, 0, 1, 5, 0, 4, 0, 2, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zAlgorithm(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}
