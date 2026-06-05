// 磁盘监控配置，对应 config.yaml 中 disk-list 节点。
package config

type Disk struct {
	MountPoint string `mapstructure:"mount-point" json:"mount-point" yaml:"mount-point"`
}

type DiskList struct {
	Disk `yaml:",inline" mapstructure:",squash"`
}
