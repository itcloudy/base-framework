// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"github.com/hashicorp/go-version"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jmoiron/sqlx"
	"sort"
)

const noVersion = "0.0.0"

type MigrationHistoryRepository struct {
}

func (repo *MigrationHistoryRepository) CurrentVersion(DB *sqlx.DB) (version string, err error) {
	var mh models.MigrationHistory
	DB.Select(&mh, "SELECT * FROM migration_history ORDER BY id DESC ")
	if mh.Version == "" {
		return noVersion, nil
	}
	return mh.Version, err
}
func (repo *MigrationHistoryRepository) ApplyMigrations(DB *sqlx.DB, collection version.Collection, migrates map[string]string) (err error) {
	// 对版本进行排序
	sort.Sort(collection)
	// 开启事务
	tx := DB.MustBegin()

	for _, v := range collection {
		tx.MustExec(migrates[v.String()])
	}
	for _, v := range collection {
		tx.MustExec("insert into migration_history (version, data ) values ($1,$2)", v.String(), migrates[v.String()])
	}
	tx.Commit()
	return

}
func (repo *MigrationHistoryRepository) ListMigration(DB *sqlx.DB) (migrates []models.MigrationHistory, count int, err error) {
	err = DB.Select(&migrates, "SELECT * FROM migration_history")
	count = len(migrates)
	return
}
