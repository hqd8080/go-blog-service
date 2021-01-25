/**
 * @Author: hqd
 * @Description: 标签
 * @File:  tag
 * @Version: 1.0.0
 * @Date: 2021/1/17 17:08
 */

package v1

import "github.com/gin-gonic/gin"

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context)  {}
func (t Tag) List(c *gin.Context) {}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称"
// @Param state body int false "标签状态"
// @Success 200 {object} error "成功"
// @Failure 400 {object} error "请求错误"
// @Failure 500 {object} error "服务器错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {

}

func (t Tag) Update(c *gin.Context) {}

func (t Tag) Delete(c *gin.Context) {}
