package mathx

import (
	"github.com/go-libraries/govaluate/eval/functions/utils"
	"math"
)

func Cos(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Cos(v), nil
	}
	return nil, utils.WrongParamsCount
}

func Cosh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Cosh(v), nil
	}
	return nil, utils.WrongParamsCount
}

func Acos(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Acos(v), nil
	}
	return nil, utils.WrongParamsCount
}

func Acosh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Acosh(v), nil
	}
	return nil, utils.WrongParamsCount
}
