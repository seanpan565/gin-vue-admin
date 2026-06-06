<!-- 菜单图标选择 -->
<template>
  <div class="w-full">
    <el-select
      v-model="value"
      clearable
      filterable
      placeholder="请选择"
      class="w-full"
    >
      <template #prefix>
        <el-icon>
          <component v-if="value" :is="value" />
        </el-icon>
      </template>
      <el-option
        v-for="item in iconOptions"
        :key="item.key"
        class="select__option_item"
        :label="item.key"
        :value="item.key"
      >
        <span class="gva-icon" style="padding: 3px 0 0" :class="item.label">
          <el-icon>
            <component v-if="item.label" :is="item.label" />
          </el-icon>
        </span>
        <span style="text-align: left">{{ item.key }}</span>
      </el-option>
    </el-select>
  </div>
</template>

<script setup>
  import { computed, onMounted, ref } from 'vue'
  import config from '@/core/config'

  defineOptions({
    name: 'Icon'
  })

  const value = defineModel()
  const options = ref([])

  onMounted(async () => {
    const mod = await import('./menu-icons.json')
    options.value = mod.default || mod
  })

  const iconOptions = computed(() => options.value.concat(config.logs))
</script>

<style lang="scss">
  .gva-icon {
    color: rgb(132, 146, 166);
    font-size: 14px;
    margin-right: 10px;
  }

  .select__option_item {
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }
</style>
