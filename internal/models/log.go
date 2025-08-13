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

package models

type (
	Log       = UnilogLog
	UnilogLog struct {
		Id         uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:日志Id" json:"id"`                    // 日志Id
		UserId     uint   `gorm:"column:user_id;type:uint;not null;default:0;comment:用户Id;index" json:"user_id"`               // 用户Id
		UserName   string `gorm:"column:user_name;type:varchar(100);not null;default:'';comment:用户名称;index" json:"user_name"`  // 用户名称
		ClientIP   string `gorm:"column:client_ip;type:varchar(100);not null;default:'';comment:客户端IP;index" json:"client_ip"` // 客户端IP
		Type1      string `gorm:"column:type1;type:varchar(100);not null;default:'';comment:类型1;index" json:"type1"`           // 类型1
		Type2      string `gorm:"column:type2;type:varchar(100);not null;default:'';comment:类型2;index" json:"type2"`           // 类型2
		Type3      string `gorm:"column:type3;type:varchar(100);not null;default:'';comment:类型3;index" json:"type3"`           // 类型3
		Content    string `gorm:"column:content;type:varchar(500);not null;default:'';comment:日志内容" json:"content"`            // 日志内容
		CreateTime string `gorm:"column:create_time;type:varchar(20);not null;default:'';comment:创建时间" json:"create_time"`     // 创建时间
		UpdateTime string `gorm:"column:update_time;type:varchar(20);not null;default:'';comment:修改时间" json:"update_time"`     // 修改时间
	}
)
