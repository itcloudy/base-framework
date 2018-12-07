// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package models

import "time"

type User struct {
	ID        int       `json:"id" db:"id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at" comment:"记录更新时间"`
	Username  string    `json:"username" db:"username"` // 用户名
	Alias     string    `json:"alias" db:"alias"`       //昵称
}

// TableName returns name of table
func (mh *User) TableName() string {
	return "users"
}
