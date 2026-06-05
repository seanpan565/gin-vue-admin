// GVA 框架 Vue 插件：注册全局资源
import { register } from './global'
import config from './config'
import packageInfo from '../../package.json'
import { externalLinks } from './site'

export default {
  install: (app) => {
    register(app)
    console.log(
      `[${config.appName}] v${packageInfo.version} | 前端 :${import.meta.env.VITE_CLI_PORT} | API :${import.meta.env.VITE_SERVER_PORT} | ${externalLinks.gvaDocs}`
    )
  }
}
