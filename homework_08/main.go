package main

import "homework_08/boot"

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.MysqlSetup()
	boot.RedisSetup()

}
