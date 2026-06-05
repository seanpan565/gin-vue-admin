<!-- 仪表盘 Wiki 链接 -->
<template>
  <div class="grid grid-cols-2 gap-2">
    <a
      v-for="item in wikis"
      :key="item.url"
      :href="item.url"
      class="text-sm text-black/70 dark:text-white/70 no-underline hover:text-[var(--el-color-primary)] dark:hover:text-white"
      target="_blank"
      rel="noopener noreferrer"
    >
      {{ item.title }}
    </a>
  </div>
</template>

<script setup>
  import { externalLinks, siteFeatures } from '@/core/site'

  const wikis = [
    { title: 'GVA 文档', url: externalLinks.gvaDocs },
    { title: 'Gin 文档', url: externalLinks.ginDocs },
    { title: 'Vue3 文档', url: externalLinks.vueDocs },
    { title: 'Swagger', url: externalLinks.swagger(import.meta.env.VITE_SERVER_PORT) },
    { title: 'GitHub', url: externalLinks.github }
  ]

  if (siteFeatures.pluginMarket) {
    wikis.push({ title: '插件市场', url: 'https://plugin.gin-vue-admin.com/' })
  }
</script>

<style scoped lang="scss"></style>
