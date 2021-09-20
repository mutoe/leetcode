package sqrt_decomposition

import "testing"

func TestSqrtDecomposition(t *testing.T) {
	t.Run("Sqrt decomposition", func(t *testing.T) {
		nums := []value{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
		sqrt := Constructor(nums)

		if got := sqrt.sumRange(0, 9); got != 55 {
			t.Errorf("sumRange(0, 9) = %v, want 55", got)
		}
		if got := sqrt.sumRange(6, 17); got != 150 {
			t.Errorf("sumRange(6, 17) = %v, want 150", got)
		}

		sqrt.update(1, 10)
		if got := sqrt.sumRange(1, 8); got != 52 {
			t.Errorf("sumRange(1, 8) = %v, want 52", got)
		}
		sqrt.update(16, 200)
		if got := sqrt.sumRange(11, 17); got != 288 {
			t.Errorf("sumRange(11, 17) = %v, want 288", got)
		}
	})
}
