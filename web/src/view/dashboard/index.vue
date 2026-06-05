<!-- 仪表盘首页 -->
<template>
  <div class="h-full gva-container2 overflow-auto bg-slate-50/60 dark:bg-slate-900">
    <div class="space-y-4 p-4 lg:p-6">
      <section
        class="relative overflow-hidden rounded-xl border border-slate-200/80 bg-white px-5 py-6 shadow-sm dark:border-slate-700 dark:bg-slate-800"
      >
        <div class="relative flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
          <div>
            <p class="text-xs tracking-[0.2em] text-slate-500 dark:text-slate-400">DASHBOARD</p>
            <h1 class="mt-2 text-xl font-semibold text-slate-900 dark:text-slate-100 lg:text-2xl">
              欢迎回来，{{ appName }}
            </h1>
            <p class="mt-2 text-sm text-slate-600 dark:text-slate-300">
              {{ today }} · 核心业务概览与开发快捷入口
            </p>
          </div>
          <div class="flex flex-wrap items-center gap-2">
            <el-button type="primary" @click="goSwagger">API 文档</el-button>
            <el-button @click="goDocs">项目文档</el-button>
          </div>
        </div>
      </section>

      <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 xl:grid-cols-3">
        <gva-card>
          <gva-chart :type="1" title="访问人数" />
        </gva-card>
        <gva-card>
          <gva-chart :type="2" title="新增客户" />
        </gva-card>
        <gva-card>
          <gva-chart :type="3" title="解决数量" />
        </gva-card>
      </div>

      <div class="grid grid-cols-1 items-stretch gap-4 xl:grid-cols-12">
        <div class="grid grid-cols-1 gap-4 content-start xl:col-span-8 xl:h-full">
          <gva-card title="内容数据">
            <gva-chart :type="4" />
          </gva-card>

          <gva-card title="最新更新">
            <gva-table />
          </gva-card>
        </div>

        <div class="flex flex-col gap-4 xl:col-span-4 xl:h-full">
          <gva-card title="快捷功能" show-action custom-class="min-h-[300px]">
            <gva-quick-link />
          </gva-card>
          <gva-card title="公告" show-action custom-class="min-h-[260px]">
            <gva-notice />
          </gva-card>
          <gva-card title="文档" show-action custom-class="min-h-[120px]">
            <gva-wiki />
          </gva-card>
          <div
            class="relative min-h-[180px] flex-1 overflow-hidden rounded-lg border border-slate-200 bg-slate-900 p-5 text-white shadow-sm dark:border-slate-700"
          >
            <div class="relative">
              <div class="inline-flex rounded-full bg-white/10 px-3 py-1 text-xs">开发提示</div>
              <h3 class="mt-3 text-lg font-semibold">商城模块开发建议</h3>
              <p class="mt-2 text-sm text-slate-200/90">
                后端在 server/model|service|api|router/mall 下扩展；前端在 web/src/view/mall 与 api/mall 下扩展。权限与菜单继续复用 system 模块。
              </p>
              <div class="mt-4 flex flex-wrap gap-2 text-xs">
                <span class="rounded-full bg-white/10 px-2.5 py-1">Router → API → Service</span>
                <span class="rounded-full bg-white/10 px-2.5 py-1">参考 example/customer</span>
                <span class="rounded-full bg-white/10 px-2.5 py-1">金额用分(int64)</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
  import { computed } from 'vue'
  import config from '@/core/config'
  import { externalLinks } from '@/core/site'
  import { GvaTable, GvaChart, GvaWiki, GvaNotice, GvaQuickLink, GvaCard } from './components'

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

  const goDocs = () => {
    window.open(externalLinks.gvaDocs, '_blank', 'noopener,noreferrer')
  }

  defineOptions({
    name: 'Dashboard'
  })
</script>

<style lang="scss" scoped></style>
