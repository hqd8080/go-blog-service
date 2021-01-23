/**
 * @Author: hqd
 * @Description: base model
 * @File:  model
 * @Version: 1.0.0
 * @Date: 2021/1/17 16:35
 */

package model

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hqd8080/go-blog-service/conf"
	"github.com/jinzhu/gorm"
)

var dbMaster *gorm.DB
var dbSlave *gorm.DB

// base model
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	IsDeleted uint       `json:"is_deleted"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func Master() *gorm.DB {
	return dbMaster
}

func Slave() *gorm.DB {
	return dbSlave
}

func setDBConnection(master, slave *gorm.DB) {
	dbMaster = master
	dbSlave = slave
}

func NewDBConnection(MysqlConf *conf.MysqlConf) error {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		MysqlConf.User,
		MysqlConf.Password,
		MysqlConf.Host,
		MysqlConf.DBName,
		MysqlConf.Charset,
		MysqlConf.ParseTime,
	))
	if err != nil {
		return err
	}
	if conf.ServerConfig.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(MysqlConf.MaxIdleConns)
	db.DB().SetMaxOpenConns(MysqlConf.MaxOpenConns)
	setDBConnection(db, db)
	return nil
}
