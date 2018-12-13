// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/models"
)

type IMenuService interface {
	ServiceGetSelfMenu(roles []string) (menus []models.MenuList, err error)
	ServiceGetRoleByID(id string) (role models.MenuDetail, err error)
	ServiceMenuCreate(model models.MenuCreate) (result models.MenuDetail, err error)
	ServiceMenuUpdate(update models.MenuUpdate) (result models.MenuDetail, err error)
	ServiceMenuDelete(ids []string) (err error)
	ServiceGetAllMenu(page, size int, order string, query string, queryArgs ...interface{}) (menus []*models.MenuList, pagination conf.Pagination, err error)
}
