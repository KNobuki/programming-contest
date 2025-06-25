package main

import (
	"reflect"
	"sort"
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

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a > b",
			args: args{a: 5, b: 3},
			want: 5,
		},
		{
			name: "a < b",
			args: args{a: 3, b: 5},
			want: 5,
		},
		{
			name: "a == b",
			args: args{a: 5, b: 5},
			want: 5,
		},
		{
			name: "negative numbers",
			args: args{a: -3, b: -5},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a > b",
			args: args{a: 5, b: 3},
			want: 3,
		},
		{
			name: "a < b",
			args: args{a: 3, b: 5},
			want: 3,
		},
		{
			name: "a == b",
			args: args{a: 5, b: 5},
			want: 5,
		},
		{
			name: "negative numbers",
			args: args{a: -3, b: -5},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive",
			args: args{n: 5},
			want: 5,
		},
		{
			name: "negative",
			args: args{n: -5},
			want: 5,
		},
		{
			name: "zero",
			args: args{n: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.n); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pow(t *testing.T) {
	type args struct {
		a int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2^3",
			args: args{a: 2, n: 3},
			want: 8,
		},
		{
			name: "3^4",
			args: args{a: 3, n: 4},
			want: 81,
		},
		{
			name: "5^0",
			args: args{a: 5, n: 0},
			want: 1,
		},
		{
			name: "10^1",
			args: args{a: 10, n: 1},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args.a, tt.args.n); got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "gcd(12, 8)",
			args: args{x: 12, y: 8},
			want: 4,
		},
		{
			name: "gcd(14, 21)",
			args: args{x: 14, y: 21},
			want: 7,
		},
		{
			name: "gcd(17, 19)",
			args: args{x: 17, y: 19},
			want: 1,
		},
		{
			name: "gcd with 0",
			args: args{x: 5, y: 0},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lcm(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lcm(12, 8)",
			args: args{x: 12, y: 8},
			want: 24,
		},
		{
			name: "lcm(14, 21)",
			args: args{x: 14, y: 21},
			want: 42,
		},
		{
			name: "lcm(7, 5)",
			args: args{x: 7, y: 5},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcm(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfInts(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal array",
			args: args{a: []int{1, 2, 3, 4, 5}},
			want: 15,
		},
		{
			name: "empty array",
			args: args{a: []int{}},
			want: 0,
		},
		{
			name: "single element",
			args: args{a: []int{42}},
			want: 42,
		},
		{
			name: "negative numbers",
			args: args{a: []int{-1, -2, -3}},
			want: -6,
		},
		{
			name: "mixed positive and negative",
			args: args{a: []int{1, -2, 3, -4, 5}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfInts(tt.args.a); got != tt.want {
				t.Errorf("sumOfInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxOfInts(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal array",
			args: args{a: []int{1, 5, 3, 9, 2}},
			want: 9,
		},
		{
			name: "all negative",
			args: args{a: []int{-5, -2, -8, -1}},
			want: -1,
		},
		{
			name: "single element",
			args: args{a: []int{42}},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxOfInts(tt.args.a); got != tt.want {
				t.Errorf("maxOfInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minOfInts(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal array",
			args: args{a: []int{1, 5, 3, 9, 2}},
			want: 1,
		},
		{
			name: "all negative",
			args: args{a: []int{-5, -2, -8, -1}},
			want: -8,
		},
		{
			name: "single element",
			args: args{a: []int{42}},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOfInts(tt.args.a); got != tt.want {
				t.Errorf("minOfInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniqueInts(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "with duplicates",
			args: args{a: []int{1, 2, 3, 2, 1, 4}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "no duplicates",
			args: args{a: []int{1, 2, 3, 4}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "all same",
			args: args{a: []int{5, 5, 5, 5}},
			want: []int{5},
		},
		{
			name: "empty",
			args: args{a: []int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueInts(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniqueInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniqueStrings(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "with duplicates",
			args: args{a: []string{"apple", "banana", "apple", "cherry"}},
			want: []string{"apple", "banana", "cherry"},
		},
		{
			name: "no duplicates",
			args: args{a: []string{"apple", "banana", "cherry"}},
			want: []string{"apple", "banana", "cherry"},
		},
		{
			name: "empty",
			args: args{a: []string{}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueStrings(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniqueStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2 is prime",
			args: args{x: 2},
			want: true,
		},
		{
			name: "3 is prime",
			args: args{x: 3},
			want: true,
		},
		{
			name: "4 is not prime",
			args: args{x: 4},
			want: false,
		},
		{
			name: "17 is prime",
			args: args{x: 17},
			want: true,
		},
		{
			name: "1 is not prime",
			args: args{x: 1},
			want: false,
		},
		{
			name: "0 is not prime",
			args: args{x: 0},
			want: false,
		},
		{
			name: "negative is not prime",
			args: args{x: -5},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.x); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_primeFact(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "12 = 2^2 * 3",
			args: args{x: 12},
			want: map[int]int{2: 2, 3: 1},
		},
		{
			name: "60 = 2^2 * 3 * 5",
			args: args{x: 60},
			want: map[int]int{2: 2, 3: 1, 5: 1},
		},
		{
			name: "17 is prime",
			args: args{x: 17},
			want: map[int]int{17: 1},
		},
		{
			name: "1",
			args: args{x: 1},
			want: map[int]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := primeFact(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("primeFact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divisor(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "divisors of 12",
			args: args{x: 12},
			want: []int{1, 12, 2, 6, 3, 4},
		},
		{
			name: "divisors of 1",
			args: args{x: 1},
			want: []int{1},
		},
		{
			name: "divisors of 16",
			args: args{x: 16},
			want: []int{1, 16, 2, 8, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divisor(tt.args.x)
			// Sort both slices for comparison since order might vary
			sort.Ints(got)
			want := make([]int, len(tt.want))
			copy(want, tt.want)
			sort.Ints(want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("divisor() = %v, want %v", got, want)
			}
		})
	}
}

func Test_Dsu(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		d := NewDsu(5)

		// Initially all elements are separate
		if d.Same(0, 1) {
			t.Errorf("Initially 0 and 1 should not be in same set")
		}

		// Merge 0 and 1
		d.Merge(0, 1)
		if !d.Same(0, 1) {
			t.Errorf("After merge, 0 and 1 should be in same set")
		}

		// Size check
		if got := d.Size(0); got != 2 {
			t.Errorf("Size of set containing 0 = %v, want 2", got)
		}

		// Merge 2 and 3
		d.Merge(2, 3)

		// Merge two groups
		d.Merge(1, 2)
		if !d.Same(0, 3) {
			t.Errorf("After merging groups, 0 and 3 should be in same set")
		}
		if got := d.Size(0); got != 4 {
			t.Errorf("Size of merged set = %v, want 4", got)
		}
	})

	t.Run("groups", func(t *testing.T) {
		d := NewDsu(5)
		d.Merge(0, 2)
		d.Merge(1, 3)

		groups := d.Groups()
		if len(groups) != 3 { // {0,2}, {1,3}, {4}
			t.Errorf("Number of groups = %v, want 3", len(groups))
		}
	})
}

func Test_reverseInts(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal array",
			args: args{a: []int{1, 2, 3, 4, 5}},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "empty array",
			args: args{a: []int{}},
			want: []int{},
		},
		{
			name: "single element",
			args: args{a: []int{42}},
			want: []int{42},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseInts(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseString(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal string",
			args: args{a: "hello"},
			want: "olleh",
		},
		{
			name: "empty string",
			args: args{a: ""},
			want: "",
		},
		{
			name: "single char",
			args: args{a: "a"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.args.a); got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_madd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal addition",
			args: args{a: 10, b: 20},
			want: 30,
		},
		{
			name: "with modulo",
			args: args{a: 998244350, b: 5},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := madd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("madd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_msub(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal subtraction",
			args: args{a: 30, b: 10},
			want: 20,
		},
		{
			name: "negative result",
			args: args{a: 5, b: 10},
			want: 998244348,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := msub(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("msub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mmul(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal multiplication",
			args: args{a: 10, b: 20},
			want: 200,
		},
		{
			name: "large numbers",
			args: args{a: 1000000, b: 1000000},
			want: 757402647, // (10^6 * 10^6) % mod
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mmul(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("mmul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mpow(t *testing.T) {
	type args struct {
		a int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2^10",
			args: args{a: 2, n: 10},
			want: 1024,
		},
		{
			name: "3^0",
			args: args{a: 3, n: 0},
			want: 1,
		},
		{
			name: "large power",
			args: args{a: 2, n: 30},
			want: 75497471, // 2^30 % mod
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mpow(tt.args.a, tt.args.n); got != tt.want {
				t.Errorf("mpow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eratosthenes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name      string
		args      args
		wantPrime []int // indices that should be true
	}{
		{
			name:      "primes up to 20",
			args:      args{n: 20},
			wantPrime: []int{2, 3, 5, 7, 11, 13, 17, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := eratosthenes(tt.args.n)
			for _, p := range tt.wantPrime {
				if !got[p] {
					t.Errorf("eratosthenes() marked %v as non-prime, but it is prime", p)
				}
			}
			// Check some non-primes
			nonPrimes := []int{0, 1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20}
			for _, np := range nonPrimes {
				if np <= tt.args.n && got[np] {
					t.Errorf("eratosthenes() marked %v as prime, but it is not prime", np)
				}
			}
		})
	}
}

func Test_nextPermutation(t *testing.T) {
	type args struct {
		aa []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantRet bool
	}{
		{
			name:    "normal permutation",
			args:    args{aa: []int{1, 2, 3}},
			want:    []int{1, 3, 2},
			wantRet: true,
		},
		{
			name:    "last permutation",
			args:    args{aa: []int{3, 2, 1}},
			want:    []int{3, 2, 1},
			wantRet: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aa := make([]int, len(tt.args.aa))
			copy(aa, tt.args.aa)
			gotRet := nextPermutation(aa)
			if gotRet != tt.wantRet {
				t.Errorf("nextPermutation() returned %v, want %v", gotRet, tt.wantRet)
			}
			if !reflect.DeepEqual(aa, tt.want) {
				t.Errorf("nextPermutation() modified array to %v, want %v", aa, tt.want)
			}
		})
	}
}

func Test_genPerm(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "size 5",
			args: args{size: 5},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "size 0",
			args: args{size: 0},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genPerm(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genPerm() = %v, want %v", got, tt.want)
			}
		})
	}
}
