// Excel 导出配置，对应 config.yaml 中 excel 节点。
package config

type Excel struct {
	Dir string `mapstructure:"dir" json:"dir" yaml:"dir"`
}
