// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package models

const noVersion = "0.0.0"

// MigrationHistory is model
type MigrationHistory struct {
	ID          int64
	Version     string
	DateApplied int64
}

// TableName returns name of table
func (mh *MigrationHistory) TableName() string {
	return "migration_history"
}

/*
// CurrentVersion returns current version of database migrations
func (mh *MigrationHistory) CurrentVersion() (string, error) {


	err := DBConn.Last(mh).Error

	if mh.Version == "" {
		return noVersion, nil
	}

	return mh.Version, err
}

// ApplyMigration executes database schema and writes migration history
func (mh *MigrationHistory) ApplyMigration(version string, query string) (err error) {
	tx := DBConn.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
		DBConn.Close()
	}()

	err = tx.Exec(query).Error

	return DBConn.Create(&MigrationHistory{Version: version, DateApplied: time.Now().Unix()}).Error
}
*/
