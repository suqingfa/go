package leetcode

import (
	"os"
	"runtime/pprof"
	"testing"
)

func TestTable(t *testing.T) {
	file, _ := os.CreateTemp(os.TempDir(), "cpu.prof")
	println("cpu.prof:", file.Name())
	defer file.Close()
	pprof.StartCPUProfile(file)
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
}
