// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package models

import "time"

// MigrationHistoryRepository is model
type MigrationHistory struct {
	ID        int       `json:"id" db:"id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty"  db:"created_at" comment:"记录创建时间"`
	Version   string    `json:"version" db:"version" comment:"版本"`
	Data      string    `json:"data" db:"data" comment:"升级数据"`
	Installed bool      `json:"installed" db:"-" comment:"升级"`
}

// TableName returns name of table
func (mh *MigrationHistory) TableName() string {
	return "migration_history"
}
