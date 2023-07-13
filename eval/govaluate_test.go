package eval

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
	"testing"
)

type T1 struct {
	A  string
	FF []T1
	N  int
}

func TestFunctions(t *testing.T) {
	//expression, err := NewEvaluableExpression("first_name == 'Julian' && emp_id == 302 && birth_date >= '2016-12-12' && birth_date <= '2016-12-25'")
	//fmt.Println(int64(time.Second / (4195 * time.Nanosecond)))
	//return
	parameters := make(map[string]interface{}, 8)
	parameters["first_name"] = "Julian"
	parameters["emp_id"] = 302
	parameters["birth_date"] = "2016-12-15"
	//obj := T1{
	//	A: "test",
	//	FF: []T1{
	//		{A: "fff", N: 1},
	//		{A: "BBB", N: 2},
	//	},
	//}
	//parameters["obj"]
	//bts, _ := json.Marshal(obj)
	//var ff map[string]interface{}
	//json.Unmarshal(bts, &ff)
	//parameters["obj"] = obj
	funcs := map[string]ExpressionFunction{
		"add": func(arguments ...interface{}) (interface{}, error) {
			return cast.ToInt64(arguments[0]) + cast.ToInt64(arguments[1]), nil
		},
		"strlen": func(arguments ...interface{}) (interface{}, error) {
			return len(arguments[0].(string)), nil
		},
		"slice_len": func(arguments ...interface{}) (interface{}, error) {
			return reflect.ValueOf(arguments[0]).Len(), nil
		},
		"sum": func(arguments ...interface{}) (interface{}, error) {
			vt := reflect.ValueOf(arguments[0])
			var res int64
			var field string
			if len(arguments) > 1 {
				field = arguments[1].(string)
			}
			switch vt.Kind() {
			case reflect.Slice:
				for i := 0; i < vt.Len(); i++ {
					if field != "" {
						vf := vt.Index(i).FieldByName(field)
						if vf.CanInt() {
							res += vf.Int()
						}
					} else {
						vf := vt.Index(i)
						if vf.CanInt() {
							res += vf.Int()
						}
					}
				}
			}
			return res, nil
		},
	}
	expression, err := NewEvaluableExpressionWithFunctions("add(slice_len([obj.FF]),sum([obj.FF], 'N')) ", funcs)
	if err != nil {
		fmt.Println("parse", err.Error())
		return
	}
	result, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

func BenchmarkNew(b *testing.B) {
	parameters := make(map[string]interface{}, 8)
	parameters["first_name"] = "Julian"
	parameters["emp_id"] = 302
	parameters["birth_date"] = "2016-12-15"
	obj := T1{
		A: "test",
		FF: []T1{
			{A: "fff", N: 1},
			{A: "BBB", N: 2},
		},
	}
	//parameters["obj"]
	//bts, _ := json.Marshal(obj)
	//var ff map[string]interface{}
	//json.Unmarshal(bts, &ff)
	parameters["obj"] = obj
	funcs := map[string]ExpressionFunction{
		"add": func(arguments ...interface{}) (interface{}, error) {
			return cast.ToInt64(arguments[0]) + cast.ToInt64(arguments[1]), nil
		},
		"strlen": func(arguments ...interface{}) (interface{}, error) {
			return len(arguments[0].(string)), nil
		},
		"slice_len": func(arguments ...interface{}) (interface{}, error) {
			return reflect.ValueOf(arguments[0]).Len(), nil
		},
		"sum": func(arguments ...interface{}) (interface{}, error) {
			vt := reflect.ValueOf(arguments[0])
			var res int64
			var field string
			if len(arguments) > 1 {
				field = arguments[1].(string)
			}
			switch vt.Kind() {
			case reflect.Slice:
				for i := 0; i < vt.Len(); i++ {
					if field != "" {
						vf := vt.Index(i).FieldByName(field)
						if vf.CanInt() {
							res += vf.Int()
						}
					} else {
						vf := vt.Index(i)
						if vf.CanInt() {
							res += vf.Int()
						}
					}
				}
			}
			return res, nil
		},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expression, err := NewEvaluableExpressionWithFunctions("add(slice_len([obj.FF]),sum([obj.FF], 'N')) ", funcs)
		if err != nil {
			b.Errorf("%s", err.Error())
			return
		}
		_, err = expression.Evaluate(parameters)
		if err != nil {
			b.Errorf("%s", err.Error())
		}
		//b.Logf("%+v", result)
	}
}
