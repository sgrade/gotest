package checkifall1sareatleastlengthkplacesaway

import "testing"

func TestKLengthApart(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 0, 0, 0, 1, 0, 0, 1}, 2, true},
		{[]int{1, 0, 0, 1, 0, 1}, 2, false},
	}

	for _, test := range tests {
		got := kLengthApart(test.nums, test.k)
		if got != test.want {
			t.Errorf("kLengthApart(%v, %d) = %v, want %v", test.nums, test.k, got, test.want)
		}
	}
}
