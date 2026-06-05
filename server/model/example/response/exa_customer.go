package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/example"

// 客户管理相关响应结构
type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
