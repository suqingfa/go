package test

import (
	"math"
	"testing"
)

func TestIEEE754(t *testing.T) {
	// float32 1 8 23
	list := []uint32{
		0x00000000, // 0
		0x00000001, // 最小非规范数
		0x007fffff, // 最大非规范数
		0x00800000, // 最小规范数
		0x3f800000, // 1.0
		0x40000000, // 2.0
		0x7f7fffff, // 最大规范数
		0x7f800000, // +Inf
		0x7f800001, // NaN
		0x7f800002, // NaN
	}

	for _, i := range list {
		f := math.Float32frombits(i)
		t.Logf("%032b %08x %e\n", i, i, f)
	}
}
