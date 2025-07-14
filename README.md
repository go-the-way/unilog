# unilog
The unified log package

# Features
- Rich Data Type Supported
- Recursion Supported
- Customizable Format
- Data Persistence
- Expr Supported
- Extendable

# Usage
```
package main

import (
	"fmt"

	"github.com/go-the-way/unilog"
)

type (
	loggerObj struct {
		Int8  int8
		Int16 int16
		Int   int
		Int32 int32
		Int64 int64

		Uint8  uint8
		Uint16 uint16
		Uint   uint
		Uint32 uint32
		Uint64 uint64

		Float32 float32
		Float64 float64
		String  string
		Map     map[string]any
		Array   [2]string
		Slice   []string

		Obj1    loggerObj1 `log:"obj_1,%s{%s}"`
		Obj2    loggerObj2 `log:",inner"`
		Ref1Val string     `log:"ref_1_val,ref:Ref1.Name,%s[%s => %s]"`
		Ref1    loggerRef1 `log:"-"`
	}
	loggerObj1 struct{ Name string }
	loggerObj2 struct{ Name string }
	loggerRef1 struct{ Name string }
)

func main() {
	obj := &loggerObj{
		String:  "Polo",
		Map:     map[string]any{"halo": "haha"},
		Array:   [2]string{"Apple", "Coco"},
		Slice:   []string{"Apple", "Coco"},
		Obj1:    loggerObj1{"Pero"},
		Obj2:    loggerObj2{"Kale"},
		Ref1Val: "Joo",
		Ref1:    loggerRef1{"Cola"},
	}
	fmt.Println(unilog.GetFields(obj).Log())
	// Outputs:
	// Int8[0],Int16[0],Int[0],Int32[0],Int64[0],Uint8[0],Uint16[0],Uint[0],Uint32[0],Uint64[0],Float32[0],
	// Float64[0],String[Polo],Map[halo:haha],Array[Apple,Coco],Slice[Apple,Coco],
	// obj_1{Name[Pero]},Name[Kale],ref_1_val[Cola => Joo]
}
```