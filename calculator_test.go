package govaluate

import (
	"fmt"
	"testing"
)

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
	t.Logf("%+v", e)
	fmt.Println(res)
}

type MrpDict struct {
	Basic struct {
		Level  string
		Status string
	}
}

func TestNewCalculatorBool(t *testing.T) {
	cal := NewCalculator()
	cal.SetCalculatorRelations([]CalculatorRelation{
		{
			OutputField: "",
			Expression:  "[Basic.Level] == 'Top1' && [Basic.Status] == 'New'",
		},
	})
	res := false
	e := cal.Expression(MrpDict{
		Basic: struct {
			Level  string
			Status string
		}{Level: "Top1", Status: "New"},
	}, &res)
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
		{
			OutputField: "StringValue",
			Expression:  "'c'",
		},
	})
	res := CalculatorOutPut{}

	for i := 0; i < b.N; i++ {
		//3次运算 赋值
		_ = cal.Expression(CalculatorInPut{}, &res)
	}
	b.ReportAllocs()
}
