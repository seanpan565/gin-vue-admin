// Package ast 提供自动代码生成所需的 Go AST 解析、构造与修改工具。
//
// ast.go 定义 AST 通用辅助函数，用于菜单/API/字典结构体生成及 import 查询。
package ast

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

// AddImport 向 AST 节点追加 import 路径（已存在则跳过）。
func AddImport(astNode ast.Node, imp string) {
	impStr := fmt.Sprintf("\"%s\"", imp)
	ast.Inspect(astNode, func(node ast.Node) bool {
		if genDecl, ok := node.(*ast.GenDecl); ok {
			if genDecl.Tok == token.IMPORT {
				for i := range genDecl.Specs {
					if impNode, ok := genDecl.Specs[i].(*ast.ImportSpec); ok {
						if impNode.Path.Value == impStr {
							return false
						}
					}
				}
				genDecl.Specs = append(genDecl.Specs, &ast.ImportSpec{
					Path: &ast.BasicLit{
						Kind:  token.STRING,
						Value: impStr,
					},
				})
			}
		}
		return true
	})
}

// FindFunction 按函数名查找首个 FuncDecl 节点。
func FindFunction(astNode ast.Node, FunctionName string) *ast.FuncDecl {
	var funcDeclP *ast.FuncDecl
	ast.Inspect(astNode, func(node ast.Node) bool {
		if funcDecl, ok := node.(*ast.FuncDecl); ok {
			if funcDecl.Name.String() == FunctionName {
				funcDeclP = funcDecl
				return false
			}
		}
		return true
	})
	return funcDeclP
}

// FindArray 查找指定包名与类型名的数组字面量赋值节点。
func FindArray(astNode ast.Node, identName, selectorExprName string) *ast.CompositeLit {
	var assignStmt *ast.CompositeLit
	ast.Inspect(astNode, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			for _, expr := range node.Rhs {
				if exprType, ok := expr.(*ast.CompositeLit); ok {
					if arrayType, ok := exprType.Type.(*ast.ArrayType); ok {
						sel, ok1 := arrayType.Elt.(*ast.SelectorExpr)
						x, ok2 := sel.X.(*ast.Ident)
						if ok1 && ok2 && x.Name == identName && sel.Sel.Name == selectorExprName {
							assignStmt = exprType
							return false
						}
					}
				}
			}
		}
		return true
	})
	return assignStmt
}

// CreateMenuStructAst 将菜单配置转换为 AST 复合字面量表达式列表。
func CreateMenuStructAst(menus []system.SysBaseMenu) *[]ast.Expr {
	var menuElts []ast.Expr
	for i := range menus {
		elts := []ast.Expr{ // 结构体的字段
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "ParentId"},
				Value: &ast.BasicLit{Kind: token.INT, Value: "0"},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Path"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", menus[i].Path)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Name"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", menus[i].Name)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Hidden"},
				Value: &ast.Ident{Name: "false"},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Component"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", menus[i].Component)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Sort"},
				Value: &ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%d", menus[i].Sort)},
			},
			&ast.KeyValueExpr{
				Key: &ast.Ident{Name: "Meta"},
				Value: &ast.CompositeLit{
					Type: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "model"},
						Sel: &ast.Ident{Name: "Meta"},
					},
					Elts: []ast.Expr{
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Title"},
							Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", menus[i].Title)},
						},
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Icon"},
							Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", menus[i].Icon)},
						},
					},
				},
			},
		}

		// 添加菜单参数
		if len(menus[i].Parameters) > 0 {
			var paramElts []ast.Expr
			for _, param := range menus[i].Parameters {
				paramElts = append(paramElts, &ast.CompositeLit{
					Type: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "model"},
						Sel: &ast.Ident{Name: "SysBaseMenuParameter"},
					},
					Elts: []ast.Expr{
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Type"},
							Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", param.Type)},
						},
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Key"},
							Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", param.Key)},
						},
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Value"},
							Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", param.Value)},
						},
					},
				})
			}
			elts = append(elts, &ast.KeyValueExpr{
				Key: &ast.Ident{Name: "Parameters"},
				Value: &ast.CompositeLit{
					Type: &ast.ArrayType{
						Elt: &ast.SelectorExpr{
							X:   &ast.Ident{Name: "model"},
							Sel: &ast.Ident{Name: "SysBaseMenuParameter"},
						},
					},
					Elts: paramElts,
				},
			})
		}

		// 添加菜单按钮
		if len(menus[i].MenuBtn) > 0 {
			var btnElts []ast.Expr
			for _, btn := range menus[i].MenuBtn {
				btnElts = append(btnElts, &ast.CompositeLit{
					Type: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "model"},
						Sel: &ast.Ident{Name: "SysBaseMenuBtn"},
					},
					Elts: []ast.Expr{
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Name"},
							Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", btn.Name)},
						},
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Desc"},
							Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", btn.Desc)},
						},
					},
				})
			}
			elts = append(elts, &ast.KeyValueExpr{
				Key: &ast.Ident{Name: "MenuBtn"},
				Value: &ast.CompositeLit{
					Type: &ast.ArrayType{
						Elt: &ast.SelectorExpr{
							X:   &ast.Ident{Name: "model"},
							Sel: &ast.Ident{Name: "SysBaseMenuBtn"},
						},
					},
					Elts: btnElts,
				},
			})
		}

		menuElts = append(menuElts, &ast.CompositeLit{
			Type: nil,
			Elts: elts,
		})
	}
	return &menuElts
}

