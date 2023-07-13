package mathx

import (
	"github.com/go-libraries/govaluate/eval/functions/utils"
)

func Max(args ...interface{}) (interface{}, error) {
	if len(args) > 0 {
		var floatMax float64
		for idx, arg := range args {
			v, ok := utils.ToFloat64(arg)
			if !ok {
				return nil, utils.NewWrongParamType(idx)
			}
			if idx == 0 {
				floatMax = v
			} else if v > floatMax {
				floatMax = v
			}
		}
		return floatMax, nil
	}
	return nil, utils.WrongParamsCount
}

func Min(args ...interface{}) (interface{}, error) {
	if len(args) > 0 {
		var floatMax float64
		for idx, arg := range args {
			v, ok := utils.ToFloat64(arg)
			if !ok {
				return nil, utils.NewWrongParamType(idx)
			}
			if idx == 0 {
				floatMax = v
			} else if v < floatMax {
				floatMax = v
			}
		}
		return floatMax, nil
	}
	return nil, utils.WrongParamsCount
}
