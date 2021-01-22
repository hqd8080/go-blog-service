/**
 * @Author: hqd
 * @Description: 文章
 * @File:  article
 * @Version: 1.0.0
 * @Date: 2021/1/17 17:08
 */

package v1

import "github.com/gin-gonic/gin"

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context)    {}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
