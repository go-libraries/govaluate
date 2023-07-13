package mathx

import (
	"github.com/go-libraries/govaluate/eval/functions/utils"
	"math"
)

func Tan(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Tan(v), nil
	}
	return nil, utils.WrongParamsCount
}

func Tanh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Tanh(v), nil
	}
	return nil, utils.WrongParamsCount
}

func Atan(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Atan(v), nil
	}
	return nil, utils.WrongParamsCount
}

func Atan2(args ...interface{}) (interface{}, error) {
	if len(args) == 2 {
		v1, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		v2, ok := utils.ToFloat64(args[1])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(1)
		}
		return math.Atan2(v1, v2), nil
	}
	return nil, utils.WrongParamsCount
}

func Atanh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Atanh(v), nil
	}
	return nil, utils.WrongParamsCount
}
