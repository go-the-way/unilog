package models

type (
	Log       = UnilogLog
	UnilogLog struct {
		Id         uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:日志Id" json:"id"`                // 日志Id
		UserId     uint   `gorm:"column:user_id;type:uint;not null;default:0;comment:用户Id" json:"user_id"`                 // 用户Id
		UserName   string `gorm:"column:user_name;type:varchar(20);not null;default:'';comment:用户名称" json:"user_name"`     // 用户名称
		ClientIP   string `gorm:"column:client_ip;type:varchar(50);not null;default:'';comment:客户端IP" json:"client_ip"`    // 客户端IP
		Content    string `gorm:"column:content;type:varchar(500);not null;default:'';comment:日志内容" json:"content"`        // 日志内容
		CreateTime string `gorm:"column:create_time;type:varchar(20);not null;default:'';comment:创建时间" json:"create_time"` // 创建时间
		UpdateTime string `gorm:"column:update_time;type:varchar(20);not null;default:'';comment:修改时间" json:"update_time"` // 修改时间
	}
)
