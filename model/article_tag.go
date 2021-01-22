/**
 * @Author: hqd
 * @Description: article tag model
 * @File:  article_tag
 * @Version: 1.0.0
 * @Date: 2021/1/17 16:44
 */

package model

type ArticleTag struct {
	*Model
	ArticleID uint `json:"article_id"`
	TagID     uint `json:"tag_id"`
}

func (at ArticleTag) TableName() string {
	return "blog_article_tag"
}
