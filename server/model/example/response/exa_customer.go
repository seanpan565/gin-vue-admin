package response

import "mall-admin/server/model/example"

// 客户管理相关响应结构
type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
