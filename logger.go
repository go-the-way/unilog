// Copyright 2025 unilog Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unilog

import (
	"fmt"
	"reflect"

	"github.com/go-the-way/unilog/internal/logger"
)

var supportedKind = map[reflect.Kind]struct{}{
	reflect.Bool:       {},
	reflect.Int:        {},
	reflect.Int8:       {},
	reflect.Int16:      {},
	reflect.Int32:      {},
	reflect.Int64:      {},
	reflect.Uint:       {},
	reflect.Uint8:      {},
	reflect.Uint16:     {},
	reflect.Uint32:     {},
	reflect.Uint64:     {},
	reflect.Float32:    {},
	reflect.Float64:    {},
	reflect.Complex64:  {},
	reflect.Complex128: {},
	reflect.Array:      {},
	reflect.Map:        {},
	reflect.Slice:      {},
	reflect.String:     {},
	reflect.Struct:     {},
}

func Callback[LOG logger.Logger](option ...func(req LOG) (laReq LogAddReq)) func(req LOG) {
	return func(req LOG) {
		logName := req.LogName()
		if logName == "" {
			logName = "unknown"
		}
		fieldsContent := ""
		if len(req.LogFields()) > 0 {
			fieldsContent = req.LogFields().String()
		}
		userdata := req.LogUser()
		clientIP := req.LogClientIP()
		content := fmt.Sprintf("%s{%s}", logName, fieldsContent)
		laReq := LogAddReq{UserId: userdata.UserId, UserName: userdata.UserName, ClientIP: clientIP, Content: content}
		if len(option) > 0 {
			if opt := option[0]; opt != nil {
				laReq = opt(req)
			}
		}
		LogAdd(laReq)
	}
}

func GetFieldsFromTag(struct0 any) (fields []Field) {
	v := reflect.ValueOf(struct0)
	if !v.IsValid() {
		panic("unilog: the struct value is invalid.")
	}
	if v = rv(v); v.Kind() != reflect.Struct {
		panic("unilog: the struct value is not supported.")
	}
	return getSupportedFields(v)
}

func getSupportedFields(v reflect.Value) (fields []Field) {
	for i := 0; i < v.NumField(); i++ {
		fd := v.Type().Field(i)
		fdv := v.Field(i)
		logName, ok := fd.Tag.Lookup("log")
		if !ok {
			logName = fd.Name
		}
		if logName != "-" && v.IsValid() {
			fdv = rv(fdv)
			if _, supported := supportedKind[fdv.Kind()]; supported {
				if fdv.Kind() == reflect.Struct {
					fields = append(fields, getSupportedFields(fdv)...)
				} else if fdv.CanInterface() {
					fields = append(fields, Field{Name: logName, Value: fdv.Interface()})
				}
			}
		}
	}
	return
}

func rv(v reflect.Value) (vv reflect.Value) {
	vv = v
	for vv.Kind() == reflect.Pointer {
		vv = vv.Elem()
	}
	return
}
