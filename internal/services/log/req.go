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

type (
	GetPageReq struct {
		Page  int `form:"page"`
		Limit int `form:"limit"`

		OrderBy string `form:"order_by" json:"order_by"` // 排序

		Id          uint   `form:"id"`           // 日志Id
		UserId      uint   `form:"user_id"`      // 用户Id
		UserName    string `form:"user_name"`    // 用户名
		ClientIP    string `form:"client_ip"`    // 客户端IP
		Type1       string `form:"type1"`        // 类型1
		Type2       string `form:"type2"`        // 类型2
		Type3       string `form:"type3"`        // 类型3
		Content     string `form:"content"`      // 日志内容
		CreateTime1 string `form:"create_time1"` // 创建时间
		CreateTime2 string `form:"create_time2"` // 创建时间
		UpdateTime1 string `form:"update_time1"` // 修改时间
		UpdateTime2 string `form:"update_time2"` // 修改时间
	}
	IdReq struct {
		Id uint `validate:"min(1,日志Id不能为空)" json:"id"`
	}
	GetReq IdReq
	AddReq struct {
		UserId   uint   `json:"user_id"`   // 用户Id
		UserName string `json:"user_name"` // 用户名称
		ClientIP string `json:"client_ip"` // 客户端IP
		Type1    string `json:"type1"`     // 类型1
		Type2    string `json:"type2"`     // 类型2
		Type3    string `json:"type3"`     // 类型3
		Content  string `json:"content"`   // 日志内容
		Callback func(req AddReq)
	}
	UpdateReq struct {
		IdReq    `validate:"valid(T)"`
		AddReq   `validate:"valid(T)"`
		Callback func(req UpdateReq)
	}
	DeleteReq struct {
		IdReq
		Callback func(req DeleteReq)
	}
)
