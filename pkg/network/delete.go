package network

import (
	"context"
	"fmt"

	"github.com/lfshao/ztctl/pkg/config"
	"github.com/lfshao/ztctl/pkg/output"
)

// Delete 删除指定的ZeroTier网络
func Delete(networkID string) error {
	if networkID == "" {
		return fmt.Errorf("网络ID不能为空")
	}

	c, err := config.GetConfig().GetZTClient()
	if err != nil {
		return fmt.Errorf("连接ZeroTier Central失败: %v", err)
	}

	ctx := context.Background()

	// 删除网络
	if err := c.DeleteNetwork(ctx, networkID); err != nil {
		return fmt.Errorf("删除网络失败: %v", err)
	}

	output.DefaultPrinter.PrintMsg(output.LevelSuccess, "成功删除网络 %s", networkID)
	return nil
}
