package output

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Result
// 抽象出返回值，只需要填充到Obj对象，一级的反射对象会产生最多，因此优化只生成一次
type Result struct {
	Obj   interface{}
	value reflect.Value
	flag  bool
}

func (p *Result) Set(name string, value interface{}) error {
	if !p.flag {
		rv := reflect.ValueOf(p.Obj)
		if rv.Kind() != reflect.Pointer || rv.IsNil() {
			return errors.New("设置对象必须是引用值")
		}
		p.flag = true
		p.value = rv
	}
	if name == "" {
		rv := reflect.ValueOf(value)
		p.value.Elem().Set(rv)
		return nil
	}
	var arr []string
	if strings.Contains(name, ".") {
		arr = strings.Split(name, ".")
		name = arr[0]
	} else {
		arr = make([]string, 1)
		arr[0] = name
	}

	return ReflectSetValue(p.value, value, "", arr...)
}

func ReflectSetValue(val reflect.Value, setVal interface{}, preKey string, keys ...string) (err error) {
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
		return errors.New("暂时不支持对数组进行赋值")

	case reflect.Invalid:

	}
	if !vt.IsValid() || vt.Kind() == reflect.Invalid {
		return fmt.Errorf("不存在查询值 %s.%s ", preKey, k)
	}
	//if vt.Kind() == reflect.Pointer {
	//	if vt.IsNil() {
	//		vt.Set(reflect.New(vt.Type().Elem()))
	//	}
	//	vt = vt.Elem()
	//}
	vt = decAlloc(vt)
	if len(keys) == 1 {
		defer func() {
			x := recover()
			if x != nil {
				err = fmt.Errorf("设置值失败 %s.%s 期望type:%s 赋值value的type: %s 实际value:%+v",
					preKey, k, vt.Type().String(), reflect.TypeOf(setVal).String(), setVal)
			}
		}()
		vt.Set(reflect.ValueOf(setVal))
		return nil
	}
	return ReflectSetValue(vt, setVal, keys[0], keys[1:]...)
}

func decAlloc(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		v = v.Elem()
	}
	return v
}
