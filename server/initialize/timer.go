// timer.go 注册后台定时任务（如每日清理过期日志）。
package initialize

import (
	"mall-admin/server/task"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"mall-admin/server/global"
)

// Timer 启动定时任务调度器并注册内置清理任务。
func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // 定时任务方法定在task文件包中
			if err != nil {
				global.GVA_LOG.Error("定时清理数据库失败", zap.Error(err))
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			global.GVA_LOG.Error("注册定时任务失败", zap.Error(err))
		}

		// 其他定时任务定在这里 参考上方使用方法

		//_, err := global.GVA_Timer.AddTaskByFunc("定时任务标识", "corn表达式", func() {
		//	具体执行内容...
		//  ......
		//}, option...)
		//if err != nil {
		//	fmt.Println("add timer error:", err)
		//}
	}()
}
