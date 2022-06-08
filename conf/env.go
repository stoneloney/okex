package conf

import "flag"

var (
	SwitchEnv = flag.String("env", "develop", "rt")
)

// GetEnv 获取环境
func GetEnv() string {
	return *SwitchEnv
}

//---------------------------------------------- 环境变量定义 ----------------------------------------------//
const (
	EnvRelease = "release"   //正式环境
	EnvDevelop = "develop"   //开发联调环境
)
