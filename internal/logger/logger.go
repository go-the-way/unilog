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

// Logger is an interface combining LogInfo and LogUserClientIP for comprehensive logging.
type Logger interface {
	LogInfo
	LogUserClientIP
}

// LogInfo defines methods for retrieving log name and fields.
type LogInfo interface {
	// LogName returns the name associated with the log entry.
	LogName() (name string)
	// LogFields returns a slice of fields for logging.
	LogFields() (fields FieldSlice)
}

// LogUserClientIP combines LogUser and LogClientIP interfaces for user and client IP logging.
type LogUserClientIP interface {
	LogUser
	LogClientIP
}

// LogUser defines a method for retrieving user data for logging.
type LogUser interface {
	// LogUser returns the user data associated with the log entry.
	LogUser() (userdata Userdata)
}

// LogClientIP defines a method for retrieving the client IP address for logging.
type LogClientIP interface {
	// LogClientIP returns the client IP address associated with the log entry.
	LogClientIP() (clientIP string)
}

// Userdata represents user information for logging.
type Userdata struct {
	UserId   uint   // UserId is the unique identifier of the user.
	UserName string // UserName is the name of the user.
}
