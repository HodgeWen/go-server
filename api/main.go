/*
 * @Author: whj
 * @Date: 2023-01-17 16:16:18
 * @LastEditors: whj
 * @LastEditTime: 2023-01-18 15:30:23
 * @FilePath: /test-go/api/main.go
 *
 */
package api

import (
	// 改变包的地址从而更改包的版本
	handler "test-go/api/v1"

	"github.com/gin-gonic/gin"
)

var server = gin.Default()

const VER = "api/v1"

var Version = server.Group(VER)

func LoadRouter(moduleName string, handler handler.Api) {
	module := Version.Group(moduleName)

	if handler.Page != nil {
		module.GET("page", handler.Page)
	}
	if handler.List != nil {
		module.GET("list", handler.List)
	}
	if handler.Delete != nil {
		module.DELETE(":id", handler.Delete)
	}
}

func init() {
	LoadRouter("auth", handler.Auth)

	server.Run(":7777")
}
