package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ztctl",
	Short: "ZeroTier网络管理工具",
	Long: `ztctl是一个命令行工具，用于管理ZeroTier网络。
它可以帮助你创建、删除、修改网络，以及管理网络成员。`,
	SilenceUsage: true,
}

func Execute() error {
	return rootCmd.Execute()
}