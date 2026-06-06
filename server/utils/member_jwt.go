// member_jwt.go C 端会员 JWT 签发与解析工具。
package utils

import (
	"errors"
	"time"

	"mall-admin/server/global"
	mallReq "mall-admin/server/model/mall/request"
	jwt "github.com/golang-jwt/jwt/v5"
)

const memberTokenAudience = "GVA-MALL"

// MemberJWT C 端会员 JWT 管理器，与后台 JWT 共用 SigningKey，通过 Audience 区分。
type MemberJWT struct {
	SigningKey []byte
}

// NewMemberJWT 创建会员 JWT 管理器。
func NewMemberJWT() *MemberJWT {
	return &MemberJWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}
}

// CreateMemberClaims 构建会员 JWT Claims。
func (j *MemberJWT) CreateMemberClaims(base mallReq.MemberBaseClaims) mallReq.MemberCustomClaims {
	bf, _ := ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	return mallReq.MemberCustomClaims{
		MemberBaseClaims: base,
		BufferTime:       int64(bf / time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{memberTokenAudience},
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),
			Issuer:    global.GVA_CONFIG.JWT.Issuer,
		},
	}
}

// CreateMemberToken 签发会员 token。
func (j *MemberJWT) CreateMemberToken(claims mallReq.MemberCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseMemberToken 解析会员 token。
func (j *MemberJWT) ParseMemberToken(tokenString string) (*mallReq.MemberCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &mallReq.MemberCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, TokenExpired
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, TokenMalformed
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, TokenSignatureInvalid
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, TokenNotValidYet
		default:
			return nil, TokenInvalid
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*mallReq.MemberCustomClaims); ok && token.Valid {
			if !memberAudienceMatch(claims, memberTokenAudience) {
				return nil, TokenInvalid
			}
			return claims, nil
		}
	}
	return nil, TokenValid
}

func memberAudienceMatch(claims *mallReq.MemberCustomClaims, audience string) bool {
	for _, aud := range claims.Audience {
		if aud == audience {
			return true
		}
	}
	return false
}

// MemberLoginToken 为会员生成登录 token。
func MemberLoginToken(base mallReq.MemberBaseClaims) (token string, claims mallReq.MemberCustomClaims, err error) {
	j := NewMemberJWT()
	claims = j.CreateMemberClaims(base)
	token, err = j.CreateMemberToken(claims)
	return
}
