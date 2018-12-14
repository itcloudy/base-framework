// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/models"
)

type IUserService interface {
	ServiceGetSelf(id int) (user models.UserDetail, err error)
	ServiceGetUserByID(id int) (user models.UserDetail, err error)
	ServiceGetUserByUserName(username string) (user models.UserDetail, err error)
	ServiceUserCreate(userCreate models.UserCreate) (user models.UserDetail, err error)
	ServiceUserDelete(ids []int) (err error)
	ServiceCheckUser(username, pwd string) (user models.UserDetail, err error)
	ServiceGetAllUser(page, size int, order string, query string, queryArgs ...interface{}) (users []*models.UserList, pagination conf.Pagination, err error)
}
