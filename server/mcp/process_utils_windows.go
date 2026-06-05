//go:build windows

// Package mcpTool Windows 平台 MCP 托管进程工具。
package mcpTool
// process_utils_windows.go Windows 下的进程分离与终止。

import (
	"errors"
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

const windowsStillActive = 259

// prepareDetachedProcess 以隐藏窗口方式启动子进程。
func prepareDetachedProcess(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x00000008 | 0x00000200 | 0x08000000,
	}
}

func processExists(pid int) bool {
	if pid <= 0 {
		return false
	}

	handle, err := windows.OpenProcess(windows.PROCESS_QUERY_LIMITED_INFORMATION, false, uint32(pid))
	if err != nil {
		return false
	}
	defer windows.CloseHandle(handle)

	var code uint32
	if err := windows.GetExitCodeProcess(handle, &code); err != nil {
		return false
	}

	return code == windowsStillActive
}

func terminateProcess(pid int) error {
	if pid <= 0 {
		return nil
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	err = process.Kill()
	if err == nil || errors.Is(err, os.ErrProcessDone) {
		return nil
	}

	return err
}
