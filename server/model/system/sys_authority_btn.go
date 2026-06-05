package system

// SysAuthorityBtn 角色菜单按钮权限表模型
type SysAuthorityBtn struct {
	AuthorityId      uint           `gorm:"comment:角色ID"`
	SysMenuID        uint           `gorm:"comment:菜单ID"`
	SysBaseMenuBtnID uint           `gorm:"comment:菜单按钮ID"`
	SysBaseMenuBtn   SysBaseMenuBtn ` gorm:"comment:按钮详情"`
}
