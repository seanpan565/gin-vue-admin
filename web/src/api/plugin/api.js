// 插件市场 API（默认关闭，见 core/site.js 中 siteFeatures.pluginMarket）
import service from '@/utils/request'

export const getShopPluginList = (params) => {
  return service({
    baseURL: 'plugin',
    url: '/shopPlugin/getShopPluginList',
    method: 'get',
    params
  })
}
