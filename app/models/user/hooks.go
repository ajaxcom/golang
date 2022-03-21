package user

import (
	"goblog/pkg/password"

	"gorm.io/gorm"
)

// BeforCreate GORM的模型钩子，创建模型前调用
func (user *User) BeforCreate(tx *gorm.DB) (err error) {
	user.Password = password.Hash(user.Password)
	return
}

// BeforeSave GORM 的模型钩子，在保存和更新模型前调用
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if !password.IsHash(user.Password) {
		user.Password = password.Hash(user.Password)
	}
	return
}
