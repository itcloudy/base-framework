// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package repositories

import (
	"github.com/itcloudy/base-framework/pkg/models"
)

type IUserRepository interface {
	//根据用户名查找用户
	FindUserByUserName(username string) (models.User, error)
	//根据用户ID查找用户
	FindUserByID(id string) (models.User, error)
}
