// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/quarkcloudio/quark-go/v2/pkg/dal"
	"github.com/quarkcloudio/quark-micro/cmd/content/api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库配置
func dbInit() {

	// 配置信息
	var (
		dbUser     = config.Mysql.Username
		dbPassword = config.Mysql.Password
		dbHost     = config.Mysql.Host
		dbPort     = config.Mysql.Port
		dbName     = config.Mysql.Database
		dbCharset  = config.Mysql.Charset
	)

	// 数据库配置信息
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&parseTime=True&loc=Local"

	dal.InitDB(mysql.Open(dsn), &gorm.Config{})
}

func main() {

	// 初始化数据库
	dbInit()

	// 初始化服务
	h := server.Default(server.WithHostPorts(config.Service.Host))

	register(h)
	h.Spin()
}
