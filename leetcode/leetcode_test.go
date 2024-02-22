package leetcode

import (
	"os"
	"runtime/pprof"
	"testing"
)

func TestTable(t *testing.T) {
	file, _ := os.CreateTemp("", "cpu.prof")
	println("cpu.prof:", file.Name())
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_ = pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	tests := []struct {
		name   string
		output int
		input  int
	}{
		{},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
		})
	}

	const N int = 1e4
	nums := make([]int, N)
	for i := range N {
		nums[i] = 1
	}

	matrix := make([][]int, N)
	for i := range N {
		matrix[i] = make([]int, N)
	}
}
