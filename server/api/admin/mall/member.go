// member.go 商城后台会员管理 API 处理器。
package mall

import (
	"mall-admin/server/global"
	"mall-admin/server/model/common/response"
	mallReq "mall-admin/server/model/mall/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// MemberApi 后台会员管理接口。
type MemberApi struct{}

// GetMemberList
// @Tags      MallMember
// @Summary   会员列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      mallReq.MemberSearch  true  "分页及筛选"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}
// @Router    /mall/member/list [post]
func (a *MemberApi) GetMemberList(c *gin.Context) {
	var pageInfo mallReq.MemberSearch
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := memberService.GetMemberList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取会员列表失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	for i := range list {
		list[i].Password = ""
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetMemberDetail
// @Tags      MallMember
// @Summary   会员详情
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      mallReq.MemberByID  true  "会员ID"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /mall/member/detail [post]
func (a *MemberApi) GetMemberDetail(c *gin.Context) {
	var req mallReq.MemberByID
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	member, err := memberService.GetMemberByID(req.ID)
	if err != nil {
		response.FailWithMessage("会员不存在", c)
		return
	}
	member.Password = ""

	response.OkWithDetailed(member, "获取成功", c)
}
