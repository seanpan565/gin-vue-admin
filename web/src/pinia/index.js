// Pinia 状态管理入口，统一导出各 store
import { createPinia } from 'pinia'
import { useAppStore } from '@/pinia/modules/app'
import { useUserStore } from '@/pinia/modules/user'
import { useDictionaryStore } from '@/pinia/modules/dictionary'

const store = createPinia()

export { store, useAppStore, useUserStore, useDictionaryStore }
