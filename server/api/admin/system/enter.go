// Package system 系统管理模块API
package system

import "mall-admin/server/service"

// ApiGroup 系统模块API分组
type ApiGroup struct {
	DBApi
	JwtApi
	BaseApi
	SystemApi
	CasbinApi
	SystemApiApi
	AuthorityApi
	DictionaryApi
	AuthorityMenuApi
	OperationRecordApi
	DictionaryDetailApi
	AuthorityBtnApi
	SysParamsApi
	SysVersionApi
	SysErrorApi
	LoginLogApi
	ApiTokenApi
}

// 注入 service 层依赖
var (
	apiService              = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService              = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService             = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService           = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	casbinService           = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	baseMenuService         = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	authorityService        = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	authorityBtnService     = service.ServiceGroupApp.SystemServiceGroup.AuthorityBtnService
	systemConfigService     = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	sysParamsService        = service.ServiceGroupApp.SystemServiceGroup.SysParamsService
	operationRecordService  = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
	dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	sysVersionService       = service.ServiceGroupApp.SystemServiceGroup.SysVersionService
	sysErrorService         = service.ServiceGroupApp.SystemServiceGroup.SysErrorService
	loginLogService         = service.ServiceGroupApp.SystemServiceGroup.LoginLogService
	apiTokenService         = service.ServiceGroupApp.SystemServiceGroup.ApiTokenService
)
