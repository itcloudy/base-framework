// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/models"
)

type IRoleService interface {
	ServiceGetRoleByID(id int) (role models.RoleDetail, err error)
	ServiceRoleCreate(model models.RoleCreate) (result models.RoleDetail, err error)
	ServiceRoleUpdate(update models.RoleCreate) (result models.RoleDetail, err error)
	ServiceRoleDelete(ids []int) (err error)
	ServiceGetAllRole(page, size int, order string, query string, queryArgs ...interface{}) (total int, users []*models.RoleList, pagination conf.Pagination, err error)
}
