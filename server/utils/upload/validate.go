// validate.go 上传文件类型白名单校验。
package upload

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"mall-admin/server/global"
)

var defaultAllowedExts = map[string]struct{}{
	".jpg": {}, ".jpeg": {}, ".png": {}, ".gif": {}, ".webp": {}, ".bmp": {}, ".svg": {},
	".pdf": {}, ".doc": {}, ".docx": {}, ".xls": {}, ".xlsx": {}, ".ppt": {}, ".pptx": {},
	".txt": {}, ".csv": {}, ".zip": {}, ".rar": {}, ".7z": {},
	".mp3": {}, ".mp4": {}, ".wav": {}, ".avi": {}, ".mov": {}, ".mkv": {},
}

// ValidateUploadFile 校验上传文件扩展名是否在白名单内。
func ValidateUploadFile(file *multipart.FileHeader) error {
	if file == nil {
		return fmt.Errorf("文件不能为空")
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext == "" {
		return fmt.Errorf("文件缺少扩展名")
	}
	allowed := allowedExtSet()
	if _, ok := allowed[ext]; !ok {
		return fmt.Errorf("不允许上传该类型文件: %s", ext)
	}
	dangerous := []string{".php", ".jsp", ".asp", ".aspx", ".exe", ".sh", ".bat", ".cmd", ".ps1", ".html", ".htm", ".js", ".mjs", ".cjs"}
	for _, d := range dangerous {
		if ext == d {
			return fmt.Errorf("不允许上传可执行或脚本类文件")
		}
	}
	return nil
}

func allowedExtSet() map[string]struct{} {
	exts := global.GVA_CONFIG.Security.UploadAllowedExts
	if len(exts) == 0 {
		return defaultAllowedExts
	}
	set := make(map[string]struct{}, len(exts))
	for _, e := range exts {
		e = strings.ToLower(strings.TrimSpace(e))
		if e == "" {
			continue
		}
		if !strings.HasPrefix(e, ".") {
			e = "." + e
		}
		set[e] = struct{}{}
	}
	return set
}
