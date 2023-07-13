package mathx

import (
	"github.com/go-libraries/govaluate/eval/functions/utils"
	"math"
)

func Exp(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Exp(v), nil
	}
	return nil, utils.WrongParamsCount
}

func Exp2(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Exp2(v), nil
	}
	return nil, utils.WrongParamsCount
}
