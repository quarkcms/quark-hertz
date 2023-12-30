// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/quarkcloudio/quark-go/v2/pkg/adapter/hertzadapter"
	admininstall "github.com/quarkcloudio/quark-go/v2/pkg/app/admin/install"
	adminmiddleware "github.com/quarkcloudio/quark-go/v2/pkg/app/admin/middleware"
	adminservice "github.com/quarkcloudio/quark-go/v2/pkg/app/admin/service"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
	appinstall "github.com/quarkcloudio/quark-micro/cmd/admin/biz/install"
	"github.com/quarkcloudio/quark-micro/cmd/admin/biz/service"
	"github.com/quarkcloudio/quark-micro/cmd/admin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	h := server.Default(server.WithHostPorts(config.App.Host))

	// 注册路由
	register(h)

	// 加载html模板
	h.LoadHTMLGlob("../../web/template/*")

	// 静态文件
	h.StaticFile("/admin/", "../../web/app/admin/index.html")

	// 静态文件目录
	fs := &app.FS{Root: "../../web/app", IndexNames: []string{"index.html"}}
	h.StaticFS("/", fs)

	// 配置信息
	var (
		appKey     = config.App.Key
		dbUser     = config.Mysql.Username
		dbPassword = config.Mysql.Password
		dbHost     = config.Mysql.Host
		dbPort     = config.Mysql.Port
		dbName     = config.Mysql.Database
		dbCharset  = config.Mysql.Charset
	)

	// 数据库配置信息
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&parseTime=True&loc=Local"

	// 配置资源
	getConfig := &builder.Config{
		AppKey: appKey,
		DBConfig: &builder.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
		Providers:  append(adminservice.Providers, service.Provider...),
		StaticPath: "../../web",
	}

	// 实例化对象
	b := builder.New(getConfig)

	// 初始化安装
	admininstall.Handle()

	// 初始化本项目数据库
	appinstall.Handle()

	// 中间件
	b.Use(adminmiddleware.Handle)

	// 适配hertz
	hertzadapter.Adapter(b, h)

	// 启动服务
	h.Spin()
}
