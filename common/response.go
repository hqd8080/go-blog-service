/**
 * @Author: hqd
 * @Description: 响应处理
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2021/1/23 13:30
 */

package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (resp *Response) ResponseData(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	resp.Ctx.JSON(http.StatusOK, data)
}

func (resp *Response) ResponseList(list interface{}, total int) {
	resp.Ctx.JSON(
		http.StatusOK,
		gin.H{
			"list":  list,
			"pager": Pager{Page: GetPage(resp.Ctx), PageSize: GetPageSize(resp.Ctx), Total: total},
		},
	)
}

func (resp *Response) ResponseError(code int, status int, msg string) {
	data := gin.H{"code": code, "msg": msg}
	resp.Ctx.JSON(status, data)
}
