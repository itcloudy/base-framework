// Copyright 2018  itcloudy@qq.com  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

//广播
type Broadcast struct {
	ID          int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Title       string    `json:"title" gorm:"column:title" comment:"标题"`
	Description string    `json:"description"  gorm:"column:description" comment:"内容"`
	Sended      bool      `json:"sended"  gorm:"column:sended" comment:"已发"`
}

func (mh *Broadcast) TableName() string {
	return "broadcast"
}
