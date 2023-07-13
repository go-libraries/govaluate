package input

import (
	"fmt"
	"reflect"
	"strings"
)

// EvalParameters
// 抽象出输入值，只需要填充到Obj对象，一级的反射对象会产生最多，因此优化只生成一次
type EvalParameters struct {
	Obj   interface{}
	value reflect.Value
	flag  bool
}

func (p *EvalParameters) Get(name string) (interface{}, error) {
	var arr []string
	if strings.Contains(name, ".") {
		arr = strings.Split(name, ".")
		name = arr[0]
	} else {
		arr = make([]string, 1)
		arr[0] = name
	}
	if !p.flag {
		p.flag = true
		p.value = reflect.ValueOf(p.Obj)
	}
	return ReflectGetValue(p.value, "", arr...)
}

func ReflectGetValue(val reflect.Value, preKey string, keys ...string) (interface{}, error) {
	k := keys[0]

	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}
	val = reflect.Indirect(val)
	var vt reflect.Value
	switch val.Kind() {
	case reflect.Struct:
		vt = val.FieldByName(k)
	case reflect.Map:
		vt = val.MapIndex(reflect.ValueOf(k))
	case reflect.Array, reflect.Slice:
		l := val.Len()
		index := GetDigitValue(k)
		i := index
		if index == -1 {
			i = 0
		}
		if l-1 < i {
			return nil, fmt.Errorf(" 数组 %s.%s 超出最大下标%d", preKey, k, l-1)
		}
		if index == -1 {
			vt = val.Index(0)
			return ReflectGetValue(vt, preKey, keys...)
		}
		vt = val.Index(index)
	case reflect.Invalid:

	}
	if !vt.IsValid() || !valid(vt) {
		return nil, fmt.Errorf("不存在查询值 %s.%s ", preKey, k)
	}
	if len(keys) == 1 {
		return vt.Interface(), nil
	}
	return ReflectGetValue(vt, keys[0], keys[1:]...)
}

func IsSingleDigit(data byte) int {
	digit := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i, item := range digit {
		if data == item {
			return i
		}
	}
	return -1
}

func GetDigitValue(data string) int {
	res := 0
	for i := 0; i < len(data); i++ {
		v := IsSingleDigit(data[i])
		if v == -1 {
			return -1
		}
		res = res*10 + v
	}
	return res
}

func valid(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Invalid:
		return false
	case reflect.Pointer:
		return !v.IsNil()
	}
	return true
}
