package base

type RoleDTO struct {
	RoleUuid string `json:"role_uuid"   orm:"role_uuid"   description:"角色UUID"` // 角色UUID
	Name     string `json:"name"        orm:"name"        description:"角色名称"`   // 角色名称
	Code     string `json:"code"        orm:"code"        description:"角色编码"`   // 角色编码
}
