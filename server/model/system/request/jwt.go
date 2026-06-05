// Package request 定义系统模块 API 请求参数结构。
package request

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// CustomClaims JWT 自定义载荷结构
// CustomClaims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}
