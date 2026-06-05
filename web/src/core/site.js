/**
 * 站点级开关与外链配置（本地商城开发可在此统一调整）
 */
export const siteFeatures = {
  /** 是否展示插件市场相关入口与代理 */
  pluginMarket: false,
  /** 是否展示商业授权/购买类提示 */
  commercialPrompts: false
}

export const externalLinks = {
  github: 'https://github.com/flipped-aurora/gin-vue-admin',
  gvaDocs: 'https://www.gin-vue-admin.com/',
  swagger: (port) => `http://127.0.0.1:${port}/swagger/index.html`,
  vueDocs: 'https://cn.vuejs.org/',
  ginDocs: 'https://gin-gonic.com/zh-cn/docs/'
}
