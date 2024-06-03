package leetcode

import (
	"bufio"
	"encoding/json"
	"fmt"
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

func createByType(tp reflect.Type, text string) any {
	switch tp.Kind() {
	case reflect.String:
		return text[1 : len(text)-1]
	case reflect.Int:
		value := 0
		_, _ = fmt.Sscan(text, &value)
		return value
	case reflect.Int64:
		value := int64(0)
		_, _ = fmt.Sscan(text, &value)
		return value
	case reflect.Float64:
		value := 0.0
		_, _ = fmt.Sscan(text, &value)
		return value
	case reflect.Bool:
		value := true
		_, _ = fmt.Sscan(text, &value)
		return value
	case reflect.Slice:
		switch tp.Elem().Kind() {
		case reflect.String:
			value := make([]string, 0)
			_ = json.Unmarshal([]byte(text), &value)
			return value
		case reflect.Int:
			value := make([]int, 0)
			_ = json.Unmarshal([]byte(text), &value)
			return value
		case reflect.Int64:
			value := make([]int64, 0)
			_ = json.Unmarshal([]byte(text), &value)
			return value
		case reflect.Float64:
			value := make([]float64, 0)
			_ = json.Unmarshal([]byte(text), &value)
			return value
		case reflect.Bool:
			value := make([]bool, 0)
			_ = json.Unmarshal([]byte(text), &value)
			return value
		case reflect.Slice:
			switch tp.Elem().Elem().Kind() {
			case reflect.String:
				value := make([][]string, 0)
				_ = json.Unmarshal([]byte(text), &value)
				return value
			case reflect.Int:
				value := make([][]int, 0)
				_ = json.Unmarshal([]byte(text), &value)
				return value
			case reflect.Int64:
				value := make([][]int64, 0)
				_ = json.Unmarshal([]byte(text), &value)
				return value
			case reflect.Float64:
				value := make([][]float64, 0)
				_ = json.Unmarshal([]byte(text), &value)
				return value
			case reflect.Bool:
				value := make([][]bool, 0)
				_ = json.Unmarshal([]byte(text), &value)
				return value
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

	scanner := bufio.NewScanner(data)
	for {
		arg := make([]any, len(argsType))

		for i := range arg {
			if !scanner.Scan() {
				return args, nil
			}

			arg[i] = createByType(argsType[i], scanner.Text())
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

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		result = append(result, reflect.ValueOf(createByType(resultType, scanner.Text())))
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
		assert.Equal(t, result[i].Interface(), values[0].Interface())
	}
}

var fn any = nil
