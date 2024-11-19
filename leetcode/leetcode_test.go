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

func createByType(tp reflect.Type, text []byte) (any, error) {
	switch tp.Kind() {
	case reflect.String:
		value := ""
		err := json.Unmarshal(text, &value)
		return value, err
	case reflect.Int:
		value := 0
		err := json.Unmarshal(text, &value)
		return value, err
	case reflect.Int64:
		value := int64(0)
		err := json.Unmarshal(text, &value)
		return value, err
	case reflect.Float64:
		value := 0.0
		err := json.Unmarshal(text, &value)
		return value, err
	case reflect.Bool:
		value := false
		err := json.Unmarshal(text, &value)
		return value, err
	case reflect.Slice:
		switch tp.Elem().Kind() {
		case reflect.String:
			value := make([]string, 0)
			err := json.Unmarshal(text, &value)
			return value, err
		case reflect.Int:
			value := make([]int, 0)
			err := json.Unmarshal(text, &value)
			return value, err
		case reflect.Int64:
			value := make([]int64, 0)
			err := json.Unmarshal(text, &value)
			return value, err
		case reflect.Float64:
			value := make([]float64, 0)
			err := json.Unmarshal(text, &value)
			return value, err
		case reflect.Bool:
			value := make([]bool, 0)
			err := json.Unmarshal(text, &value)
			return value, err
		case reflect.Slice:
			switch tp.Elem().Elem().Kind() {
			case reflect.String:
				value := make([][]string, 0)
				err := json.Unmarshal(text, &value)
				return value, err
			case reflect.Int:
				value := make([][]int, 0)
				err := json.Unmarshal(text, &value)
				return value, err
			case reflect.Int64:
				value := make([][]int64, 0)
				err := json.Unmarshal(text, &value)
				return value, err
			case reflect.Float64:
				value := make([][]float64, 0)
				err := json.Unmarshal(text, &value)
				return value, err
			case reflect.Bool:
				value := make([][]bool, 0)
				err := json.Unmarshal(text, &value)
				return value, err
			default:
				panic("unhandled default case " + tp.String())
			}
		default:
			panic("unhandled default case " + tp.String())
		}
	default:
		panic("unhandled default case " + tp.String())
	}
}

func TestCreateByType(t *testing.T) {
	tests := []struct {
		name   string
		tp     reflect.Type
		text   string
		result any
	}{
		{"", reflect.TypeFor[string](), `"abc"`, "abc"},
		{"", reflect.TypeFor[int](), `123`, 123},
		{"", reflect.TypeFor[int64](), `123`, int64(123)},
		{"", reflect.TypeFor[float64](), `123.4`, 123.4},
		{"", reflect.TypeFor[bool](), `true`, true},
		{"", reflect.TypeFor[bool](), `false`, false},

		{"", reflect.TypeFor[[]string](), `["abc","bcd"]`, []string{"abc", "bcd"}},
		{"", reflect.TypeFor[[]int](), `[123,234]`, []int{123, 234}},
		{"", reflect.TypeFor[[]int64](), `[123,234]`, []int64{int64(123), int64(234)}},
		{"", reflect.TypeFor[[]float64](), `[123.4,234.5]`, []float64{123.4, 234.5}},
		{"", reflect.TypeFor[[]bool](), `[true,false]`, []bool{true, false}},

		{"", reflect.TypeFor[[][]string](), `[["abc","bcd"]]`, [][]string{{"abc", "bcd"}}},
		{"", reflect.TypeFor[[][]int](), `[[123,234]]`, [][]int{{123, 234}}},
		{"", reflect.TypeFor[[][]int64](), `[[123,234]]`, [][]int64{{int64(123), int64(234)}}},
		{"", reflect.TypeFor[[][]float64](), `[[123.4,234.5]]`, [][]float64{{123.4, 234.5}}},
		{"", reflect.TypeFor[[][]bool](), `[[true,false]]`, [][]bool{{true, false}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value, err := createByType(test.tp, []byte(test.text))
			assert.NoError(t, err)
			assert.Equal(t, test.result, value)
		})
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
	buffer := make([]byte, 1024*1024*100)
	scanner.Buffer(buffer, len(buffer))
	for {
		arg := make([]any, len(argsType))

		for i := range arg {
			if !scanner.Scan() {
				return args, nil
			}

			value, err := createByType(argsType[i], scanner.Bytes())
			if err != nil {
				return nil, err
			}
			arg[i] = value
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
	buffer := make([]byte, 1024*1024*100)
	scanner.Buffer(buffer, len(buffer))
	for scanner.Scan() {
		value, err := createByType(resultType, scanner.Bytes())
		if err != nil {
			return nil, err
		}
		result = append(result, reflect.ValueOf(value))
	}

	return result, nil
}

func TestFn(t *testing.T) {
	if fn == nil {
		return
	}

	// load test data
	valueOfFn := reflect.ValueOf(fn)
	assert.Equal(t, reflect.Func, valueOfFn.Kind())

	args, err := loadData("data.txt", valueOfFn)
	assert.NoError(t, err)

	result, err := loadResult("result.txt", valueOfFn)
	assert.NoError(t, err)
	assert.Equal(t, len(args), len(result))

	// start cpu profile
	file, _ := os.CreateTemp("", "cpu.prof")
	t.Log("cpu profile:", file.Name())
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_ = pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	// run tests
	for i, arg := range args {
		t.Run("", func(t *testing.T) {
			values := valueOfFn.Call(arg)
			assert.Equal(t, 1, len(values))
			assert.Equal(t, result[i].Interface(), values[0].Interface())
		})
	}
}
