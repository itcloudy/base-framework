// Copyright 2018  itcloudy@qq.com  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

//角色拥有api
type RoleApi struct {
	ID          int        `json:"id" gorm:"column:id" comment:"主键ID"`                             //主键
	CreatedAt   time.Time  `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"` //记录创建时间
	UpdatedAt   time.Time  `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Role        *Role      `json:"role" yaml:"role" comment:"角色"`
	RoleID      int        `json:"role_id" yaml:"role_id" comment:"角色ID"`
	SystemApi   *SystemApi `json:"system_api" yaml:"system_api" comment:"系统接口"`
	SystemApiID int        `json:"system_api_id" yaml:"system_api_id" comment:"系统接口ID"`
}

func (mh *RoleApi) TableName() string {
	return "role_api"
}
