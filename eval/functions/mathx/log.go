package mathx

import (
	"github.com/go-libraries/govaluate/eval/functions/utils"
	"math"
)

func Log(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Log(v), nil
	}
	return nil, utils.WrongParamsCount
}

func Log10(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Log10(v), nil
	}
	return nil, utils.WrongParamsCount
}
