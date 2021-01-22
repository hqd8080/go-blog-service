/**
 * @Author: hqd
 * @Description: base model
 * @File:  model
 * @Version: 1.0.0
 * @Date: 2021/1/17 16:35
 */

package model

import "time"

// base model
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	IsDeleted uint       `json:"is_deleted"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
