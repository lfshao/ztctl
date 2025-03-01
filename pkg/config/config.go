package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/lfshao/ztctl/pkg/output"
	ztcentral "github.com/zerotier/go-ztcentral"
)

// Config 配置管理器
type Config struct {
	ZerotierToken string
	ztClient      *ztcentral.Client
	ztClientOnce  sync.Once
}

var (
	instance *Config
	once     sync.Once
)

// GetConfig 获取配置实例（单例模式）
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := instance.init(); err != nil {
			output.DefaultPrinter.PrintMsg(output.LevelError, "初始化配置失败\n原因: %v", err)
			os.Exit(1)
		}
	})
	return instance
}

// init 初始化配置
func (c *Config) init() error {
	// 获取ZeroTier Token
	token := os.Getenv("ZEROTIER_CENTRAL_TOKEN")
	if token == "" {
		return fmt.Errorf("未设置ZEROTIER_CENTRAL_TOKEN环境变量\n请设置环境变量: export ZEROTIER_CENTRAL_TOKEN='你的API Token'")
	}
	c.ZerotierToken = token
	return nil
}

// GetZerotierToken 获取ZeroTier Token
func (c *Config) GetZerotierToken() string {
	return c.ZerotierToken
}

// GetZTClient 获取ZeroTier客户端实例（单例模式）
func (c *Config) GetZTClient() (*ztcentral.Client, error) {
	var err error
	c.ztClientOnce.Do(func() {
		c.ztClient, err = ztcentral.NewClient(c.ZerotierToken)
	})
	return c.ztClient, err
}
