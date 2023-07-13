govaluate
====

[![Build Status](https://travis-ci.org/Knetic/govaluate.svg?branch=master)](https://travis-ci.org/Knetic/govaluate)
[![Godoc](https://godoc.org/github.com/Knetic/govaluate?status.png)](https://godoc.org/github.com/Knetic/govaluate)

# 作用

此包是一个类似计算器功能的包，内部包含自定义语法解析

因为有些功能在原包体上不支持，因此进行了二次开发,并进行性能优化

example:
```go
type CalculatorOutPut struct {
    IntValue    float64 `json:"int_value"` // 如果用默认的运算符，返回值都会是float64类型  如果需要指定类型，就需要编写元函数
    FloatValue  float64 `json:"float_value"`
    StringValue string  `json:"string_value"`
}

type CalculatorInPut struct {
    IntValue    float64 `json:"int_value"`
    FloatValue  float64 `json:"float_value"`
    StringValue string  `json:"string_value"`
}

func TestNewCalculator(t *testing.T) {
    cal := NewCalculator()
    cal.SetCalculatorRelations([]CalculatorRelation{
        {
            OutputField: "IntValue",
            Expression:  "IntValue+1",
        },
        {
            OutputField: "FloatValue",
            Expression:  "FloatValue+3",
        },
        {
            OutputField: "StringValue",
            Expression:  "'a'",
        },
    })
    res := CalculatorOutPut{}
    e := cal.Expression(CalculatorInPut{}, &res)
    t.Logf("%+v %+v", e, res)
}

func BenchmarkNewCalculator(b *testing.B) {
    cal := NewCalculator()
    cal.SetCalculatorRelations([]CalculatorRelation{
        {
            OutputField: "IntValue",
            Expression:  "IntValue+1",
        },
        {
            OutputField: "FloatValue",
            Expression:  "FloatValue+3",
        },
        {
            OutputField: "StringValue",
            Expression:  "'a'",
        },
    })
    res := CalculatorOutPut{}

    for i := 0; i < b.N; i++ {
        _ = cal.Expression(CalculatorInPut{}, &res)
    }
    b.ReportAllocs()
}

```