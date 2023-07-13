package mathx

import (
	"github.com/go-libraries/govaluate/eval/functions/utils"
	"math"
)

func Sin(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Sin(v), nil
	}
	return nil, utils.WrongParamsCount
}

func Sinh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Sinh(v), nil
	}
	return nil, utils.WrongParamsCount
}

func Asin(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Asin(v), nil
	}
	return nil, utils.WrongParamsCount
}

func Asinh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Asinh(v), nil
	}
	return nil, utils.WrongParamsCount
}
