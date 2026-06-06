package system

// BaseApi 验证码生成与校验接口
import (
	"time"

	"mall-admin/server/global"
	"mall-admin/server/model/common/response"
	systemRes "mall-admin/server/model/system/response"
	captchaStore "mall-admin/server/utils/captcha"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type BaseApi struct{}

// Captcha
// @Tags      Base
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	key := c.ClientIP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, captchaStore.Store(c.Request.Context()))
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.GVA_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}

// UploadConfig
// @Tags      Base
// @Summary   获取前端上传大小限制（与 security 配置一致）
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]float64,msg=string}
// @Router    /base/uploadConfig [get]
func (b *BaseApi) UploadConfig(c *gin.Context) {
	imageMB, videoMB := uploadSizeLimits()
	response.OkWithDetailed(gin.H{
		"maxImageMB": imageMB,
		"maxVideoMB": videoMB,
	}, "获取成功", c)
}

func uploadSizeLimits() (imageMB, videoMB float64) {
	imageMB = global.GVA_CONFIG.Security.UploadMaxImageMB
	if imageMB <= 0 {
		imageMB = 0.5
	}
	videoMB = global.GVA_CONFIG.Security.UploadMaxVideoMB
	if videoMB <= 0 {
		videoMB = 5
	}
	return imageMB, videoMB
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}
