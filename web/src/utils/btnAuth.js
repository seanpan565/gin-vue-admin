// 按钮权限 composable：读取当前路由 meta.btns
import { useRoute } from 'vue-router'
import { reactive } from 'vue'
export const useBtnAuth = () => {
  const route = useRoute()
  return route.meta.btns || reactive({})
}
