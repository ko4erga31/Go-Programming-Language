package popcount

import (
	"math/bits"
	"testing"
)

var sink int

func TestPopCountV2(t *testing.T) {
	for _, x := range []uint64{0, 1, 0xFF, 0x0102040810204080, ^uint64(0)} {
		if got, want := PopCountV2(x), bits.OnesCount64(x); got != want {
			t.Errorf("PopCountV2(%#x) = %d, want %d", x, got, want)
		}
	}
}

func BenchmarkPopCountV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = PopCountV1(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = PopCountV2(0x1234567890ABCDEF)
	}
}

func BenchmarkOnesCount64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = bits.OnesCount64(0x1234567890ABCDEF)
	}
}
