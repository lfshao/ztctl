package network

import (
	"context"
	"fmt"

	"github.com/lfshao/ztctl/pkg/config"
	"github.com/lfshao/ztctl/pkg/output"
)

// Get 获取所有网络及其成员信息
func Get() error {
	c, err := config.GetConfig().GetZTClient()
	if err != nil {
		return fmt.Errorf("连接ZeroTier Central失败: %v", err)
	}

	ctx := context.Background()

	// 获取网络列表
	networks, err := c.GetNetworks(ctx)
	if err != nil {
		return fmt.Errorf("获取网络列表失败: %v", err)
	}

	if len(networks) == 0 {
		output.DefaultPrinter.PrintMsg(output.LevelInfo, "当前没有任何网络")
		return nil
	}

	// 准备表格数据
	headers := []string{"网络ID", "网络名称", "成员数量"}
	rows := make([][]string, 0, len(networks))

	// 处理网络信息
	for _, n := range networks {
		members, err := c.GetMembers(ctx, *n.Id)
		if err != nil {
			output.DefaultPrinter.PrintMsg(output.LevelWarning, "获取网络 %s 的成员列表失败: %v", *n.Id, err)
			continue
		}

		rows = append(rows, []string{*n.Id, *n.Config.Name, fmt.Sprintf("%d", len(members))})
	}

	// 打印表格
	output.DefaultPrinter.PrintTable(headers, rows)
	return nil
}

// Create 创建新的ZeroTier网络
func Create(name string, description string) error {
	if name == "" {
		return fmt.Errorf("网络名称不能为空")
	}

	c, err := config.GetConfig().GetZTClient()
	if err != nil {
		return fmt.Errorf("连接ZeroTier Central失败: %v", err)
	}

	ctx := context.Background()

	// 创建网络
	network, err := c.NewNetwork(ctx, name, nil)
	if err != nil {
		return fmt.Errorf("创建网络失败: %v", err)
	}

	output.DefaultPrinter.PrintMsg(output.LevelSuccess, "成功创建网络 %s", *network.Id)
	return nil
}
