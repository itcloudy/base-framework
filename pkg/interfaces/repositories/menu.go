// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package repositories

import (
	"github.com/itcloudy/base-framework/pkg/models"
)

type IMenuRepository interface {
	//获得某些角色的菜单信息
	FindMenuByRoles(roleSlice []string) (menus []models.Menu, err error)
}
