// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package migration

import (
	"github.com/hashicorp/go-version"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/logs"
	"go.uber.org/zap"
)

type migration struct {
	version string // 版本
	data    string // sql语句

}

type database interface {
	CurrentVersion() (string, error)
	ApplyMigration(string, string) error
}

func migrate(db database, appVer *version.Version, migrations []*migration) error {
	dbVerString, err := db.CurrentVersion()
	if err != nil {
		logs.Logger.Error("get migrate version failed", zap.Error(err))
		return err
	}

	dbVer, err := version.NewVersion(dbVerString)
	if err != nil {
		logs.Logger.Error("parse version failed", zap.Error(err), zap.String("version", dbVerString))

		return err
	}

	// if the database version is up-to-date
	if !dbVer.LessThan(appVer) {
		return nil
	}

	for _, m := range migrations {
		mgrVer, err := version.NewVersion(m.version)
		if err != nil {
			logs.Logger.Error("parse version failed", zap.Error(err), zap.String("version", m.version))
			return err
		}
		if !dbVer.LessThan(mgrVer) {
			continue
		}
		err = db.ApplyMigration(m.version, m.data)
		if err != nil {
			logs.Logger.Error("apply migration failed", zap.Error(err), zap.String("version", m.version))

			return err
		}
		logs.Logger.Info("apply migration success", zap.String("version", m.version))
	}

	return nil
}

func runMigrations(db database, migrationList []*migration) error {
	appVer, err := version.NewVersion(consts.VERSION)
	if err != nil {
		logs.Logger.Error("parse version failed", zap.Error(err))

		return err
	}
	return migrate(db, appVer, migrationList)
}

// InitMigrate applies initial migrations
func InitMigrate(db database) error {
	return runMigrations(db, migrations)
}

// UpdateMigrate applies update migrations
func UpdateMigrate(db database) error {
	return runMigrations(db, updateMigrations)
}
