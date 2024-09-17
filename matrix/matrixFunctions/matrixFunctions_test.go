package matrix

import (
	matrix "matrix/matrixStruct"
	inter "matrix/numInterface"
	"reflect"
	"testing"
)

//453

type testData[T inter.Number] struct {
	list1 matrix.Matrix[T]
	list2 matrix.Matrix[T]
	val   T
	want  matrix.Matrix[T]
}

func TestSum(t *testing.T) {
	tests := []testData[int]{
		{
			list1: CreateMatrix[int]([][]int{{1, 2, 3}, {4, 5, 6}}),
			list2: CreateMatrix[int]([][]int{{1, 2, 3}, {4, 5, 6}}),
			want:  CreateMatrix[int]([][]int{{2, 4, 6}, {8, 10, 12}}),
		},
		{
			list1: CreateMatrix[int]([][]int{{1, 2}, {4, 5}}),
			list2: CreateMatrix[int]([][]int{{-2, -3}, {-10, 5}}),
			want:  CreateMatrix[int]([][]int{{-1, -1}, {-6, 10}}),
		},
	}

	for _, test := range tests {
		got := Sum[int](test.list1, test.list2)
		if !reflect.DeepEqual(got.Matrix(), test.want.Matrix()) {
			t.Errorf("Sum (%d, %d) = \"%d\", want \"%d\"", test.list1.Matrix(), test.list2.Matrix(), got.Matrix(), test.want.Matrix())
		}
	}
}

func TestDifference(t *testing.T) {
	tests := []testData[int]{
		{
			list1: CreateMatrix[int]([][]int{{1, 2, 3}, {4, 5, 6}}),
			list2: CreateMatrix[int]([][]int{{1, 2, 3}, {4, 5, 6}}),
			want:  CreateMatrix[int]([][]int{{0, 0, 0}, {0, 0, 0}}),
		},
		{
			list1: CreateMatrix[int]([][]int{{1, 2}, {4, -5}}),
			list2: CreateMatrix[int]([][]int{{-2, -3}, {-10, 5}}),
			want:  CreateMatrix[int]([][]int{{3, 5}, {14, -10}}),
		},
	}

	for _, test := range tests {
		got := Difference[int](test.list1, test.list2)
		if !reflect.DeepEqual(got.Matrix(), test.want.Matrix()) {
			t.Errorf("Difference (%d, %d) = \"%d\", want \"%d\"", test.list1.Matrix(), test.list2.Matrix(), got.Matrix(), test.want.Matrix())
		}
	}
}

func TestProduct(t *testing.T) {
	tests := []testData[int]{
		{
			list1: CreateMatrix[int]([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}),
			list2: CreateMatrix[int]([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}),
			want:  CreateMatrix[int]([][]int{{30, 36, 42}, {66, 81, 96}, {102, 126, 150}}),
		},
		{
			list1: CreateMatrix[int]([][]int{{1, 2}, {4, -5}}),
			list2: CreateMatrix[int]([][]int{{-2, -3}, {-10, 5}}),
			want:  CreateMatrix[int]([][]int{{-22, 7}, {42, -37}}),
		},
	}

	for _, test := range tests {
		got := Product[int](test.list1, test.list2)
		if !reflect.DeepEqual(got.Matrix(), test.want.Matrix()) {
			t.Errorf("Product (%d, %d) = \"%d\", want \"%d\"", test.list1.Matrix(), test.list2.Matrix(), got.Matrix(), test.want.Matrix())
		}
	}
}

func TestMultiplicationByNumber(t *testing.T) {
	tests := []testData[int]{
		{
			list1: CreateMatrix[int]([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}),
			val:   -5,
			want:  CreateMatrix[int]([][]int{{-5, -10, -15}, {-20, -25, -30}, {-35, -40, -45}}),
		},
		{
			list1: CreateMatrix[int]([][]int{{1, 2}, {4, -5}}),
			val:   0,
			want:  CreateMatrix[int]([][]int{{0, 0}, {0, 0}}),
		},
	}

	for _, test := range tests {
		got := MultiplicationByNumber[int](test.list1, test.val)
		if !reflect.DeepEqual(got.Matrix(), test.want.Matrix()) {
			t.Errorf("MultiplicationByNumbe (%d, %d) = \"%d\", want \"%d\"", test.list1.Matrix(), test.val, got.Matrix(), test.want.Matrix())
		}
	}
}

func TestTransposition(t *testing.T) {
	tests := []testData[int]{
		{
			list1: CreateMatrix[int]([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}),
			want:  CreateMatrix[int]([][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}),
		},
		{
			list1: CreateMatrix[int]([][]int{{1, 2}, {4, -5}}),
			want:  CreateMatrix[int]([][]int{{1, 4}, {2, -5}}),
		},
	}

	for _, test := range tests {
		got := Transposition[int](test.list1)
		if !reflect.DeepEqual(got.Matrix(), test.want.Matrix()) {
			t.Errorf("Transposition (%d) = \"%d\", want \"%d\"", test.list1.Matrix(), got.Matrix(), test.want.Matrix())
		}
	}
}
