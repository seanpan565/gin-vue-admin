// interfaces.go 定义 AST 处理器统一接口，供各类代码注入器实现。
package ast

import (
	"go/ast"
	"io"
)

// Ast AST 处理器接口，统一解析、注入、回滚与格式化流程。
type Ast interface {
	// Parse 解析文件/代码
	Parse(filename string, writer io.Writer) (file *ast.File, err error)
	// Rollback 回滚
	Rollback(file *ast.File) error
	// Injection 注入
	Injection(file *ast.File) error
	// Format 格式化输出
	Format(filename string, writer io.Writer, file *ast.File) error
}
