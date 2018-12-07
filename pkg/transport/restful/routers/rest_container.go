// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package routers

import (
	"errors"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/repositories/common"
	"github.com/itcloudy/base-framework/pkg/services"
	"github.com/itcloudy/base-framework/pkg/transport/restful/controllers"
	"sync"
)

type kernel struct{}

var (
	k             *kernel
	containerOnce sync.Once
)

func restContainer() IRestContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}

type IRestContainer interface {
	UserContainer() controllers.UserController
	MenuContainer() controllers.MenuController
}

func (k *kernel) UserContainer() controllers.UserController {
	dbType := conf.Config.DB.DbType
	userService := services.UserService{}

	switch dbType {
	case "mysql":
	case "postgres":
		userService.IUserRepository = &common.UserRepository{DB: conf.DBConn}
		break

	default:
		panic(errors.New("un support sql type:" + dbType))

	}
	userController := controllers.UserController{UserService: userService}
	return userController
}
func (k *kernel) MenuContainer() controllers.MenuController {
	dbType := conf.Config.DB.DbType
	menuService := services.MenuService{}
	switch dbType {
	case "mysql":
	case "postgres":
		menuService.IMenuRepository = &common.MenuRepository{DB: conf.DBConn}
		break
	default:
		panic(errors.New("un support sql type:" + dbType))

	}
	menuController := controllers.MenuController{MenuService: menuService}
	return menuController
}
