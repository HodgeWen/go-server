/*
 * @Author: whj emailwen@qq.com
 * @Date: 2023-01-17 14:29:28
 * @LastEditors: whj
 * @LastEditTime: 2023-01-18 15:43:53
 * @FilePath: /test-go/api/v1/auth.go
 * @Description:
 *
 * Copyright (c) 2023 by whj emailwen@qq.com, All Rights Reserved.
 */

package v1

import (
	"time"

	"github.com/gin-gonic/gin"
)

type AuthApi struct {
	Api
}

var Auth = Api{

	/**
	 * @description:
	 * @param {*gin.Context} ctx
	 * @return {*}
	 */
	Page: func(ctx *gin.Context) {
		time.Sleep(8 * time.Second)
		ctx.JSON(200, gin.H{
			"msg": "page ok",
		})
	},
	/**
	 * @description: dsf
	 * @param {*gin.Context} ctx 上下文
	 * @return {*}
	 */
	List: func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "list ok",
		})
	},
	Delete: func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": ctx.Param("id"),
		})
	},
}
