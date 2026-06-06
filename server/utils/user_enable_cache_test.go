package utils

import (
	"testing"

	"mall-admin/server/global"
)

func TestInvalidateUserEnableCache(t *testing.T) {
	global.BlackCache.Set(UserEnableCachePrefix+"test-uuid", true, 0)
	InvalidateUserEnableCache("test-uuid")
	if _, ok := global.BlackCache.Get(UserEnableCachePrefix + "test-uuid"); ok {
		t.Fatal("expected user enable cache to be cleared")
	}
}

func TestInvalidateUserEnableCacheEmptyUUID(t *testing.T) {
	InvalidateUserEnableCache("")
}
