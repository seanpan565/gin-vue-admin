// ECharts 按需注册 + vue-echarts 懒加载入口。
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { LineChart } from 'echarts/charts'
import {
  AxisPointerComponent,
  GraphicComponent,
  GridComponent,
  TooltipComponent
} from 'echarts/components'

let registered = false

function registerEcharts() {
  if (registered) {
    return
  }
  use([
    CanvasRenderer,
    LineChart,
    GridComponent,
    TooltipComponent,
    GraphicComponent,
    AxisPointerComponent
  ])
  registered = true
}

let chartComponentPromise = null

/** 懒加载 vue-echarts 组件（仅首次调用时注册 ECharts 模块）。 */
export function loadChartComponent() {
  if (!chartComponentPromise) {
    chartComponentPromise = import('vue-echarts').then((mod) => {
      registerEcharts()
      return mod.default
    })
  }
  return chartComponentPromise
}
