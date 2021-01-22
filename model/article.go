/**
 * @Author: hqd
 * @Description: article model
 * @File:  article
 * @Version: 1.0.0
 * @Date: 2021/1/17 16:47
 */

package model

type Article struct {
	*Model
	Title       string
	Description string
	Content     string
	CoverImgURL string
	Visits      int
	Status      uint
	Sort        int
}

func (a Article) TableName() string {
	return "blog_article"
}
