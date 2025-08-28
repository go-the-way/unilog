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

// UserId returns a CallbackFunc that sets the UserId field of a LogAddReq.
func UserId(a uint) CallbackFunc {
	return func(req *LogAddReq) {
		req.UserId = a
	}
}

// UserName returns a CallbackFunc that sets the UserName field of a LogAddReq.
func UserName(a string) CallbackFunc {
	return func(req *LogAddReq) {
		req.UserName = a
	}
}

// ClientIP returns a CallbackFunc that sets the ClientIP field of a LogAddReq.
func ClientIP(a string) CallbackFunc {
	return func(req *LogAddReq) {
		req.ClientIP = a
	}
}

// Type1 returns a CallbackFunc that sets the Type1 field of a LogAddReq.
func Type1(a string) CallbackFunc {
	return func(req *LogAddReq) {
		req.Type1 = a
	}
}

// Type2 returns a CallbackFunc that sets the Type2 field of a LogAddReq.
func Type2(a string) CallbackFunc {
	return func(req *LogAddReq) {
		req.Type2 = a
	}
}

// Type3 returns a CallbackFunc that sets the Type3 field of a LogAddReq.
func Type3(a string) CallbackFunc {
	return func(req *LogAddReq) {
		req.Type3 = a
	}
}

// Type4 returns a CallbackFunc that sets the Type4 field of a LogAddReq.
func Type4(a string) CallbackFunc {
	return func(req *LogAddReq) {
		req.Type4 = a
	}
}

// Type5 returns a CallbackFunc that sets the Type5 field of a LogAddReq.
func Type5(a string) CallbackFunc {
	return func(req *LogAddReq) {
		req.Type5 = a
	}
}

// Content returns a CallbackFunc that sets the Content field of a LogAddReq.
func Content(a string) CallbackFunc {
	return func(req *LogAddReq) {
		req.Content = a
	}
}

// Type1Admin returns a CallbackFunc that sets the Type1 field of a LogAddReq to "admin".
func Type1Admin() CallbackFunc {
	return Type1("admin")
}

// Type1User returns a CallbackFunc that sets the Type1 field of a LogAddReq to "user".
func Type1User() CallbackFunc {
	return Type1("user")
}
