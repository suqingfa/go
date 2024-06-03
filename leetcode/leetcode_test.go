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
		arg := make([]reflect.Value, len(argsType))
		for i, argType := range argsType {
			if !scanner.Scan() {
				return args, nil
			}

			text := scanner.Text()
			var value any
			switch argType.Kind() {
			case reflect.String:
				value = text[1 : len(text)-1]
			case reflect.Int:
				t := 0
				_, err := fmt.Sscan(text, &t)
				if err != nil {
					return nil, err
				}
				value = t
			case reflect.Float64:
				t := 0.0
				_, err := fmt.Sscan(text, &t)
				if err != nil {
					return nil, err
				}
				value = t
			case reflect.Slice:
				switch argType.Elem().Kind() {
				case reflect.String:
					value = make([]string, 0)
				case reflect.Int:
					value = make([]int, 0)
				case reflect.Float64:
					value = make([]float64, 0)
				case reflect.Slice:
					switch argType.Elem().Elem().Kind() {
					case reflect.String:
						value = make([][]string, 0)
					case reflect.Int:
						value = make([][]int, 0)
					case reflect.Float64:
						value = make([][]float64, 0)
					default:
						panic("unhandled default case")
					}
				default:
					panic("unhandled default case")
				}
				err := json.Unmarshal([]byte(text), &value)
				if err != nil {
					return nil, err
				}
			default:
				panic("unhandled default case")
			}
			arg[i] = reflect.ValueOf(value)
		}
		args = append(args, arg)
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
		text := scanner.Text()
		var value any
		switch resultType.Kind() {
		case reflect.String:
			value = text[1 : len(text)-1]
		case reflect.Int:
			t := 0
			_, err := fmt.Sscan(text, &t)
			if err != nil {
				return nil, err
			}
			value = t
		case reflect.Float64:
			t := 0.0
			_, err := fmt.Sscan(text, &t)
			if err != nil {
				return nil, err
			}
			value = t
		default:
			panic("unhandled default case")
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
