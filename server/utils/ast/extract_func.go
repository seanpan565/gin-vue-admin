// extract_func.go 根据源码位置提取完整函数定义。
package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// ExtractFuncSourceByPosition 根据文件路径与行号，提取包含该行的整个方法源码
// 返回：方法名、完整源码、起止行号
func ExtractFuncSourceByPosition(filePath string, line int) (name string, source string, startLine int, endLine int, err error) {
	src, readErr := os.ReadFile(filePath)
	if readErr != nil {
		err = fmt.Errorf("read file failed: %w", readErr)
		return
	}

	fset := token.NewFileSet()
	file, parseErr := parser.ParseFile(fset, filePath, src, parser.ParseComments)
	if parseErr != nil {
		err = fmt.Errorf("parse file failed: %w", parseErr)
		return
	}

	var target *ast.FuncDecl
	ast.Inspect(file, func(n ast.Node) bool {
		fd, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}
		s := fset.Position(fd.Pos()).Line
		e := fset.Position(fd.End()).Line
		if line >= s && line <= e {
			target = fd
			startLine = s
			endLine = e
			return false
		}
		return true
	})

	if target == nil {
		err = fmt.Errorf("no function encloses line %d in %s", line, filePath)
		return
	}

	start := fset.Position(target.Pos()).Offset
	end := fset.Position(target.End()).Offset
	if start < 0 || end > len(src) || start >= end {
		err = fmt.Errorf("invalid offsets for function: start=%d end=%d len=%d", start, end, len(src))
		return
	}
	source = string(src[start:end])
	name = target.Name.Name
	return
}
