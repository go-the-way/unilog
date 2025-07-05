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
	FieldSlice []Field
)
