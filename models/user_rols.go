package models

type userRols struct {
	// 超级管理员
	SuperAdmin int
	// 管理员
	Admin int
	// 编辑
	Edit int
	// 用户认证编辑
	UserEdit int
	// 普通用户
	User int
}

// UserRols is 用户角色
var UserRols = userRols{
	SuperAdmin: 100,
	Admin:      90,
	Edit:       80,
	UserEdit:   70,
	User:       0,
}