// CreateApiStructAst 将 API 配置转换为 AST 复合字面量表达式列表。
func CreateApiStructAst(apis []system.SysApi) *[]ast.Expr {
	var apiElts []ast.Expr
	for i := range apis {
		elts := []ast.Expr{ // 结构体的字段
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Path"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", apis[i].Path)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Description"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", apis[i].Description)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "ApiGroup"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", apis[i].ApiGroup)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Method"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", apis[i].Method)},
			},
		}
		apiElts = append(apiElts, &ast.CompositeLit{
			Type: nil,
			Elts: elts,
		})
	}
	return &apiElts
}

// CheckImport 检查文件是否已包含指定 import 路径。
func CheckImport(file *ast.File, importPath string) bool {
	for _, imp := range file.Imports {
		// Remove quotes around the import path
		path := imp.Path.Value[1 : len(imp.Path.Value)-1]

		if path == importPath {
			return true
		}
	}

	return false
}

func clearPosition(astNode ast.Node) {
	ast.Inspect(astNode, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.Ident:
			// 清除位置信息
			node.NamePos = token.NoPos
		case *ast.CallExpr:
			// 清除位置信息
			node.Lparen = token.NoPos
			node.Rparen = token.NoPos
		case *ast.BasicLit:
			// 清除位置信息
			node.ValuePos = token.NoPos
		case *ast.SelectorExpr:
			// 清除位置信息
			node.Sel.NamePos = token.NoPos
		case *ast.BinaryExpr:
			node.OpPos = token.NoPos
		case *ast.UnaryExpr:
			node.OpPos = token.NoPos
		case *ast.StarExpr:
			node.Star = token.NoPos
		}
		return true
	})
}

// CreateStmt 将 Go 表达式字符串解析为 ExprStmt 语句节点。
func CreateStmt(statement string) *ast.ExprStmt {
	expr, err := parser.ParseExpr(statement)
	if err != nil {
		log.Fatal(err)
	}
	clearPosition(expr)
	return &ast.ExprStmt{X: expr}
}

// IsBlockStmt 判断节点是否为 BlockStmt。
func IsBlockStmt(node ast.Node) bool {
	_, ok := node.(*ast.BlockStmt)
	return ok
}

// VariableExistsInBlock 检查代码块内是否已定义指定变量名。
func VariableExistsInBlock(block *ast.BlockStmt, varName string) bool {
	exists := false
	ast.Inspect(block, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			for _, expr := range node.Lhs {
				if ident, ok := expr.(*ast.Ident); ok && ident.Name == varName {
					exists = true
					return false
				}
			}
		}
		return true
	})
	return exists
}

// CreateDictionaryStructAst 将字典配置转换为 AST 复合字面量表达式列表。
func CreateDictionaryStructAst(dictionaries []system.SysDictionary) *[]ast.Expr {
	var dictElts []ast.Expr
	for i := range dictionaries {
		statusStr := "true"
		if dictionaries[i].Status != nil && !*dictionaries[i].Status {
			statusStr = "false"
		}

		elts := []ast.Expr{
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Name"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", dictionaries[i].Name)},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Type"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", dictionaries[i].Type)},
			},
			&ast.KeyValueExpr{
				Key: &ast.Ident{Name: "Status"},
				Value: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "utils"},
						Sel: &ast.Ident{Name: "Pointer"},
					},
					Args: []ast.Expr{
						&ast.Ident{Name: statusStr},
					},
				},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "Desc"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", dictionaries[i].Desc)},
			},
		}

		if len(dictionaries[i].SysDictionaryDetails) > 0 {
			var detailElts []ast.Expr
			for _, detail := range dictionaries[i].SysDictionaryDetails {
				detailStatusStr := "true"
				if detail.Status != nil && !*detail.Status {
					detailStatusStr = "false"
				}

				detailElts = append(detailElts, &ast.CompositeLit{
					Type: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "model"},
						Sel: &ast.Ident{Name: "SysDictionaryDetail"},
					},
					Elts: []ast.Expr{
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Label"},
							Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", detail.Label)},
						},
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Value"},
							Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", detail.Value)},
						},
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Extend"},
							Value: &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", detail.Extend)},
						},
						&ast.KeyValueExpr{
							Key: &ast.Ident{Name: "Status"},
							Value: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   &ast.Ident{Name: "utils"},
									Sel: &ast.Ident{Name: "Pointer"},
								},
								Args: []ast.Expr{
									&ast.Ident{Name: detailStatusStr},
								},
							},
						},
						&ast.KeyValueExpr{
							Key:   &ast.Ident{Name: "Sort"},
							Value: &ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%d", detail.Sort)},
						},
					},
				})
			}
			elts = append(elts, &ast.KeyValueExpr{
				Key: &ast.Ident{Name: "SysDictionaryDetails"},
				Value: &ast.CompositeLit{
					Type: &ast.ArrayType{Elt: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "model"},
						Sel: &ast.Ident{Name: "SysDictionaryDetail"},
					}},
					Elts: detailElts,
				},
			})
		}

		dictElts = append(dictElts, &ast.CompositeLit{
			Type: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "model"},
				Sel: &ast.Ident{Name: "SysDictionary"},
			},
			Elts: elts,
		})
	}
	return &dictElts
}
