// Package mcpTool 字典相关上游 HTTP 调用封装。
package mcpTool

// dictionary_http.go 字典 CRUD 上游 API 代理。

import (
	"context"
	"net/url"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

// exportedDictionary 导出的字典结构。
type exportedDictionary struct {
	Name                 string                       `json:"name"`
	Type                 string                       `json:"type"`
	Status               *bool                        `json:"status"`
	Desc                 string                       `json:"desc"`
	SysDictionaryDetails []system.SysDictionaryDetail `json:"sysDictionaryDetails"`
}

func fetchDictionaryList(ctx context.Context, keyword string) ([]system.SysDictionary, error) {
	query := url.Values{}
	if keyword != "" {
		query.Set("name", keyword)
	}

	resp, err := getUpstream[[]system.SysDictionary](ctx, "/sysDictionary/getSysDictionaryList", query)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func findDictionaryByType(ctx context.Context, dictType string) (*system.SysDictionary, error) {
	dictionaries, err := fetchDictionaryList(ctx, dictType)
	if err != nil {
		return nil, err
	}

	for _, dictionary := range dictionaries {
		if dictionary.Type == dictType {
			dict := dictionary
			return &dict, nil
		}
	}

	return nil, nil
}

func exportDictionary(ctx context.Context, id uint) (*exportedDictionary, error) {
	query := url.Values{}
	query.Set("id", strconv.FormatUint(uint64(id), 10))

	resp, err := getUpstream[exportedDictionary](ctx, "/sysDictionary/exportSysDictionary", query)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func createDictionary(ctx context.Context, dictionary system.SysDictionary) error {
	_, err := postUpstream[map[string]any](ctx, "/sysDictionary/createSysDictionary", dictionary)
	return err
}

func createDictionaryDetail(ctx context.Context, detail system.SysDictionaryDetail) error {
	_, err := postUpstream[map[string]any](ctx, "/sysDictionaryDetail/createSysDictionaryDetail", detail)
	return err
}

func enabledBoolPointer() *bool {
	return utils.Pointer(true)
}
