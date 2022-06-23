package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"imall/global"
	"time"
)

// Mysql 配置MySQl数据库
func Mysql() {
	m := global.Config.Mysql
	var dsn = fmt.Sprintf("%s:%s@%s", m.Username, m.Password, m.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		return
	}
	sqlDb, err := db.DB()
	if err != nil {
		fmt.Printf("mysql error: %s", err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	global.Db = db
}
