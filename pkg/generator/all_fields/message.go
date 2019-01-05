// Copyright 2018  itcloudy@qq.com  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

//消息通知
type Message struct {
	ID          int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	User        *User     `json:"user" gorm:"-" comment:"用户"`
	UserID      int       `json:"user_id" gorm:"column:user_id" comment:"用户ID"`
	Title       string    `json:"title" gorm:"column:title" comment:"标题"`
	Description string    `json:"description"  gorm:"column:description" comment:"内容"`
	NoticeType  string    `json:"type" gorm:"column:type" comment:"类型"`
	Status      string    `json:"status" gorm:"column:status" comment:"状态"`
	Extra       string    `json:"extra" gorm:"column:extra" comment:"备注"`
	Readed      bool      `json:"readed" gorm:"column:readed" comment:"已读"`
}

func (mh *Message) TableName() string {
	return "message"
}
