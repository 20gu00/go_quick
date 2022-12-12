package mysql

import (
	"fmt"
	"go_forum/common/setUp/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sqlx.DB
var gormDB *gorm.DB

// sqlx
func InitMysql(cfg *config.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		cfg.UserName,
		cfg.MysqlPassword,
		cfg.MysqlAddr,
		cfg.MysqlPort,
		cfg.DBName,
	)
	//viper.GetString("mysql.user_name"),
	//viper.GetString("mysql.password"),
	//viper.GetString("mysql.addr"),
	//viper.GetInt("mysql.port"),
	//viper.GetString("mysql.db_name"),
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.MaxConnection)
	db.SetMaxIdleConns(cfg.MaxIdle)
	return
}

//包内全局
func DBClose() {
	_ = db.Close()
}

// gorm
func InitMysqlGorm(cfg *config.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		cfg.UserName,
		cfg.MysqlPassword,
		cfg.MysqlAddr,
		cfg.MysqlPort,
		cfg.DBName,
	)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	fmt.Println(gormDB)
	return
}

// 不用close
