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

package log

import (
	"github.com/go-the-way/unilog/internal/models"
	"github.com/go-the-way/unilog/internal/pkg"
)

func (req *AddReq) transform() *models.Log {
	return &models.Log{
		UserId:     req.UserId,
		UserName:   req.UserName,
		ClientIP:   req.ClientIP,
		Type1:      req.Type1,
		Type2:      req.Type2,
		Type3:      req.Type3,
		Content:    req.Content,
		CreateTime: pkg.TimeNowStr(),
		UpdateTime: pkg.TimeNowStr(),
	}
}

func (req *UpdateReq) transform() map[string]any {
	return map[string]any{
		"user_id":     req.UserId,
		"user_name":   req.UserName,
		"client_ip":   req.ClientIP,
		"type1":       req.Type1,
		"type2":       req.Type2,
		"type3":       req.Type3,
		"content":     req.Content,
		"update_time": pkg.TimeNowStr(),
	}
}
