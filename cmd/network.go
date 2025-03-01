package cmd

import (
	"github.com/lfshao/ztctl/pkg/network"
	"github.com/lfshao/ztctl/pkg/output"
	"github.com/spf13/cobra"
)

var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "管理ZeroTier网络",
	Long:  `用于创建、列出、删除和修改ZeroTier网络。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有网络",
	RunE: func(cmd *cobra.Command, args []string) error {
		return network.Get()
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "创建新的网络",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		desc, _ := cmd.Flags().GetString("description")

		if err := network.Create(name, desc); err != nil {
			output.DefaultPrinter.PrintMsg(output.LevelError, "创建网络失败: %v", err)
			return
		}
		output.DefaultPrinter.PrintMsg(output.LevelSuccess, "网络创建成功")
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除指定的网络",
	Run: func(cmd *cobra.Command, args []string) {
		networkID, _ := cmd.Flags().GetString("id")
		if err := network.Delete(networkID); err != nil {
			output.DefaultPrinter.PrintMsg(output.LevelError, "删除网络失败: %v", err)
			return
		}
		output.DefaultPrinter.PrintMsg(output.LevelSuccess, "网络删除成功")
	},
}

var membersCmd = &cobra.Command{
	Use:   "members",
	Short: "列出指定网络的所有成员",
	Run: func(cmd *cobra.Command, args []string) {
		networkID, _ := cmd.Flags().GetString("id")
		if err := network.ListMembers(networkID); err != nil {
			output.DefaultPrinter.PrintMsg(output.LevelError, "%v", err)
			return
		}
	},
}

var authorizeCmd = &cobra.Command{
	Use:   "authorize",
	Short: "授权成员加入网络",
	Run: func(cmd *cobra.Command, args []string) {
		networkID, _ := cmd.Flags().GetString("network-id")
		memberID, _ := cmd.Flags().GetString("member-id")
		if err := network.AuthorizeMember(networkID, memberID); err != nil {
			output.DefaultPrinter.PrintMsg(output.LevelError, "%v", err)
			return
		}
	},
}

var deauthorizeCmd = &cobra.Command{
	Use:   "deauthorize",
	Short: "取消成员的网络访问权限",
	Run: func(cmd *cobra.Command, args []string) {
		networkID, _ := cmd.Flags().GetString("network-id")
		memberID, _ := cmd.Flags().GetString("member-id")
		if err := network.DeauthorizeMember(networkID, memberID); err != nil {
			output.DefaultPrinter.PrintMsg(output.LevelError, "%v", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(networkCmd)
	networkCmd.AddCommand(listCmd)
	networkCmd.AddCommand(createCmd)
	networkCmd.AddCommand(deleteCmd)
	networkCmd.AddCommand(membersCmd)
	networkCmd.AddCommand(authorizeCmd)
	networkCmd.AddCommand(deauthorizeCmd)

	createCmd.Flags().String("name", "", "网络名称")
	createCmd.Flags().String("description", "", "网络描述")
	deleteCmd.Flags().String("id", "", "网络ID")
	membersCmd.Flags().String("id", "", "网络ID")
	authorizeCmd.Flags().String("network-id", "", "网络ID")
	authorizeCmd.Flags().String("member-id", "", "成员ID")
	deauthorizeCmd.Flags().String("network-id", "", "网络ID")
	deauthorizeCmd.Flags().String("member-id", "", "成员ID")
}
