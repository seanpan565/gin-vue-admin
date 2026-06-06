import { getBaseUrl } from '@/utils/format'

const defaults = { maxImageMB: 0.5, maxVideoMB: 5 }
let cached = null
let pending = null

export async function getUploadLimits() {
  if (cached) {
    return cached
  }
  if (!pending) {
    pending = fetch(`${getBaseUrl()}/base/uploadConfig`)
      .then((res) => res.json())
      .then((body) => {
        if (body?.code === 0 && body.data) {
          cached = {
            maxImageMB: Number(body.data.maxImageMB) || defaults.maxImageMB,
            maxVideoMB: Number(body.data.maxVideoMB) || defaults.maxVideoMB
          }
        } else {
          cached = { ...defaults }
        }
        return cached
      })
      .catch(() => {
        cached = { ...defaults }
        return cached
      })
      .finally(() => {
        pending = null
      })
  }
  return pending
}
