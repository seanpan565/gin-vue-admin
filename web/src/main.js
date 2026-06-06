// 应用入口：初始化 Vue 实例并挂载全局插件
import './style/element_visiable.scss'
import 'element-plus/theme-chalk/dark/css-vars.css'
import 'uno.css'
import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import { setupVueRootValidator } from 'vite-check-multiple-dom/client';

import 'element-plus/dist/index.css'
// 引入商城管理后台初始化
import './core/mall-admin'
// 引入封装的router
import router from '@/router/index'
import '@/permission'
import run from '@/core/mall-admin.js'
import auth from '@/directive/auth'
import clickOutSide from '@/directive/clickOutSide'
import { store } from '@/pinia'
import App from './App.vue'
import '@/core/error-handel'

const app = createApp(App)

app.config.productionTip = false

setupVueRootValidator(app, {
    lang: 'zh'
  })

app
  .use(run)
  .use(ElementPlus)
  .use(store)
  .use(auth)
  .use(clickOutSide)
  .use(router)
  .mount('#app')
export default app
