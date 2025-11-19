package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/justrightnow/0G/cmd"
)

// main 程序入口函数
// 负责加载环境变量并启动命令行界面
func main() {
	// 加载 .env 环境变量文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 执行命令行程序
	cmd.Execute()
}