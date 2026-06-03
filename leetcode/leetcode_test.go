package leetcode

import (
	"bufio"
	"encoding/json"
	"math"
	"os"
	"reflect"
	"runtime/pprof"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func createByTypeValue(tp reflect.Type, v any) any {
	value := reflect.New(tp)
	switch tp.Kind() {
	case reflect.Bool:
		value.Elem().SetBool(v.(bool))
	case reflect.Uint8:
		value.Elem().SetUint(uint64(v.(string)[0]))
	case reflect.Int:
		value.Elem().SetInt(int64(v.(float64)))
	case reflect.Int64:
		value.Elem().SetInt(int64(v.(float64)))
	case reflect.Float32:
		value.Elem().SetFloat(float64(v.(float32)))
	case reflect.Float64:
		value.Elem().SetFloat(v.(float64))
	case reflect.String:
		value.Elem().SetString(v.(string))
	case reflect.Slice:
		v := v.([]any)
		value = reflect.MakeSlice(tp, len(v), len(v))
		for i, a := range v {
			e := createByTypeValue(tp.Elem(), a)
			value.Index(i).Set(reflect.ValueOf(e))
		}
		return value.Interface()
	default:
		panic("unhandled default case " + tp.String())
	}

	return value.Elem().Interface()
}

func createByType(tp reflect.Type, text []byte) (any, error) {
	if tp == reflect.PointerTo(reflect.TypeFor[TreeNode]()) {
		var ints []*int
		err := json.Unmarshal(text, &ints)
		if err != nil {
			return nil, err
		}

		if len(ints) == 0 {
			return nil, nil
		}

		root := &TreeNode{Val: *ints[0]}

		q := []*TreeNode{root}
		for i := 1; len(q) > 0 && i < len(ints); i += 2 {
			cur := q[0]
			q = q[1:]

			if ints[i] != nil {
				cur.Left = &TreeNode{Val: *ints[i]}
				q = append(q, cur.Left)
			}

			if ints[i+1] != nil {
				cur.Right = &TreeNode{Val: *ints[i+1]}
				q = append(q, cur.Right)
			}
		}

		return root, nil
	}

	var v any
	err := json.Unmarshal(text, &v)
	if err != nil {
		return nil, err
	}

	return createByTypeValue(tp, v), nil
}

func TestCreateByType(t *testing.T) {
	tests := []struct {
		name   string
		tp     reflect.Type
		text   string
		result any
	}{
		{"string", reflect.TypeFor[string](), `"abc"`, "abc"},
		{"byte", reflect.TypeFor[byte](), `"a"`, byte('a')},
		{"int", reflect.TypeFor[int](), `123`, 123},
		{"int64", reflect.TypeFor[int64](), `123`, int64(123)},
		{"float64", reflect.TypeFor[float64](), `123.4`, 123.4},
		{"bool", reflect.TypeFor[bool](), `true`, true},
		{"bool", reflect.TypeFor[bool](), `false`, false},

		{"[]string", reflect.TypeFor[[]string](), `["abc","bcd"]`, []string{"abc", "bcd"}},
		{"[]byte", reflect.TypeFor[[]byte](), `["a","b"]`, []byte{'a', 'b'}},
		{"[]int", reflect.TypeFor[[]int](), `[123,234]`, []int{123, 234}},
		{"[]int64", reflect.TypeFor[[]int64](), `[123,234]`, []int64{int64(123), int64(234)}},
		{"[]float64", reflect.TypeFor[[]float64](), `[123.4,234.5]`, []float64{123.4, 234.5}},
		{"[]bool", reflect.TypeFor[[]bool](), `[true,false]`, []bool{true, false}},

		{"[][]string", reflect.TypeFor[[][]string](), `[["abc","bcd"]]`, [][]string{{"abc", "bcd"}}},
		{"[][]byte", reflect.TypeFor[[][]byte](), `[["a","b"]]`, [][]byte{{'a', 'b'}}},
		{"[][]int", reflect.TypeFor[[][]int](), `[[123,234]]`, [][]int{{123, 234}}},
		{"[][]int64", reflect.TypeFor[[][]int64](), `[[123,234]]`, [][]int64{{int64(123), int64(234)}}},
		{"[][]float64", reflect.TypeFor[[][]float64](), `[[123.4,234.5]]`, [][]float64{{123.4, 234.5}}},
		{"[][]bool", reflect.TypeFor[[][]bool](), `[[true,false]]`, [][]bool{{true, false}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value, err := createByType(test.tp, []byte(test.text))
			assert.NoError(t, err)
			if !equal(reflect.ValueOf(test.result), reflect.ValueOf(value)) {
				assert.Equal(t, test.result, value)
			}
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

// equal test func
func equal(a, b reflect.Value) bool {
	if a.Kind() != b.Kind() {
		return false
	}

	switch a.Kind() {
	case reflect.Float64:
		return math.Abs(a.Float()-b.Float()) < 1e-5
	case reflect.Array | reflect.Slice:
		if a.Len() != b.Len() {
			return false
		}

		for i := 0; i < a.Len(); i++ {
			if !equal(a.Index(i), b.Index(i)) {
				return false
			}
		}
	default:
		return a.Interface() == b.Interface()
	}

	return true
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

	// 不是调试时超时退出
	debug := false
	statusFile, _ := os.Open("/proc/self/status")
	scanner := bufio.NewScanner(statusFile)
	for scanner.Scan() {
		text := scanner.Text()
		if !strings.HasPrefix(text, "TracerPid") {
			continue
		}

		split := strings.Split(text, ":")
		ppid, _ := strconv.Atoi(split[1][1:])
		debug = ppid > 0
		break
	}

	// start cpu profile
	file, _ := os.CreateTemp("", "cpu.prof")
	t.Log("cpu profile:", file.Name())
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_ = pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	if !debug {
		go func(c <-chan time.Time) {
			<-c
			pprof.StopCPUProfile()
			_ = file.Close()
			os.Exit(0)
		}(time.After(3 * time.Second))
	}

	// run tests
	for i, arg := range args {
		t.Run("", func(t *testing.T) {
			values := valueOfFn.Call(arg)
			assert.Equal(t, 1, len(values))
			if !equal(result[i], values[0]) {
				assert.Equal(t, result[i].Interface(), values[0].Interface())
			}
		})
	}
}
