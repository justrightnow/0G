// cmd 包包含所有命令行相关功能
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd 根命令，作为所有子命令的入口点
var rootCmd = &cobra.Command{
	Use:   "0g-storage-example",
	Short: "0g storage example",
	Long:  `0G 存储示例程序，提供文件上传和下载功能`,
}


// Execute is the command line entrypoint.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}