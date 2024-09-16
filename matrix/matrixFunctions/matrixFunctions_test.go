package matrix

import (
	matrix "matrix/matrixStruct"
	"reflect"
	"testing"
)

//453

type testData[T matrix.Number] struct {
	list1 matrix.Matrix[T]
	list2 matrix.Matrix[T]
	want  matrix.Matrix[T]
}

func TestSum(t *testing.T) {
	tests := []testData[int]{
		testData[int]{
			list1: CreateMatrix[int]([][]int{{1, 2, 3}, {4, 5, 6}}),
			list2: CreateMatrix[int]([][]int{{1, 2, 3}, {4, 5, 6}}),
			want:  CreateMatrix[int]([][]int{{2, 4, 6}, {8, 10, 12}}),
		},
		testData[int]{
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
		testData[int]{
			list1: CreateMatrix[int]([][]int{{1, 2, 3}, {4, 5, 6}}),
			list2: CreateMatrix[int]([][]int{{1, 2, 3}, {4, 5, 6}}),
			want:  CreateMatrix[int]([][]int{{0, 0, 0}, {0, 0, 0}}),
		},
		testData[int]{
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
