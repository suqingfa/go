package leetcode

import (
	"github.com/stretchr/testify/assert"
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
			assert.True(t, test.output == 0)
		})
	}
}
