package mathx

import (
	"github.com/go-libraries/govaluate/eval/functions/utils"
	"math"
)

func Sqrt(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := utils.ToFloat64(args[0])
		if !ok {
			return math.NaN(), utils.NewWrongParamType(0)
		}
		return math.Sqrt(v), nil
	}
	return nil, utils.WrongParamsCount
}
