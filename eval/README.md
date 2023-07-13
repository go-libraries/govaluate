govaluate
====

[![Build Status](https://travis-ci.org/Knetic/govaluate.svg?branch=master)](https://travis-ci.org/Knetic/govaluate)
[![Godoc](https://godoc.org/github.com/Knetic/govaluate?status.png)](https://godoc.org/github.com/Knetic/govaluate)

# 作用

此包是一个类似计算器功能的包，内部包含自定义语法解析

因为有些功能在原包体上不支持，因此进行了二次开发

example:
```go
type T1 struct {
    A  string
    FF []T1
    N  int
}

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
	expression, err := NewEvaluableExpressionWithFunctions("add(slice_len([obj.FF]),sum([obj.FF], 'N')) ", funcs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
```