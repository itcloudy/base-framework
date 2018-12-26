// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

type FileUpload struct {
	ID        int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	User      *User     `json:"user" gorm:"-" comment:"用户"`
	UserID    int       `json:"user_id" gorm:"column:user_id" comment:"用户ID"`
	FileName  string    `json:"file_name" gorm:"column:file_name" comment:"文件名"`
	Target    string    `json:"target" gorm:"column:target" comment:"目标文件"`
	Type      string    `json:"type" gorm:"column:type" comment:"存储类型"`
}

func (mh *FileUpload) TableName() string {
	return "file_upload"
}
