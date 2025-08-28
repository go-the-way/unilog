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

	"github.com/go-the-way/unilog/internal/logger"
)

// CallbackFunc defines a function type for processing a LogAddReq.
type CallbackFunc func(req *LogAddReq)

// Callback generates a logging function for a given Logger type, applying optional callback functions.
// It constructs a log entry using the Logger's name, fields, user data, and client IP, then passes it to LogAdd.
// The optional callback function can modify the LogAddReq before it is logged.
func Callback[LOG logger.Logger](opts ...CallbackFunc) func(req LOG) {
	return func(req LOG) {
		// Get the log name, defaulting to "unknown" if not provided.
		logName := req.LogName()
		if logName == "" {
			logName = "unknown"
		}

		// Generate the fields content by logging the fields slice, if any.
		fieldsContent := ""
		if len(req.LogFields()) > 0 {
			fieldsContent = req.LogFields().Log()
		}

		// Retrieve user data and client IP from the Logger.
		userdata := req.LogUser()
		clientIP := req.LogClientIP()

		// Construct the log content by combining the log name and fields content.
		content := fmt.Sprintf("%s{%s}", logName, fieldsContent)

		// Create a LogAddReq instance with user data, client IP, and content.
		req0 := LogAddReq{
			UserId:   userdata.UserId,
			UserName: userdata.UserName,
			ClientIP: clientIP,
			Content:  content,
		}

		// Apply the optional callback function to modify the LogAddReq, if provided.
		if len(opts) > 0 {
			if opt := opts[0]; opt != nil {
				opt(&req0)
			}
		}

		// Pass the constructed LogAddReq to the LogAdd function for logging.
		LogAdd(req0)
	}
}
