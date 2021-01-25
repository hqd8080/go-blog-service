/**
 * @Author: hqd
 * @Description: 分页
 * @File:  pagination
 * @Version: 1.0.0
 * @Date: 2021/1/23 13:17
 */

package common

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hqd8080/go-blog-service/conf"
)

func GetPage(c *gin.Context) int {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize <= 0 {
		return conf.AppConfig.DefaultPageSize
	}
	if pageSize > conf.AppConfig.MaxPageSize {
		return conf.AppConfig.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
