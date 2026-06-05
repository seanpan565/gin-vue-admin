// 插件市场 API：获取可安装插件列表
import service from '@/utils/request'

export const getShopPluginList = (params) => {
  return service({
    baseURL: "plugin",
    url: '/shopPlugin/getShopPluginList',
    method: 'get',
    params
  })
}