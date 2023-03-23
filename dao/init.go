package dao

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mysql *gorm.DB

func init() {
	// init mysql
	var err error
	MysqlUsername := viper.GetString("mysql.username")
	MysqlPassword := viper.GetString("mysql.password")
	MysqlHost := viper.GetString("mysql.host")
	MysqlPort := viper.GetString("mysql.port")
	MysqlDatabase := viper.GetString("mysql.database")

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		MysqlUsername, MysqlPassword, MysqlHost, MysqlPort, MysqlDatabase)

	Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
	})
	if err != nil {
		panic(err)
	}

}
