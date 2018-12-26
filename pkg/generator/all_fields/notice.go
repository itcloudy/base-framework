// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

type Notice struct {
	ID          int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Title       string    `json:"title" gorm:"column:title" comment:"标题"`
	Description string    `json:"description"  gorm:"column:description" comment:"内容"`
	NoticeType  string    `json:"type" gorm:"column:type" comment:"类型"`
	Status      string    `json:"status" gorm:"column:status" comment:"状态"`
	Extra       string    `json:"extra" gorm:"column:extra" comment:"备注"`
}

func (mh *Notice) TableName() string {
	return "notice"
}
