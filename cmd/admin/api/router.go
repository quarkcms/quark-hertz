// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "github.com/quarkcms/quark-hertz/cmd/admin/api/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {

	// 后台路由组
	rg := r.Group("/api/admin")
	rg.GET("/login/:resource", handler.Login)
	rg.GET("/:resource/index", handler.ResourceIndex)
}
