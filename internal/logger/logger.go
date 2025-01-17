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

package logger

import (
	"fmt"
	"strings"
)

type (
	Logger interface {
		Info
		UserClientIP
	}
	Info interface {
		LogName() (name string)
		LogFields() (fields FieldSlice)
	}
	UserClientIP interface {
		User
		ClientIP
	}
	User     interface{ LogUser() (userdata Userdata) }
	ClientIP interface{ LogClientIP() (clientIP string) }
	Userdata struct {
		UserId   uint
		UserName string
	}
	Field struct {
		Name  string
		Value any
	}
	FieldSlice       []Field
	FieldSliceOption struct {
		FieldFormat  string
		FieldJoinSep string
	}
)

func (a FieldSlice) String(option ...func(opt *FieldSliceOption)) string {
	var opt = FieldSliceOption{FieldJoinSep: ","}
	if len(option) > 0 {
		if optFn := option[0]; optFn != nil {
			optFn(&opt)
		}
	}
	var formatS []string
	if ft := opt.FieldFormat; len(ft) > 0 {
		formatS = append(formatS, ft)
	}
	sep := opt.FieldJoinSep
	var strS []string
	for _, aa := range a {
		strS = append(strS, aa.String(formatS...))
	}
	return strings.Join(strS, sep)
}

func (a Field) String(format ...string) string {
	ft := "%s[%v]"
	if len(format) > 0 {
		if ft0 := format[0]; len(ft0) > 0 {
			ft = ft0
		}
	}
	return fmt.Sprintf(ft, a.Name, a.Value)
}
