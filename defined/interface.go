package defined

type IInput interface {
	Get(name string) (interface{}, error)
}

type IOutput interface {
	Set(name string, val interface{}) error
}

type IExpression interface {
	CheckExpressionValid(exp string) bool
	DoExpress(exp string, input IInput) (result interface{}, err error)
}
