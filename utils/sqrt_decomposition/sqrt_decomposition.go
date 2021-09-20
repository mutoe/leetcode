package sqrt_decomposition

import "math"

type value int

type SqrtDecomposition struct {
	data   []value
	N      int
	blocks []value
	B      int
	Bn     int
}

func Constructor(data []value) SqrtDecomposition {
	sqrt := math.Sqrt(float64(len(data)))
	blockSize := int(sqrt)
	blockCount := int(math.Ceil(sqrt))

	block := make([]value, blockCount)
	for i := 0; i < len(data); i++ {
		block[i/blockSize] += data[i]
	}

	return SqrtDecomposition{
		data:   data,
		N:      len(data),
		blocks: block,
		B:      blockSize,
		Bn:     blockCount,
	}
}

func (sqrt SqrtDecomposition) sumRange(a, b int) value {
	bs, be := a/sqrt.B, b/sqrt.B
	res := value(0)
	if bs == be {
		for i := a; i <= b; i++ {
			res += sqrt.data[i]
		}
		return res
	}
	for i := a; i < sqrt.B*(bs+1); i++ {
		res += sqrt.data[i]
	}
	for b := bs + 1; b < be; b++ {
		res += sqrt.blocks[b]
	}
	for i := be * sqrt.B; i <= b; i++ {
		res += sqrt.data[i]
	}
	return res
}

func (sqrt *SqrtDecomposition) update(i int, value value) {
	oldValue := sqrt.data[i]
	sqrt.data[i] = value
	sqrt.blocks[i/sqrt.B] += value - oldValue
}
