package network

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lfshao/ztctl/pkg/config"
	"github.com/lfshao/ztctl/pkg/output"
)

// AuthorizeMember 授权成员加入网络
func AuthorizeMember(networkID string, memberID string) error {
	if networkID == "" {
		return fmt.Errorf("网络ID不能为空")
	}
	if memberID == "" {
		return fmt.Errorf("成员ID不能为空")
	}

	// 处理成员ID格式
	if strings.Contains(memberID, networkID) {
		// 如果成员ID包含网络ID前缀，提取实际的设备ID部分
		parts := strings.Split(memberID, "-")
		if len(parts) == 2 {
			memberID = parts[1]
		}
	}

	c, err := config.GetConfig().GetZTClient()
	if err != nil {
		return fmt.Errorf("连接ZeroTier Central失败: %v", err)
	}

	ctx := context.Background()

	// 获取成员信息
	member, err := c.GetMember(ctx, networkID, memberID)
	if err != nil {
		return fmt.Errorf("获取成员信息失败: %v", err)
	}

	// 设置授权状态为true
	authorized := true
	member.Config.Authorized = &authorized

	// 更新成员信息
	_, err = c.UpdateMember(ctx, networkID, memberID, member)
	if err != nil {
		return fmt.Errorf("更新成员授权状态失败: %v", err)
	}

	output.DefaultPrinter.PrintMsg(output.LevelSuccess, "已授权成员 %s 加入网络", memberID)
	return nil
}

// DeauthorizeMember 取消成员的网络访问权限
func DeauthorizeMember(networkID string, memberID string) error {
	if networkID == "" {
		return fmt.Errorf("网络ID不能为空")
	}
	if memberID == "" {
		return fmt.Errorf("成员ID不能为空")
	}

	// 处理成员ID格式
	if strings.Contains(memberID, networkID) {
		// 如果成员ID包含网络ID前缀，提取实际的设备ID部分
		parts := strings.Split(memberID, "-")
		if len(parts) == 2 {
			memberID = parts[1]
		}
	}

	c, err := config.GetConfig().GetZTClient()
	if err != nil {
		return fmt.Errorf("连接ZeroTier Central失败: %v", err)
	}

	ctx := context.Background()

	// 获取成员信息
	member, err := c.GetMember(ctx, networkID, memberID)
	if err != nil {
		return fmt.Errorf("获取成员信息失败: %v", err)
	}

	// 设置授权状态为false
	authorized := false
	member.Config.Authorized = &authorized

	// 更新成员信息
	_, err = c.UpdateMember(ctx, networkID, memberID, member)
	if err != nil {
		return fmt.Errorf("更新成员授权状态失败: %v", err)
	}

	output.DefaultPrinter.PrintMsg(output.LevelSuccess, "已取消成员 %s 的网络访问权限", memberID)
	return nil
}

// ListMembers 列出指定网络的所有成员
func ListMembers(networkID string) error {
	if networkID == "" {
		return fmt.Errorf("\033[31m错误: 网络ID不能为空\033[0m")
	}

	c, err := config.GetConfig().GetZTClient()
	if err != nil {
		return fmt.Errorf("\033[31m错误: 连接ZeroTier Central失败\033[0m\n原因: %v\n建议: 请检查API Token是否正确", err)
	}

	ctx := context.Background()

	// 获取网络信息
	network, err := c.GetNetwork(ctx, networkID)
	if err != nil {
		return fmt.Errorf("\033[31m错误: 获取网络信息失败\033[0m\n原因: %v\n建议: 请检查网络ID是否正确", err)
	}

	// 获取成员列表
	members, err := c.GetMembers(ctx, networkID)
	if err != nil {
		return fmt.Errorf("\033[31m错误: 获取成员列表失败\033[0m\n原因: %v\n建议: 请确保网络连接正常且API Token具有足够的权限", err)
	}

	if len(members) == 0 {
		output.DefaultPrinter.PrintMsg(output.LevelInfo, "网络 %s (%s) 当前没有任何成员", *network.Config.Name, networkID)
		return nil
	}

	// 准备表格数据
	headers := []string{"成员ID", "名称", "IP地址", "状态", "最后在线时间"}
	rows := make([][]string, 0, len(members))

	// 打印网络信息
	output.DefaultPrinter.PrintMsg(output.LevelInfo, "网络: %s (%s)", *network.Config.Name, networkID)

	// 处理成员信息
	for _, m := range members {
		// 获取IP地址
		ip := "无"
		if m.Config != nil && m.Config.IpAssignments != nil {
			assignments := *m.Config.IpAssignments
			if len(assignments) > 0 {
				ip = assignments[0]
			}
		}

		// 获取在线状态
		status := "离线"
		if m.Config != nil && m.Config.Authorized != nil && *m.Config.Authorized {
			status = "在线"
		}

		// 获取名称
		name := "未命名"
		if m.Name != nil && *m.Name != "" {
			name = *m.Name
		}

		// 获取最后在线时间
		lastOnline := "从未在线"
		if m.Config != nil && m.Config.CreationTime != nil {
			lastOnline = time.Unix(*m.Config.CreationTime/1000, 0).Format("2006-01-02 15:04:05")
		}

		rows = append(rows, []string{*m.Id, name, ip, status, lastOnline})
	}

	// 打印表格
	output.DefaultPrinter.PrintTable(headers, rows)
	return nil
}
