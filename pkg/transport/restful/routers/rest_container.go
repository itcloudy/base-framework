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

func restContainer() iRestContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}

type iRestContainer interface {
	IndexContainer() controllers.IndexController
	UserContainer() controllers.UserController
	MenuContainer() controllers.MenuController
	RoleContainer() controllers.RoleController
	SystemAPIContainer() controllers.SystemAPIController
}

func (k *kernel) IndexContainer() controllers.IndexController {
	return controllers.IndexController{}
}
func (k *kernel) UserContainer() controllers.UserController {
	service := services.UserService{}

	service.IUserRepository = &common.UserRepository{}
	dbType := conf.Config.DB.DbType
	switch dbType {
	case "mysql":
	case "postgres":
		service.DB = conf.DBConn
		break
	default:
		panic(errors.New("un support sql type:" + dbType))

	}
	controller := controllers.UserController{IUserService: &service}
	return controller
}
func (k *kernel) RoleContainer() controllers.RoleController {
	service := services.RoleService{}
	service.IRoleRepository = &common.RoleRepository{}
	dbType := conf.Config.DB.DbType
	switch dbType {
	case "mysql":
	case "postgres":
		service.DB = conf.DBConn
		break
	default:
		panic(errors.New("un support sql type:" + dbType))

	}
	controller := controllers.RoleController{RoleService: service}
	return controller
}
func (k *kernel) MenuContainer() controllers.MenuController {
	service := services.MenuService{}
	service.IMenuRepository = &common.MenuRepository{}
	dbType := conf.Config.DB.DbType
	switch dbType {
	case "mysql":
	case "postgres":
		service.DB = conf.DBConn
		break
	default:
		panic(errors.New("un support sql type:" + dbType))

	}
	controller := controllers.MenuController{MenuService: service}
	return controller
}
func (k *kernel) SystemAPIContainer() controllers.SystemAPIController {
	service := services.SystemAPIService{}
	service.ISystemAPIRepository = &common.SystemAPIRepository{}
	dbType := conf.Config.DB.DbType
	switch dbType {
	case "mysql":
	case "postgres":
		service.DB = conf.DBConn
		break
	default:
		panic(errors.New("un support sql type:" + dbType))

	}
	controller := controllers.SystemAPIController{SystemAPIService: service}
	return controller
}
