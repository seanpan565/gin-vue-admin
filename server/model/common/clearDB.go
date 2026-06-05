package common

// ClearDB 定时清理数据库任务配置。
type ClearDB struct {
	TableName    string
	CompareField string
	Interval     string
}
