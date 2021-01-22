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

func (t Tag) Get(c *gin.Context)    {}
func (t Tag) List(c *gin.Context)   {}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
