package fibonacci_test

import (
	"fibonacci"
	"testing"
)

func TestLoop(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			name: "5",
			args: args{
				n: 5,
			},
			want: 5,
		},
		{
			name: "6",
			args: args{
				n: 6,
			},
			want: 8,
		},
		{
			name: "7",
			args: args{
				n: 7,
			},
			want: 13,
		},
		{
			name: "40",
			args: args{
				n: 40,
			},
			want: 102334155,
		},
		{
			name: "50",
			args: args{
				n: 50,
			},
			want: 12586269025,
		},
		{
			name: "75",
			args: args{
				n: 75,
			},
			want: 2111485077978050,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := fibonacci.Loop(tt.args.n); got != tt.want {
				t.Errorf("Loop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursive(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				n: 0,
			},
			want: 0,
		}, {
			name: "1",
			args: args{
				n: 1,
			},
			want: 1,
		}, {
			name: "2",
			args: args{
				n: 2,
			},
			want: 1,
		}, {
			name: "3",
			args: args{
				n: 3,
			},
			want: 2,
		}, {
			name: "4",
			args: args{
				n: 4,
			},
			want: 3,
		}, {
			name: "5",
			args: args{
				n: 5,
			},
			want: 5,
		}, {
			name: "6",
			args: args{
				n: 6,
			},
			want: 8,
		}, {
			name: "7",
			args: args{
				n: 7,
			},
			want: 13,
		}, {
			name: "40",
			args: args{
				n: 40,
			},
			want: 102334155,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := fibonacci.Recursive(tt.args.n); got != tt.want {
				t.Errorf("Recursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursiveCache(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			name: "5",
			args: args{
				n: 5,
			},
			want: 5,
		},
		{
			name: "6",
			args: args{
				n: 6,
			},
			want: 8,
		},
		{
			name: "7",
			args: args{
				n: 7,
			},
			want: 13,
		},
		{
			name: "40",
			args: args{
				n: 40,
			},
			want: 102334155,
		},
		{
			name: "50",
			args: args{
				n: 50,
			},
			want: 12586269025,
		},
		{
			name: "75",
			args: args{
				n: 75,
			},
			want: 2111485077978050,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := fibonacci.RecursiveCache(tt.args.n); got != tt.want {
				t.Errorf("RecursiveCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

var result int

func BenchmarkLoop20(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = fibonacci.Loop(20)
	}

	result = r
}

func BenchmarkRecursive20(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = fibonacci.Recursive(20)
	}

	result = r
}

func BenchmarkRecursiveCache20(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = fibonacci.RecursiveCache(20)
	}

	result = r
}

func BenchmarkLoop30(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = fibonacci.Loop(30)
	}

	result = r
}

func BenchmarkRecursive30(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = fibonacci.Recursive(30)
	}

	result = r
}

func BenchmarkRecursiveCache30(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = fibonacci.RecursiveCache(30)
	}

	result = r
}

func BenchmarkLoop40(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = fibonacci.Loop(40)
	}

	result = r
}

func BenchmarkRecursive40(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = fibonacci.Recursive(40)
	}

	result = r
}

func BenchmarkRecursiveCache40(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = fibonacci.RecursiveCache(40)
	}

	result = r
}
