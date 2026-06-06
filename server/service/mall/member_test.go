package mall

import (
	"testing"

	mallReq "mall-admin/server/model/mall/request"
)

func TestRegisterInvalidMobile(t *testing.T) {
	s := MemberService{}
	_, err := s.Register(mallReq.MemberRegister{
		Mobile:   "12345",
		Password: "abc12345",
	}, "127.0.0.1")
	if err == nil {
		t.Fatal("expected error for invalid mobile")
	}
	if err.Error() != "手机号格式不正确" {
		t.Fatalf("unexpected error: %v", err)
	}
}
