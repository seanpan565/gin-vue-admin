<!-- 仪表盘首页 -->
<template>
  <div class="h-full gva-container2 overflow-auto bg-slate-50/60 dark:bg-slate-900">
    <div class="space-y-4 p-4 lg:p-6">
      <section
        class="overflow-hidden rounded-xl border border-slate-200/80 bg-white px-5 py-6 shadow-sm dark:border-slate-700 dark:bg-slate-800"
      >
        <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
          <div>
            <p class="text-xs tracking-[0.2em] text-slate-500 dark:text-slate-400">DASHBOARD</p>
            <h1 class="mt-2 text-xl font-semibold text-slate-900 dark:text-slate-100 lg:text-2xl">
              欢迎回来，{{ appName }}
            </h1>
            <p class="mt-2 text-sm text-slate-600 dark:text-slate-300">
              {{ today }} · 商城管理后台
            </p>
          </div>
          <div class="flex flex-wrap items-center gap-2">
            <el-button type="primary" @click="goSwagger">API 文档</el-button>
          </div>
        </div>
      </section>

      <div class="grid grid-cols-1 items-stretch gap-4 xl:grid-cols-12">
        <div class="flex flex-col gap-4 xl:col-span-8">
          <gva-card title="快捷功能" show-action custom-class="min-h-[300px]">
            <gva-quick-link />
          </gva-card>
        </div>
        <div class="flex flex-col gap-4 xl:col-span-4">
          <gva-card title="公告" show-action custom-class="min-h-[260px]">
            <gva-notice />
          </gva-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
  import { computed } from 'vue'
  import config from '@/core/config'
  import { externalLinks } from '@/core/site'
  import { GvaNotice, GvaQuickLink, GvaCard } from './components'

  const appName = config.appName

  const today = computed(() => {
    try {
      const d = new Date()
      return d.toLocaleDateString('zh-CN', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit'
      })
    } catch (e) {
      return new Date().toISOString().slice(0, 10)
    }
  })

  const goSwagger = () => {
    window.open(externalLinks.swagger(import.meta.env.VITE_SERVER_PORT), '_blank', 'noopener,noreferrer')
  }

  defineOptions({
    name: 'Dashboard'
  })
</script>

<style lang="scss" scoped></style>
