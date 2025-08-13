// Copyright 2025 icmpkt Author. All Rights Reserved.
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

type CallbackFunc func(req *LogAddReq)

func Callback[LOG logger.Logger](opts ...CallbackFunc) func(req LOG) {
	return func(req LOG) {
		logName := req.LogName()
		if logName == "" {
			logName = "unknown"
		}
		fieldsContent := ""
		if len(req.LogFields()) > 0 {
			fieldsContent = req.LogFields().Log()
		}
		userdata := req.LogUser()
		clientIP := req.LogClientIP()
		content := fmt.Sprintf("%s{%s}", logName, fieldsContent)
		req0 := LogAddReq{UserId: userdata.UserId, UserName: userdata.UserName, ClientIP: clientIP, Content: content}
		if len(opts) > 0 {
			if opt := opts[0]; opt != nil {
				opt(&req0)
			}
		}
		LogAdd(req0)
	}
}
