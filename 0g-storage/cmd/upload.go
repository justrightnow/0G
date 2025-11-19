// cmd 包 - 上传命令实现
package cmd

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/0gfoundation/0g-storage-client/common/blockchain"
	"github.com/0gfoundation/0g-storage-client/core"
	"github.com/0gfoundation/0g-storage-client/indexer"
	"github.com/0gfoundation/0g-storage-client/transfer"
	"github.com/spf13/cobra"
)

// uploadCmd 上传命令定义
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload files to 0g storage",
	Long:  `上传文件到 0G 分布式存储网络，支持文件分片和分布式存储`,
	Run: func(cmd *cobra.Command, args []string) {
		// 创建 Web3 客户端连接区块链
		w3client := blockchain.MustNewWeb3(os.Getenv("EVM_RPC"), os.Getenv("PRIVATE_KEY"))
		defer w3client.Close()

		// 创建索引器客户端，用于查询存储节点信息
		indexerClient, err := indexer.NewClient(os.Getenv("INDEXER_RPC"))
		if err != nil {
			log.Fatalf("create indexer client error: %v", err)
		}

		ctx := context.Background()

		// 选择可用的存储节点（硬编码的节点地址）
		nodes, err := indexerClient.SelectNodes(ctx, 1, []string{
			"http://34.174.223.105:5678",
			"http://104.196.238.89:5678",
			"http://34.57.99.219:5678",
			"http://34.55.197.204:5678",
			"http://34.133.200.179:5678",
		}, "max", true)
		if err != nil {
			log.Fatalf("select nodes error: %v", err)
		}

		// 创建文件上传器
		uploader, err := transfer.NewUploader(ctx, w3client, nodes)
		if err != nil {
			log.Fatalf("create uploader error: %v", err)
		}

		// 打开要上传的文件
		file, err := core.Open(os.Getenv("FILE_NAME"))
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer file.Close()

		log.Printf("Begin to upload file ...\n")

		// 解析分片大小配置
		fragmentSizeStr := os.Getenv("FRAGMENT_SIZE")
		fragmentSize, err := strconv.ParseInt(fragmentSizeStr, 10, 64)
		if err != nil {
			log.Fatalf("Error fragment size: %v", err)
		}

		// 执行文件分片上传
		_, roots, err := uploader.SplitableUpload(ctx, file, fragmentSize, transfer.UploadOption{
			SkipTx:           true,                    // 跳过交易，只上传数据
			FinalityRequired: transfer.TransactionPacked, // 要求交易被打包
			FullTrusted:      false,                   // 不完全信任模式
			NRetries:         10,                      // 重试次数
			TaskSize:         10,                      // 任务大小
			Method:           "10",                    // 上传方法
		})
		if err != nil {
			log.Fatalf("upload file error: %v", err)
		}

		log.Printf("Upload successful!\n")
		log.Printf("Roots size: %d\n", len(roots))

		// 将根哈希转换为字符串格式
		s := make([]string, len(roots))
		for i, root := range roots {
			s[i] = root.String()
		}
		log.Printf("File uploaded in %v fragments, roots = %v", len(roots), strings.Join(s, ","))
	},
}

// init 初始化函数，将上传命令注册到根命令
func init() {
	rootCmd.AddCommand(uploadCmd)
}