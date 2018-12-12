// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import "github.com/itcloudy/base-framework/pkg/models"

type IRoleService interface {
	ServiceGetRoleByID(id string) (role models.RoleDetail, err error)
	ServiceRoleCreate(model models.RoleCreate) (result models.RoleDetail, err error)
	ServiceRoleUpdate(update models.RoleCreate) (result models.RoleDetail, err error)
	ServiceRoleDelete(ids []string) (err error)
	ServiceGetAllRole(offset, limit int, order string, query string, queryArgs ...interface{}) (count int, users []*models.RoleList, err error)
}
