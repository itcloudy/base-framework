// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package routers

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/repositories/common"
	"github.com/itcloudy/base-framework/pkg/services"
	"github.com/itcloudy/base-framework/pkg/transport/restful/controllers"
)

func (k *kernel) IndexContainer() controllers.IndexController {
	return controllers.IndexController{}
}
func (k *kernel) UserContainer() controllers.UserController {
	service := services.UserService{}

	service.IUserRepository = &common.UserRepository{}
	service.DB = conf.DBConn
	controller := controllers.UserController{IUserService: &service}
	return controller
}
func (k *kernel) RoleContainer() controllers.RoleController {
	service := services.RoleService{}
	service.IRoleRepository = &common.RoleRepository{}
	service.DB = conf.DBConn
	controller := controllers.RoleController{RoleService: service}
	return controller
}

func (k *kernel) SystemAPIContainer() controllers.SystemAPIController {
	service := services.SystemAPIService{}
	service.ISystemAPIRepository = &common.SystemAPIRepository{}
	service.DB = conf.DBConn
	controller := controllers.SystemAPIController{SystemAPIService: service}
	return controller
}
