package leetcode

import (
	"bufio"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"runtime/pprof"
	"testing"
)

// load reflect of func
func loadMethodInfo(valueOfFn reflect.Value) ([]reflect.Type, reflect.Type) {
	argsType := make([]reflect.Type, valueOfFn.Type().NumIn())
	for i := range valueOfFn.Type().NumIn() {
		argsType[i] = valueOfFn.Type().In(i)
	}
	retType := valueOfFn.Type().Out(0)

	return argsType, retType
}

func createByType(tp reflect.Type) any {
	switch tp.Kind() {
	case reflect.String:
		return ""
	case reflect.Int:
		return 0
	case reflect.Int64:
		return int64(0)
	case reflect.Float64:
		return 0.0
	case reflect.Slice:
		switch tp.Elem().Kind() {
		case reflect.String:
			return make([]string, 0)
		case reflect.Int:
			return make([]int, 0)
		case reflect.Int64:
			return make([]int64, 0)
		case reflect.Float64:
			return make([]float64, 0)
		case reflect.Slice:
			switch tp.Elem().Elem().Kind() {
			case reflect.String:
				return make([][]string, 0)
			case reflect.Int:
				return make([][]int, 0)
			case reflect.Int64:
				return make([][]int64, 0)
			case reflect.Float64:
				return make([][]float64, 0)
			default:
				panic("unhandled default case")
			}
		default:
			panic("unhandled default case")
		}
	default:
		panic("unhandled default case")
	}
}

// load data
func loadData(filename string, valueOfFn reflect.Value) ([][]reflect.Value, error) {
	data, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	args := make([][]reflect.Value, 0)

	argsType, _ := loadMethodInfo(valueOfFn)

	arg := make([]any, len(argsType))
	for i, argType := range argsType {
		arg[i] = createByType(argType)
	}

	scanner := bufio.NewScanner(data)
	for {
		for i := range arg {
			if !scanner.Scan() {
				return args, nil
			}

			text := scanner.Text()
			err := json.Unmarshal([]byte(text), &arg[i])
			if err != nil {
				return nil, err
			}
		}

		argValue := make([]reflect.Value, len(arg))
		for i, a := range arg {
			argValue[i] = reflect.ValueOf(a)
		}
		args = append(args, argValue)
	}
}

func loadResult(filename string, valueOfFn reflect.Value) ([]reflect.Value, error) {
	data, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	result := make([]reflect.Value, 0)

	_, resultType := loadMethodInfo(valueOfFn)
	value := createByType(resultType)

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		text := scanner.Text()
		err := json.Unmarshal([]byte(text), &value)
		if err != nil {
			return nil, err
		}
		result = append(result, reflect.ValueOf(value))
	}

	return result, nil
}

func TestTable(t *testing.T) {
	if fn == nil {
		return
	}

	// load test data
	valueOfFn := reflect.ValueOf(fn)
	assert.Equal(t, reflect.Func, valueOfFn.Kind())

	args, err := loadData("data.txt", valueOfFn)
	assert.Nil(t, err)

	result, err := loadResult("result.txt", valueOfFn)
	assert.Nil(t, err)
	assert.Equal(t, len(args), len(result))

	// start cpu profile
	file, _ := os.CreateTemp("", "cpu.prof")
	println("cpu.prof:", file.Name())
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_ = pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	// run tests
	for i, arg := range args {
		values := valueOfFn.Call(arg)
		assert.Equal(t, 1, len(values))
		assert.True(t, result[i].Equal(values[0]))
	}
}

var fn any = nil
