//go:build !windows

// Package mcpTool Unix 平台 MCP 托管进程工具。
package mcpTool
// process_utils_unix.go 非 Windows 下的进程分离与终止。

import (
	"errors"
	"os/exec"
	"syscall"
)

// prepareDetachedProcess 以独立进程组启动子进程。
func prepareDetachedProcess(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
}

func processExists(pid int) bool {
	if pid <= 0 {
		return false
	}

	err := syscall.Kill(pid, 0)
	return err == nil || errors.Is(err, syscall.EPERM)
}

func terminateProcess(pid int) error {
	if pid <= 0 {
		return nil
	}

	err := syscall.Kill(pid, syscall.SIGTERM)
	if err == nil || errors.Is(err, syscall.ESRCH) {
		return nil
	}

	return err
}
