/**
 * @Author: hqd
 * @Description: tag model
 * @File:  tag
 * @Version: 1.0.0
 * @Date: 2021/1/17 16:40
 */

package model

type Tag struct {
	*Model
	TagName   string `json:"tag_name"`
	TagStatus uint   `json:"tag_status"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
