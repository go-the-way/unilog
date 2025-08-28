# unilog

The `unilog` package is a powerful logging utility for Go applications, designed to provide structured, customizable, and extensible logging capabilities. It supports logging of various Go data types, including basic types, composite types, and nested structs, with advanced features like field referencing, value transformation, and integration of user and client metadata. The package is highly configurable, allowing developers to define log output formats using struct field tags and customize logging behavior through callbacks.

## Features

- **Rich Data Type Support**: Logs basic types (bool, integers, floats, strings), composite types (maps, arrays, slices), and nested structs.
- **Customizable Output**: Define log formats using `log` struct tags for precise control over field representation.
- **Field Referencing**: Reference other fields or structs in the log output using the `ref` tag option.
- **Value Transformation**: Transform field values into human-readable strings using the `transform` tag option.
- **Recursive Logging**: Automatically processes nested structs and arrays/slices with customizable formatting.
- **User and Client Metadata**: Integrates user data (ID, name) and client IP addresses via the `Logger` interface.
- **Extensible Callbacks**: Supports custom log processing through callback functions.
- **Configurable Formatting**: Customize array, map, and field separator formats for flexible output.

## Installation

To install the `unilog` package, run the following command:

```bash
go get github.com/go-the-way/unilog
```

## Tag Definitions

The `unilog` package uses the `log` struct tag to configure logging behavior for struct fields. The supported tag options are:

- **Name**: Specifies a custom field name for the log output. Defaults to the struct field name.
    - Example: `log:"custom_name"`
- **Format**: Defines a custom `printf`-style format string for the field. Defaults to `%s[%v]`.
    - Example: `log:",%s[%v]"`
- **Reference (ref)**: References another field or struct in the log output.
    - Simple reference: `log:",ref:Obj"`
    - Formatted reference: `log:",ref:Obj,%s[ref:%s]"`
    - Reference with self: `log:",ref:Obj,%s[ref:%s self:%d]"`
- **Transform**: Maps field values to human-readable strings using a transformation expression.
    - Simple transform: `log:",transform:0->unknown|1->on|2->off"`
    - Formatted transform: `log:",transform:0->unknown|1->on|2->off,%s[transformed:%s self:%d]"`
- **Inline**: Recursively logs nested struct fields as if they were part of the parent struct.
    - Example: `log:",inline"`

## Usage

The `unilog` package provides a simple yet powerful API for logging structs and their fields. The primary functions are `GetFields` for extracting loggable fields and `Callback` for generating logging functions for types implementing the `Logger` interface. Below is an example demonstrating its usage:

```go
package main

import (
	"fmt"
	
	"github.com/go-the-way/unilog"
)

type (
	// Logger struct with various fields and log tags.
	logger struct {
		Int8    int8
		Int16   int16
		Uint8   uint8
		Float32 float32
		String  string
		Map     map[string]any
		Array   [2]string
		Slice   []string
		Obj     loggerObj `log:"obj"`
		ObjFormat loggerObj `log:",%s{%v}"`
		ObjInner  loggerObj `log:",inline"`
		RefObj    string    `log:",ref:obj"`
		RefObjFormat string `log:",ref:obj,%s[ref_obj:%s]"`
		RefObjField  string `log:",ref:obj.Name"`
		Transform byte `log:",transform:0->unknown|1->on|2->off"`
		TransformFormat byte `log:",transform:0->unknown|1->on|2->off,%s[transformed:%s self:%d]"`
	}
	// Nested struct for testing recursive logging.
	loggerObj struct {
		Name string
	}
	// wrapper struct implementing the Logger interface.
	wrapper struct {
		*logger
		userdata unilog.Userdata
		clientIP string
	}
)

// Implement Logger interface methods for wrapper.
func (w *wrapper) LogName() string {
	return "example"
}
func (w *wrapper) LogFields() unilog.FieldSlice {
	return unilog.GetFields(w.logger)
}
func (w *wrapper) LogUser() unilog.Userdata {
	return w.userdata
}
func (w *wrapper) LogClientIP() string {
	return w.clientIP
}

func main() {
	// Create a logger instance with sample data.
	obj := &logger{
		String: "Polo",
		Map:    map[string]any{"name": "Halo"},
		Array:  [2]string{"Apple", "Banana"},
		Slice:  []string{"Orange", "Lemon"},
		Obj:    loggerObj{"Kellen"},
	}
	// Wrap the logger in a wrapper struct.
	w := &wrapper{
		logger:   obj,
		userdata: unilog.Userdata{UserId: 123, UserName: "JohnDoe"},
		clientIP: "192.168.1.1",
	}
	// Create a logging function with a callback.
	logFunc := unilog.Callback[*wrapper](func(req *unilog.LogAddReq) {
		fmt.Printf("Logging: UserID=%d, UserName=%s, ClientIP=%s, Content=%s\n",
			req.UserId, req.UserName, req.ClientIP, req.Content)
	})
	// Execute the logging function.
	logFunc(w)
}
```

