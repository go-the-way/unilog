# unilog
The unified log package

# Features
- Rich Data Type Supported
- Recursion Supported
- Customizable Format
- Data Persistence
- Expr Supported
- Extendable

# Tag Definition
- Name `log:"log_name"`, default: `field_name`
- Format `log:",%s[%v]"`, default: `%s[%v]`
- Ref 1 `log:",ref:Ref1.Var,%s[ref:%v]"`
- Ref 2 `log:",ref:Ref2.Var,%s[ref:%v self:%v]"`
- Transform 1 `log:",transform:1->on|2->off,%s[transformed:%v]"`
- Transform 2 `log:",transform:1->on|2->off,%s[transformed:%v self:%v]"`

# Usage
```
package main

import (
	"fmt"

	"github.com/go-the-way/unilog"
)

type (
	logger struct {
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

		Obj       loggerObj `log:""`
		ObjFormat loggerObj `log:",%s{%v}"`
		ObjInner  loggerObj `log:",inner"`

		RefObj                  string `log:",ref:Obj"`
		RefObjFormat            string `log:",ref:Obj,%s[ref_obj:%s]"`
		RefObjField             string `log:",ref:Obj.Name"`
		RefObjFieldFormat       string `log:",ref:Obj.Name,%s[ref_field:%s]"`
		RefUndefinedObj         string `log:",ref:Any"`
		RefUndefinedObjFormat   string `log:",ref:Any,%s[ref_undefined_obj:%v]"`
		RefUndefinedField       string `log:",ref:Any.Field"`
		RefUndefinedFieldFormat string `log:",ref:Any.Field,%s[ref_undefined_field:%v]"`

		Ref2ObjFormat            byte `log:",ref:Obj,%s[ref_obj:%s self:%d]"`
		Ref2ObjFieldFormat       byte `log:",ref:Obj.Name,%s[ref_field:%s self:%d]"`
		Ref2UndefinedObjFormat   byte `log:",ref:Any,%s[ref_undefined_obj:%v self:%d]"`
		Ref2UndefinedFieldFormat byte `log:",ref:Any.Field,%s[ref_undefined_field:%v self:%d]"`

		Transform       byte `log:",transform:0->unknown|1->on|2->off"`
		TransformFormat byte `log:",transform:0->unknown|1->on|2->off,%s[transformed:%s self:%d]"`
	}
	loggerObj struct{ Name string }
)

func main() {
	obj := &logger{
		String: "Polo",
		Map:    map[string]any{"name": "Halo"},
		Array:  [2]string{"Apple", "Banana"},
		Slice:  []string{"Orange", "Lemon"},
		Obj:    loggerObj{"Kellen"},
	}
	fmt.Println(unilog.GetFields(obj).Log())
	// Outputs:
	// Int8[0],Int16[0],Int[0],Int32[0],Int64[0],Uint8[0],Uint16[0],Uint[0],Uint32[0],Uint64[0],Float32[0],Float64[0],String[Polo],
	// Map[name:Halo],Array[Apple,Banana],Slice[Orange,Lemon],
	// Obj[Name[Kellen]],ObjFormat{Name[]},Name[],
	// RefObj[{Kellen}],RefObjFormat[ref_obj:{Kellen}],RefObjField[Kellen],RefObjFieldFormat[ref_field:Kellen],
	// RefUndefinedObj[<nil>],RefUndefinedObjFormat[ref_undefined_obj:<nil>],RefUndefinedField[<nil>],
	// RefUndefinedFieldFormat[ref_undefined_field:<nil>],
	// Ref2ObjFormat[ref_obj:{Kellen} self:0],Ref2ObjFieldFormat[ref_field:Kellen self:0],
	// Ref2UndefinedObjFormat[ref_undefined_obj:<nil> self:0],Ref2UndefinedFieldFormat[ref_undefined_field:<nil> self:0],
	// Transform[unknown],TransformFormat[transformed:unknown self:0]
}
```