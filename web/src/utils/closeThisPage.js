// 关闭当前标签页：通过事件总线通知 tabs 组件
import { emitter } from '@/utils/bus.js'

export const closeThisPage = () => {
  emitter.emit('closeThisPage')
}