### Output

The above code produces output similar to the following (depending on the `LogAdd` implementation):

```
Logging: UserID=123, UserName=JohnDoe, ClientIP=192.168.1.1, Content=example{Int8[0],Int16[0],Uint8[0],Float32[0],String[Polo],Map[name:Halo],Array[Apple,Banana],Slice[Orange,Lemon],obj[Name[Kellen]],ObjFormat{Name[Kellen]},Name[Kellen],RefObj[{Name[Kellen]}],RefObjFormat[ref_obj:{Name[Kellen]}],RefObjField[Kellen],Transform[unknown],TransformFormat[transformed:unknown self:0]}
```

## How It Works

1. **Struct Field Tags**: Use the `log` tag to define how each field is logged, specifying custom names, formats, references, transformations, or inline behavior.
2. **GetFields**: The `unilog.GetFields` function extracts loggable fields from a struct, processing tags and handling nested structs, arrays, and slices.
3. **FieldSlice and Log**: The `FieldSlice` type aggregates fields and generates a concatenated log string using the `Log` method, with fields separated by a configurable separator.
4. **Logger Interface**: Types implementing the `Logger` interface provide log name, fields, user data, and client IP for comprehensive logging.
5. **Callback Function**: The `Callback` function generates a logging function that constructs a `LogAddReq` with user data, client IP, and formatted content, passing it to `LogAdd` for processing.
6. **Custom Formatting**: Configure array, map, and field separator formats using `SetArrayFunc`, `SetMapFunc`, `SetFieldFormat`, `SetArrayElementFormat`, and `SetFieldJoinSep`.

## Advanced Features

### Custom Formatting

Customize the log output for a field using a `printf`-style format string:

```
ObjFormat loggerObj `log:",%s{%v}"`
```

This formats the `ObjFormat` field as `ObjFormat{Name[Kellen]}` instead of the default `obj[Name[Kellen]]`.

### Field Referencing

Reference other fields or structs using the `ref` tag option:

```
RefObjField string `log:",ref:obj.Name"`
```

This logs the `Name` field of the `obj` struct, enabling contextual logging.

### Value Transformation

Transform field values into human-readable strings using the `transform` tag option:

```
Transform byte `log:",transform:0->unknown|1->on|2->off"`
```

This maps the byte values `0`, `1`, and `2` to `unknown`, `on`, and `off`, respectively.

### Inline Structs

Log nested struct fields as part of the parent struct using the `inline` tag:

```
ObjInner loggerObj `log:",inline"`
```

This includes the `Name` field of `ObjInner` directly in the parent struct's log output.

### Custom Array and Map Formatting

Customize how arrays and maps are formatted:

```
unilog.SetArrayFunc(unilog.arrayFunc("%v", "|")) // Use "|" as array element separator.
unilog.SetMapFunc(unilog.mapFunc("%s=%v", ";"))  // Use ";" as map pair separator.
```

### Callback Customization

Use the `Callback` function with a custom callback to modify the `LogAddReq` before logging:

```
logFunc := unilog.Callback[*wrapper](func(req *unilog.LogAddReq) {
	req.Content = "Custom: " + req.Content
})
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on the [GitHub repository](https://github.com/go-the-way/unilog) with your suggestions, bug reports, or improvements.

## License

This project is licensed under the Apache License. See the [LICENSE](LICENSE) file for details.