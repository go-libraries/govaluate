package eval

import (
	"github.com/go-libraries/govaluate/eval/functions"
	"github.com/go-libraries/govaluate/eval/functions/mathx"
)

var FlowFunctions = map[string]ExpressionFunction{
	"isDefined": functions.IsDefined,
	"ifnull":    functions.IfNull,
	"if":        functions.If,
}

var MathFunctions = map[string]ExpressionFunction{
	"cos":   mathx.Cos,
	"acos":  mathx.Acos,
	"cosh":  mathx.Cosh,
	"acosh": mathx.Acosh,

	"exp":  mathx.Exp,
	"exp2": mathx.Exp2,

	"log":   mathx.Log,
	"log10": mathx.Log10,

	"round": mathx.Round,
	"floor": mathx.Floor,
	"ceil":  mathx.Ceil,
	"trunc": mathx.Trunc,

	"sin":   mathx.Sin,
	"asin":  mathx.Asin,
	"sinh":  mathx.Sinh,
	"asinh": mathx.Asinh,

	"sqrt": mathx.Sqrt,

	"tan":   mathx.Tan,
	"atan":  mathx.Atan,
	"atan2": mathx.Atan2,
	"tanh":  mathx.Tanh,
	"atanh": mathx.Atanh,

	"max": mathx.Max,
	"min": mathx.Min,
}

var Functions map[string]ExpressionFunction

func init() {
	Functions = make(map[string]ExpressionFunction)
	for k, v := range MathFunctions {
		Functions[k] = v
	}
	for k, v := range FlowFunctions {
		Functions[k] = v
	}
}
