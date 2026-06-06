// 商城会员管理 API（后台）
import service from '@/utils/request'

export const getMemberList = (data) => {
  return service({
    url: '/mall/member/list',
    method: 'post',
    data
  })
}

export const getMemberDetail = (data) => {
  return service({
    url: '/mall/member/detail',
    method: 'post',
    data
  })
}
