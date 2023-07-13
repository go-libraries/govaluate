package eval

import (
	"github.com/go-libraries/govaluate/defined"
	"sync"
)

type Eval struct {
	functions map[string]ExpressionFunction
}

var evalObj *Eval
var evalOnce sync.Once

func GetEvalWithFunctions() *Eval {
	if evalObj != nil {
		return evalObj
	}
	evalOnce.Do(func() {
		evalObj = &Eval{functions: Functions}
	})
	return evalObj
}

func (e *Eval) CheckExpressionValid(exp string) bool {
	_, err := NewEvaluableExpressionWithFunctions(exp, e.functions)
	if err != nil {
		return false
	}
	return true
}

// 新增pool防止，同样成立的公式，不会重新构造
// 添加后  6 allocs/op   291.2 ns/op
// 添加前  84 allocs/op  4000+ ns/op
var evalSuccessPool map[string]*sync.Pool
var evalLock sync.Mutex

func init() {
	evalSuccessPool = make(map[string]*sync.Pool)
}

func (e *Eval) DoExpress(exp string, input defined.IInput) (result interface{}, err error) {
	var expression *EvaluableExpression
	if v, ok := evalSuccessPool[exp]; ok {
		expression = v.Get().(*EvaluableExpression)
	} else {
		//初始化的时候构建需要先检测是否可行
		expression, err = NewEvaluableExpressionWithFunctions(exp, e.functions)
		if err != nil {
			return nil, err
		}
		evalLock.Lock()
		defer evalLock.Unlock()
		// 并发第二个锁进来时 已经初始化过了 没必要再初始化
		if v, ok := evalSuccessPool[exp]; ok {
			expression = v.Get().(*EvaluableExpression)
		} else {
			evalSuccessPool[exp] = &sync.Pool{
				New: func() interface{} {
					e1, _ := NewEvaluableExpressionWithFunctions(exp, e.functions)
					return e1
				},
			}
		}
	}
	defer evalSuccessPool[exp].Put(expression)
	return expression.EvaluateWithInterface(input)
}
