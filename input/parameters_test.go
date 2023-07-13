package input

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"reflect"
	"testing"
)

type T struct {
	A  string
	FF []T
}

func TestMapParameters_Get(t *testing.T) {
	parameters := make(map[string]interface{}, 8)
	parameters["first_name"] = "Julian"
	parameters["emp_id"] = 302
	parameters["birth_date"] = "2016-12-15"
	parameters["obj"] = T{
		A: "test",
		FF: []T{
			{
				A: "FFFFFF",
			},
		},
	}
	type TestObj struct {
		Obj interface{}
	}
	var t11 = TestObj{
		Obj: &T{
			A: "test",
			FF: []T{
				{
					A: "123F",
				},
			},
		},
	}
	//reflect2.TypeOf(t11).Indirect()
	fmt.Println(ReflectGetValue(reflect.ValueOf(t11), "", []string{"Obj", "FF", "A"}...))
	//rt := reflect2.TypeOf(parameters["obj"])
	//ptrStart := reflect2.PtrOf(parameters["obj"])
	//reflect.ValueOf(parameters)
	//ff, _ := rt.Type1().Elem().FieldByName("A")
	//fieldPtr := uintptr(ptrStart) + ff.Offset
	//fmt.Println(fieldPtr)
	//fmt.Printf("%+v\n", rt.Indirect(parameters["obj"]))
	//	fmt.Printf("%+v\n", rt.UnsafeIndirect(unsafe.Pointer(fieldPtr)))
	//rv := reflect.ValueOf(parameters["obj"])
	//fmt.Println(rv.FieldByName("A").String())

	//return
	v1 := EvalParameters{
		Obj: &T{
			A: "test",
			FF: []T{
				{
					A: "123F",
				},
			},
		},
	}
	////fmt.Println(v1.Get("obj.FF.A"))
	fmt.Println(v1.Get("FF.A"))
	reflect2.TypeOf(t11)

}

func BenchmarkReflect(b *testing.B) {
	parameters := EvalParameters{
		Obj: &T{
			A: "test",
			FF: []T{
				{
					A: "123F",
				},
			},
		},
	}
	for i := 0; i < b.N; i++ {
		_, _ = parameters.Get("FF.A")
	}
	b.ReportAllocs()
}

func BenchmarkReflect2(b *testing.B) {

	parameters := make(map[string]interface{}, 8)
	parameters["first_name"] = "Julian"
	parameters["emp_id"] = 302
	parameters["birth_date"] = "2016-12-15"
	parameters["obj"] = T{
		A: "test",
		FF: []T{
			{
				A: "FFFFFF",
			},
		},
	}

	for i := 0; i < b.N; i++ {
		rv := reflect.ValueOf(parameters["obj"])
		_ = rv.FieldByName("A").String()
	}
}
