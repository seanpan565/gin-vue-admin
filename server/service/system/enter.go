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
	AutoCodeService
	BaseMenuService
	AuthorityService
	DictionaryService
	SystemConfigService
	OperationRecordService
	DictionaryDetailService
	AuthorityBtnService
	SysExportTemplateService
	SysParamsService
	SysVersionService
	SkillsService
	AIWorkflowSession aiWorkflowSession
	AutoCodePlugin    autoCodePlugin
	AutoCodePackage   autoCodePackage
	AutoCodeHistory   autoCodeHistory
	AutoCodeTemplate  autoCodeTemplate
	SysErrorService
	LoginLogService
	ApiTokenService
}
