package functions

import "github.com/go-libraries/govaluate/eval/functions/utils"

func IfNull(args ...interface{}) (interface{}, error) {
	if len(args) > 0 {
		for _, v := range args {
			if v != nil {
				return v, nil
			}
		}
		return nil, nil
	}
	return nil, utils.WrongParamsCount
}
