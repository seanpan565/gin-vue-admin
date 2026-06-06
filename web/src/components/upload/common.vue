<!-- 通用文件上传组件，支持图片与视频 -->
<template>
  <div>
    <el-upload
      :action="`${getBaseUrl()}/fileUploadAndDownload/upload`"
      :before-upload="checkFile"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :show-file-list="false"
      :data="{'classId': props.classId}"
      :headers="{'x-token': token}"
      multiple
      class="upload-btn"
    >
      <el-button type="primary" :icon="Upload">普通上传</el-button>
    </el-upload>
  </div>
</template>

<script setup>
  import { onMounted, ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { isVideoMime, isImageMime } from '@/utils/image'
  import { getBaseUrl } from '@/utils/format'
  import { getUploadLimits } from '@/utils/uploadLimit'
  import { Upload } from "@element-plus/icons-vue";
  import { useUserStore } from "@/pinia";

  defineOptions({
    name: 'UploadCommon'
  })

  const userStore = useUserStore()

  const token = userStore.token

  const props = defineProps({
    classId: {
      type: Number,
      default: 0
    }
  })

  const emit = defineEmits(['on-success'])

  const fullscreenLoading = ref(false)
  const uploadLimits = ref({ maxImageMB: 0.5, maxVideoMB: 5 })

  onMounted(async () => {
    uploadLimits.value = await getUploadLimits()
  })

  const checkFile = (file) => {
    fullscreenLoading.value = true
    const { maxImageMB, maxVideoMB } = uploadLimits.value
    const isLtImage = file.size / 1024 / 1024 < maxImageMB
    const isLtVideo = file.size / 1024 / 1024 < maxVideoMB
    const isVideo = isVideoMime(file.type)
    const isImage = isImageMime(file.type)
    let pass = true
    if (!isVideo && !isImage) {
      ElMessage.error(
        '上传图片只能是 jpg,png,svg,webp 格式, 上传视频只能是 mp4,webm 格式!'
      )
      fullscreenLoading.value = false
      pass = false
    }
    if (!isLtVideo && isVideo) {
      ElMessage.error(`上传视频大小不能超过 ${maxVideoMB}MB`)
      fullscreenLoading.value = false
      pass = false
    }
    if (!isLtImage && isImage) {
      ElMessage.error(`未压缩的上传图片大小不能超过 ${maxImageMB}MB，请使用压缩上传`)
      fullscreenLoading.value = false
      pass = false
    }

    return pass
  }

  const uploadSuccess = (res) => {
    const { data } = res
    if (data.file) {
      emit('on-success', data.file.url)
    }
    fullscreenLoading.value = false
  }

  const uploadError = () => {
    ElMessage.error('上传失败')
    fullscreenLoading.value = false
  }
</script>
