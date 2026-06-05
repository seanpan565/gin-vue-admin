// interfaces_base.go 提供 AST 处理基类，封装文件解析、格式化及路径转换。
package ast

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/pkg/errors"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Base AST 处理基类，持有 FileSet 供解析与格式化使用。
type Base struct {
	FileSet *token.FileSet
}

// Parse 解析 Go 源文件为 AST，writer 非空时从 writer 读取内容。
func (a *Base) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
	a.FileSet = token.NewFileSet()
	if writer != nil {
		file, err = parser.ParseFile(a.FileSet, filename, nil, parser.ParseComments)
	} else {
		file, err = parser.ParseFile(a.FileSet, filename, writer, parser.ParseComments)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "[filepath:%s]打开/解析文件失败!", filename)
	}
	return file, nil
}

// Rollback 基类默认不回滚，由子类覆盖实现。
func (a *Base) Rollback(file *ast.File) error {
	return nil
}

// Injection 基类默认不注入，由子类覆盖实现。
func (a *Base) Injection(file *ast.File) error {
	return nil
}

// Format 将 AST 格式化后写回文件或 writer。
func (a *Base) Format(filename string, writer io.Writer, file *ast.File) error {
	fileSet := a.FileSet
	if fileSet == nil {
		fileSet = token.NewFileSet()
	}
	if writer == nil {
		open, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666)
		defer open.Close()
		if err != nil {
			return errors.Wrapf(err, "[filepath:%s]打开文件失败!", filename)
		}
		writer = open
	}
	err := format.Node(writer, fileSet, file)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]注入失败!", filename)
	}
	return nil
}

// RelativePath 绝对路径转相对路径
func (a *Base) RelativePath(filePath string) string {
	server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
	hasServer := strings.Index(filePath, server)
	if hasServer != -1 {
		filePath = strings.TrimPrefix(filePath, server)
		keys := strings.Split(filePath, string(filepath.Separator))
		filePath = path.Join(keys...)
	}
	return filePath
}

// AbsolutePath 相对路径转绝对路径
func (a *Base) AbsolutePath(filePath string) string {
	server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
	keys := strings.Split(filePath, "/")
	filePath = filepath.Join(keys...)
	filePath = filepath.Join(server, filePath)
	return filePath
}
