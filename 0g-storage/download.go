package main

import (
	"context"
	"log"
	// "os"
	// "strings"

	"github.com/0gfoundation/0g-storage-client/common/blockchain"
	"github.com/0gfoundation/0g-storage-client/indexer"
	// "github.com/spf13/cobra"
)

func main() {

// 创建 Web3 客户端连接区块链
	w3client := blockchain.MustNewWeb3("https://evmrpc-testnet.0g.ai/", "///")
	defer w3client.Close()

	// 创建索引器客户端，用于查询文件片段位置
	indexerClient, err := indexer.NewClient("https://indexer-storage-testnet-turbo.0g.ai")
	if err != nil {
		log.Fatalf("create indexer client error: %v", err)
	}

	ctx := context.Background()

	// 选择可用的存储节点（硬编码的节点地址）
	// nodes, err := indexerClient.SelectNodes(ctx, 1, []string{
	// 	"http://34.174.223.105:5678",
	// 	"http://104.196.238.89:5678",
	// 	"http://34.57.99.219:5678",
	// 	"http://34.55.197.204:5678",
	// 	"http://34.133.200.179:5678",
	// }, "max", true)
	// if err != nil {
	// 	log.Fatalf("select nodes error: %v", err)
	// }

	// 从环境变量获取文件的根哈希列表
	roots := "0x99ecd25a45685821069373c831afb4bbfc04b4aaf609d369576e0baf820e0571"

	// 执行文件下载：根据根哈希下载文件片段并重组为完整文件
	if err := indexerClient.Download(ctx, roots, "downloaded_file.bin", true); err != nil {
		log.Fatalf("Download file error: %v", err)
	}
	log.Printf("Download successful!\n")

}