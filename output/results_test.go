package output

import (
	"fmt"
	"testing"
)

type T1 struct {
	A  string
	FF *T1
}

func TestResult_Set(t *testing.T) {
	p := &Result{
		Obj: &T1{
			//FF: &T1{
			//	FF: &T1{},
			//},
		},
	}
	err := p.Set("FF.FF.A", "tt.args.value")
	fmt.Println(err)
	res := p.Obj.(*T1)
	fmt.Printf("%+v\n", p.Obj)
	fmt.Println("===")
	fmt.Printf("%+v\n", res.FF.FF.A)
	fmt.Println("===")
}

func BenchmarkReflectSetValue(b *testing.B) {
	p := &Result{
		Obj: &T1{},
	}
	for i := 0; i < b.N; i++ {
		_ = p.Set("FF.FF.A", "tt.args.value")
	}
	b.ReportAllocs()
}
