package initalize

import (
	"GO-CRM-API-SHOPDEV/global"
	"GO-CRM-API-SHOPDEV/po"
	"fmt"
	"gorm.io/driver/mysql"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}
func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	fmt.Println("Connection String:", s)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("InitMysql initialization success")
	global.Mdb = db

	SetPool()
	migrateTables()
}
func SetPool() {
	sqlDb, err := global.Mdb.DB()
	m := global.Config.Mysql
	if err != nil {
		fmt.Printf("mysql error %s:", err)
	}
	// set config ~~ Pool
	sqlDb.SetConnMaxLifetime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

// gen table database
func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("Migrating tables error")
	}
}
