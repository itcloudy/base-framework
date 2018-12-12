// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import "github.com/itcloudy/base-framework/pkg/models"

type IUserService interface {
	ServiceGetSelf(id string) (user models.UserDetail, err error)
	ServiceGetUserByID(id string) (user models.UserDetail, err error)
	ServiceGetUserByUserName(username string) (user models.UserDetail, err error)
	ServiceUserCreate(userCreate models.UserCreate) (user models.UserDetail, err error)
	ServiceUserDelete(ids []string) (err error)
	ServiceCheckUser(username, pwd string) (user models.UserDetail, err error)
	ServiceGetAllUser(offset, limit int, order string, query string, queryArgs ...interface{}) (users []*models.UserList, count int, err error)
}
