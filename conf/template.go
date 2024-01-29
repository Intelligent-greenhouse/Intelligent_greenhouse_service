package conf

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed template.toml
var template embed.FS

// InitConfFile 初始化配置文件
func InitConfFile(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}

	confFile := filepath.Join(dir, "conf.toml")

	// 检查文件是否存在
	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		// 从嵌入的文件系统中提取文件
		data, err := fs.ReadFile(template, "template.toml")
		if err != nil {
			panic(err)
		}

		// 将文件写入目标路径
		err = os.WriteFile(confFile, data, 0644)
		if err != nil {
			panic(err)
		}
	}
}
