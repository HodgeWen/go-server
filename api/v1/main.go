/*
 * @Author: whj
 * @Date: 2023-01-18 14:31:59
 * @LastEditors: whj
 * @LastEditTime: 2023-01-18 14:35:41
 * @FilePath: /test-go/api/v1/main.go
 *
 */
package v1

import "github.com/gin-gonic/gin"

type Api struct {
	/**
	 * @description:
	 * @param {*gin.Context} ctx
	 * @return {*}
	 */
	Page func(ctx *gin.Context)
	/**
	 * @description:
	 * @param {*gin.Context} ctx
	 * @return {*}
	 */
	List func(ctx *gin.Context)
	/**
	 * @description:
	 * @param {*gin.Context} ctx
	 * @return {*}
	 */
	Delete func(ctx *gin.Context)

	Create func(ctx *gin.Context)

	Update func(ctx *gin.Context)

	Find func(ctx *gin.Context)
}
