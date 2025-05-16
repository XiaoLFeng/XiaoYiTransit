package dto

// UserInfo 表示用户的基础信息。
//
// 包含用户的唯一标识、账号信息、联系方式和状态等信息。
type UserInfo struct {
	Id       string `json:"id"              orm:"id"              description:"用户ID"`
	Username string `json:"username"        orm:"username"        description:"用户名"`
	Password string `json:"password"        orm:"password"        description:"密码"`
	RealName string `json:"real_name"       orm:"real_name"       description:"真实姓名"`
	Email    string `json:"email"           orm:"email"           description:"邮箱"`
	Phone    string `json:"phone"           orm:"phone"           description:"手机号"`
	Avatar   string `json:"avatar"          orm:"avatar"          description:"头像"`
	RoleId   string `json:"role_id"         orm:"role_id"         description:"角色ID"`
	Status   int    `json:"status"          orm:"status"          description:"状态: 0-禁用, 1-启用"`
}
