/**
 * @Author: hqd
 * @Description: tag
 * @File:  tag
 * @Version: 1.0.0
 * @Date: 2021/1/25 18:20
 */

package service

type TagListRequest struct {
	Name  string `form:"name" binding:"max:100"`
	State uint8  `form:"state,default=1" binding:"oneOf=0 1"`
}

type CreateTagRequest struct {
	Name  string `form:"name" binding:"required,min=3,max=100"`
	State uint8  `form:"state,default=1" binding:"oneOf=0 1"`
}
