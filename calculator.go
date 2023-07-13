package govaluate

import (
	"errors"
	"github.com/go-libraries/govaluate/defined"
	"github.com/go-libraries/govaluate/eval"
	"github.com/go-libraries/govaluate/input"
	"github.com/go-libraries/govaluate/output"
	"github.com/modern-go/reflect2"
	"reflect"
)

type Calculator struct {
	Eval                defined.IExpression
	CalculatorRelations []CalculatorRelation
}

type CalculatorRelation struct {
	OutputField string `json:"output_field"`
	Expression  string `json:"expression"`
} // output_field = eval(expression)

func NewCalculator() *Calculator {
	return &Calculator{
		Eval: eval.GetEvalWithFunctions(),
	}
}

func (c *Calculator) SetCalculatorRelations(values []CalculatorRelation) {
	c.CalculatorRelations = values
}

func (c *Calculator) Expression(in interface{}, out interface{}) error {
	rt := reflect2.TypeOf(out)
	if rt.Kind() != reflect.Pointer {
		return errors.New("返回值接受值必须设置为引用类型")
	}
	result := output.Result{Obj: out}
	req := &input.EvalParameters{Obj: in}
	for _, v := range c.CalculatorRelations {
		res, err := c.Eval.DoExpress(v.Expression, req)
		if err != nil {
			return err
		}
		err = result.Set(v.OutputField, res)
		if err != nil {
			return err
		}
	}
	return nil
}
