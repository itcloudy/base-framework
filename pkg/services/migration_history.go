// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"errors"
	"github.com/hashicorp/go-version"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/migration"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/pkg/repositories"
	"github.com/itcloudy/base-framework/tools"
)

// 数据库升级服务
type MigrationService struct {
	repositories.IMigrationHistoryRepository
}

/*
获得已升级的最新版本
*/
func (service *MigrationService) GetCurrentVersion() (version string, err error) {
	return service.CurrentVersion()

}

/*
第一次初始化
*/
func (service *MigrationService) FirstMigration() (err error) {
	var (
		collection version.Collection
	)
	needUpdateMap := make(map[string]string)

	// 判断需要升级几个版本
	for _, migrate := range migration.AllInitMigrations {
		var migVersion *version.Version
		// 如果转换失败则返回
		if migVersion, err = version.NewVersion(migrate.Version); err != nil {
			return
		}
		collection = append(collection, migVersion)
		needUpdateMap[migrate.Version] = migrate.Data
	}
	// 调用repository
	if len(collection) < 1 {
		return errors.New("no need update version, code logic has problem")
	}
	return service.ApplyMigrations(collection, needUpdateMap)
}

/*
升级到某个版本，若中间存在多个，则中间版本同样升级
*/
func (service *MigrationService) UpdateToOneVersion() (err error) {
	var (
		collection version.Collection

		needVer     *version.Version
		last        *version.Version
		lastVersion string
	)
	needUpdateMap := make(map[string]string)
	if needVer, err = version.NewVersion(conf.Config.DBUpdateToVersion); err != nil {
		return
	}

	// 获得已经安装的最新版本
	if lastVersion, err = service.CurrentVersion(); err != nil {
		return
	}
	if last, err = version.NewVersion(lastVersion); err != nil {
		return
	}

	// 判断升级的版本和已升级的最后一个的大小
	if needVer.LessThan(last) || needVer.Equal(last) {
		return errors.New("need update version:  " + conf.Config.DBUpdateToVersion + " last installed version: : " + last.String())
	}
	// 判断需要升级几个版本
	for _, migrate := range migration.AllUpdateMigrations {
		var migVersion *version.Version
		// 如果转换失败则返回
		if migVersion, err = version.NewVersion(migrate.Version); err != nil {
			return
		}
		// 判断版本大小，只有大于最新版本，小于等于更新版本的才放进collection中
		if (migVersion.LessThan(needVer) && migVersion.GreaterThan(last)) || migVersion.Equal(needVer) && migVersion.GreaterThan(last) {
			collection = append(collection, migVersion)
			needUpdateMap[migrate.Version] = migrate.Data
		}
	}
	// 调用repository
	if len(collection) < 1 {
		return errors.New("no need update version, code logic has problem")
	}
	return service.ApplyMigrations(collection, needUpdateMap)
}

/*
列出所有的版本，包括系统中存在的没有安装的
*/
func (service *MigrationService) GetAllListMigration() (migrates []models.MigrationHistory, err error) {
	var installedMigrates []models.MigrationHistory
	migrates = migration.AllInitMigrations
	migrates = append(migrates, migration.AllUpdateMigrations...)
	var verSlice []string
	// 获得已经安装的版本
	installedMigrates, _ = service.ListMigration()
	//获得已安装的版本
	for _, migrate := range installedMigrates {
		verSlice = append(verSlice, migrate.Version)
	}
	//判断已经安装的版本
	for k, migrate := range migrates {
		if tools.StringInSlice(verSlice, migrate.Version) {
			migrates[k].Installed = true
		}
	}
	return
}
