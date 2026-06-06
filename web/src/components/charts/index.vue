<!-- ECharts 图表封装组件，懒加载 vue-echarts，支持窗口自适应 -->
<template>
  <component
    :is="VCharts"
    v-if="renderChart && VCharts"
    :option="options"
    :autoresize="autoResize"
    :style="{ width, height }"
  />
</template>

<script setup>
  import { nextTick, onMounted, ref, shallowRef } from 'vue'
  import { loadChartComponent } from '@/utils/echarts'
  import { useWindowResize } from '@/hooks/use-windows-resize'

  defineProps({
    options: {
      type: Object,
      default() {
        return {}
      }
    },
    autoResize: {
      type: Boolean,
      default: true
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '100%'
    }
  })

  const VCharts = shallowRef(null)
  const renderChart = ref(false)

  const mountChart = async () => {
    if (!VCharts.value) {
      VCharts.value = await loadChartComponent()
    }
    renderChart.value = false
    await nextTick()
    renderChart.value = true
  }

  onMounted(mountChart)

  useWindowResize(() => {
    if (VCharts.value) {
      mountChart()
    }
  })
</script>
