// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import "github.com/itcloudy/base-framework/pkg/models"

type IUserService interface {
	GetSelf(id string) (user models.UserDetail, err error)
	GetUserByID(id string) (user models.UserDetail, err error)
	GetUserByUserName(username string) (user models.UserDetail, err error)
	UserCreate(userCreate models.UserCreate) (user models.UserDetail, err error)
	CheckUser(username, pwd string) (user models.UserDetail, err error)
}
