// member_jwt.go C 端会员 JWT Claims 请求结构。
package request

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// MemberBaseClaims C 端会员 JWT 载荷。
type MemberBaseClaims struct {
	UUID     uuid.UUID
	ID       uint
	Mobile   string
	Nickname string
}

// MemberCustomClaims C 端会员 JWT 完整结构，Audience 为 GVA-MALL，与后台 token 区分。
type MemberCustomClaims struct {
	MemberBaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}
