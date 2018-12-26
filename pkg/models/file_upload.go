// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package models

import "time"

type FileUploadCreate struct {
	ID       int    `json:"id" gorm:"column:id" comment:"主键ID"`
	UserID   int    `json:"user_id" gorm:"column:user_id" comment:"用户ID"`
	FileName string `json:"file_name" gorm:"column:file_name" comment:"文件名"`
	Address  string `json:"address" gorm:"column:address" comment:"文件目标位置"`
	Type     string `json:"type" gorm:"column:type" comment:"存储类型"`
}

func (mh *FileUploadCreate) TableName() string {
	return "file_upload"
}

type FileUploadDetail struct {
	ID        int         `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt time.Time   `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt time.Time   `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	User      *UserDetail `json:"user" gorm:"-" comment:"用户"`
	UserID    int         `json:"user_id" gorm:"column:user_id" comment:"用户ID"`
	FileName  string      `json:"file_name" gorm:"column:file_name" comment:"文件名"`
	Address   string      `json:"address" gorm:"column:address" comment:"文件Hash"`
	Type      string      `json:"type" gorm:"column:type" comment:"存储类型"`
}

func (mh *FileUploadDetail) TableName() string {
	return "file_upload"
}

type FileUploadList struct {
	ID        int         `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt time.Time   `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt time.Time   `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	User      *UserDetail `json:"user" gorm:"-" comment:"用户"`
	UserID    int         `json:"user_id" gorm:"column:user_id" comment:"用户ID"`
	FileName  string      `json:"file_name" gorm:"column:file_name" comment:"文件名"`
	Address   string      `json:"address" gorm:"column:address" comment:"文件Hash"`
	Type      string      `json:"type" gorm:"column:type" comment:"存储类型"`
}

func (mh *FileUploadList) TableName() string {
	return "file_upload"
}
