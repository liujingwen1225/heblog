package miniblog

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

const (
	// defaultConfigName 指定了 heblog 服务的默认配置文件名.
	defaultConfigName = "heblog.yaml"
	// recommendedHomeDir 定义放置 heblog 服务配置的默认目录.
	recommendedHomeDir = ".heblog"
)

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		// 如果获取用户主目录失败，打印 `'Error: xxx` 错误，并退出程序（退出码为 1）
		cobra.CheckErr(err)
		// 将用 `$HOME/<recommendedHomeDir>` 目录加入到配置文件的搜索路径中
		viper.SetConfigFile(filepath.Join(home, recommendedHomeDir))
		// 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath(".")
		// 设置配置文件格式为 YAML (YAML 格式清晰易读，并且支持复杂的配置结构)
		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultConfigName)
	}
	viper.AutomaticEnv()
	// 读取环境变量的前缀为 MINIBLOG，如果是 heblog，将自动转变为大写。
	viper.SetEnvPrefix("MINIBLOG")
	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// 打印 viper 当前使用的配置文件，方便 Debug.
	fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
}
