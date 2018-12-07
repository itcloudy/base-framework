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
	*sqlx.DB
}

func (repo *MigrationHistoryRepository) CurrentVersion() (version string, err error) {
	var mh models.MigrationHistory
	repo.Select(&mh, "SELECT * FROM migration_history ORDER BY id DESC ")
	if mh.Version == "" {
		return noVersion, nil
	}
	return mh.Version, err
}
func (repo *MigrationHistoryRepository) ApplyMigrations(collection version.Collection, migrates map[string]string) (err error) {
	// 对版本进行排序
	sort.Sort(collection)
	// 开启事务
	tx := repo.MustBegin()

	for _, v := range collection {
		tx.MustExec(migrates[v.String()])
	}
	for _, v := range collection {
		tx.MustExec("intersect into migration_history (version, data ) values ($1,$2)", v.String(), migrates[v.String()])
	}

	tx.Commit()
	return

}
func (repo *MigrationHistoryRepository) ListMigration() (migrates []models.MigrationHistory, err error) {
	err = repo.Select(&migrates, "SELECT * FROM migration_history")
	return
}
