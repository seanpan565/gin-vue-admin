// System模块Service层：系统管理核心业务逻辑
package system

// ServiceGroup 系统模块Service聚合
type ServiceGroup struct {
	JwtService
	ApiService
	MenuService
	UserService
	CasbinService
	InitDBService
	BaseMenuService
	AuthorityService
	DictionaryService
	SystemConfigService
	OperationRecordService
	DictionaryDetailService
	AuthorityBtnService
	SysParamsService
	SysVersionService
	SysErrorService
	LoginLogService
	ApiTokenService
}
