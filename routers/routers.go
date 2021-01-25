/**
 * @Author: hqd
 * @Description: router
 * @File: routers
 * @Version: 1.0.0
 * @Date: 2021/1/17 16:53
 */

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hqd8080/go-blog-service/middleware"
	v1 "github.com/hqd8080/go-blog-service/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	api := r.Group("api/v1")
	{
		tag := v1.NewTag()
		article := v1.NewArticle()

		api.POST("/article", article.Create)            // 新增文章
		api.DELETE("/article/:id", article.Delete)      // 删除文章
		api.PUT("/article/:id", article.Update)         // 更新文章
		api.PATCH("/article/:id/state", article.Update) // 更新状态
		api.GET("/articles", article.List)              // 文章列表

		api.POST("/tags", tag.Create)            // 新增标签
		api.DELETE("/tags/:id", tag.Delete)      // 删除标签
		api.PUT("/tags/:id", tag.Update)         // 更新标签
		api.PATCH("/tags/:id/state", tag.Update) // 更新状态
		api.GET("/tags", tag.List)               // 获取标签列表

		api.POST("/website")
		api.DELETE("/website/:id")
		api.PUT("/website/:id")
		api.PATCH("/website/:id/state")
		api.GET("/websites")
	}
	return r
}
