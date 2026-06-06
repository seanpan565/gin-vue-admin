/** 网站全局配置：应用名、Tab 缓存开关及启动日志 */
import packageInfo from '../../package.json'

import { externalLinks } from './site'

const greenText = (text) => `\x1b[32m${text}\x1b[0m`

export const config = {
  appName: '商城管理后台',
  showViteLogo: true,
  keepAliveTabs: false,
  logs: []
}

export const viteLogo = (env) => {
  if (!config.showViteLogo) {
    return
  }

  const appName = env.VITE_APP_TITLE || config.appName

  console.log(greenText(`> ${appName} v${packageInfo.version}`))
  console.log(greenText(`> 前端: http://127.0.0.1:${env.VITE_CLI_PORT}`))
  console.log(greenText(`> 后端: http://127.0.0.1:${env.VITE_SERVER_PORT}`))
  console.log(greenText(`> Swagger: ${externalLinks.swagger(env.VITE_SERVER_PORT)}`))
  console.log('\n')
}

export default config
