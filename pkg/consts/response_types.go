// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package consts

const (
	//  正常返回
	Success = 20000 //正常返回
	//controller 层的错误代码为4000开始
	BindingJsonErr = 40000 // 绑定json失败
	GetFileErr     = 40008 // 获取上传文件失败
	PathParamErr   = 40008 // url参数错误
	UpdateIdErr    = 40009 // 更新ID错误

	//service 层的错误代码从50000开始	common.GenResponse(c, consts.Success, response, "")
	ServerErr             = 50000 //
	ValidateErr           = 50005 // 数据验证失败
	UserNameOrPasswordErr = 50006 // 用户名或密码错误
	PermissionErr         = 50007 // 没有权限
	TokenValidErr         = 50008 // token无效
	FilelUploadErr        = 50009 //文件上传失败
	MigrationErr          = 50010 // 升级失败

	//repository层的错误代码从60000开始
	DBInSertErr = 60000 // 数据插入失败
	DBUpdateErr = 60001 // 数据更新失败
	DBSelectErr = 60002 // 数据查询失败
	DBDeleteErr = 60003 // 数据删除失败

)
