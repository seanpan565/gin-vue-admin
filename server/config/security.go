// security.go 安全相关配置，对应 config.yaml 中 security 节点。
package config

// Security 安全加固与可观测性开关。
type Security struct {
	SwaggerEnable       bool     `mapstructure:"swagger-enable" json:"swagger-enable" yaml:"swagger-enable"`
	MetricsEnable       bool     `mapstructure:"metrics-enable" json:"metrics-enable" yaml:"metrics-enable"`
	MetricsAuthToken    string   `mapstructure:"metrics-auth-token" json:"metrics-auth-token" yaml:"metrics-auth-token"`
	ErrorReportKey      string   `mapstructure:"error-report-key" json:"error-report-key" yaml:"error-report-key"`
	AuthLimitCount      int      `mapstructure:"auth-limit-count" json:"auth-limit-count" yaml:"auth-limit-count"`
	AuthLimitTime       int      `mapstructure:"auth-limit-time" json:"auth-limit-time" yaml:"auth-limit-time"`
	MemberCaptchaEnable bool     `mapstructure:"member-captcha-enable" json:"member-captcha-enable" yaml:"member-captcha-enable"`
	UploadAllowedExts   []string  `mapstructure:"upload-allowed-exts" json:"upload-allowed-exts" yaml:"upload-allowed-exts"`
	UploadMaxImageMB    float64   `mapstructure:"upload-max-image-mb" json:"upload-max-image-mb" yaml:"upload-max-image-mb"`
	UploadMaxVideoMB    float64   `mapstructure:"upload-max-video-mb" json:"upload-max-video-mb" yaml:"upload-max-video-mb"`
}
