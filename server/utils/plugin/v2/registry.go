package plugin

// v2 插件注册表，支持插件自动发现与初始化。
import "sync"

var (
	registryMu sync.RWMutex
	registry   []Plugin
)

// Register 注册插件实例，供启动时自动初始化
func Register(p Plugin) {
	if p == nil {
		return
	}
	registryMu.Lock()
	registry = append(registry, p)
	registryMu.Unlock()
}

// Registered 返回已注册插件的快照副本
func Registered() []Plugin {
	registryMu.RLock()
	defer registryMu.RUnlock()
	out := make([]Plugin, len(registry))
	copy(out, registry)
	return out
}
