// cmd 包 - 下载命令实现
package cmd

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/0gfoundation/0g-storage-client/common/blockchain"
	"github.com/0gfoundation/0g-storage-client/indexer"
	"github.com/spf13/cobra"
)

// downloadCmd 下载命令定义
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download files from 0g storage",
	Long:  `从 0G 分布式存储网络下载文件，使用根哈希定位并重组文件`,
	Run: func(cmd *cobra.Command, args []string) {
		// 创建 Web3 客户端连接区块链
		w3client := blockchain.MustNewWeb3(os.Getenv("EVM_RPC"), os.Getenv("PRIVATE_KEY"))
		defer w3client.Close()

		// 创建索引器客户端，用于查询文件片段位置
		indexerClient, err := indexer.NewClient(os.Getenv("INDEXER_RPC"))
		if err != nil {
			log.Fatalf("create indexer client error: %v", err)
		}

		ctx := context.Background()

		// 从环境变量获取文件的根哈希列表
		roots := os.Getenv("ROOTS")

		// 执行文件下载：根据根哈希下载文件片段并重组为完整文件
		if err := indexerClient.DownloadFragments(ctx, strings.Split(roots, ","), "downloaded_file.bin", true); err != nil {
			log.Fatalf("Download file error: %v", err)
		}
		log.Printf("Download successful!\n")
	},
}

// init 初始化函数，将下载命令注册到根命令
func init() {
	rootCmd.AddCommand(downloadCmd)
}