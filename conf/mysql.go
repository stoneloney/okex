package conf

import (
	"errors"
	"fmt"
)

// 数据库配置

var MysqlCfg = map[string]map[string]map[string]string{
	EnvDevelop: {
		"db":  {"host": "118.25.191.234", "port": "3306", "dbname": "okex", "user": "root", "password": "jia119110"},
	},

	EnvRelease: {
		"db":  {"host": "118.25.191.234", "port": "3306", "dbname": "okex", "user": "root", "password": "jia119110"},
	},
}

// GetMysql 获取mysql配置
func GetMysql(name string, extParams ...string) (map[string]string, error) {
	var env string
	if len(extParams) > 0 {
		env = extParams[0]
	} else {
		env = *SwitchEnv
	}

	if _, ok := MysqlCfg[env][name]; !ok {
		return map[string]string{}, errors.New(fmt.Sprintf("Env:%s, node:%s is not exists", env, name))
	}
	c := MysqlCfg[env][name]
	return c, nil
}
