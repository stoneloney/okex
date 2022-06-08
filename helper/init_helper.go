package helper

import (
	"flag"
	"fmt"
	grom "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"okex/conf"
	"sync"
)

var (
	_dbIns  *gorm.DB

	lock sync.Mutex
)

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}

	lock.Lock()
	defer lock.Unlock()

	// 连接数据库
	mysqlCfg, err := conf.GetMysql("db")
	if err != nil {
		fmt.Println("GetMysql error:", err.Error())
	} else {
		dsn2 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlCfg["user"], mysqlCfg["password"], mysqlCfg["host"], mysqlCfg["port"], mysqlCfg["dbname"])
		_dbIns, err = gorm.Open(grom.Open(dsn2), &gorm.Config{})
		if err != nil {
			fmt.Println("gorm open error:", err.Error())
		}
	}

}

func GetDb() *gorm.DB {
	return _dbIns
}
