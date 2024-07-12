/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"jobs/cmd"
	"jobs/initialize"
)

func main() {
	initialize.InitLogger()
	initialize.InitReadProfile()
	initialize.InitReadNacos()
	initialize.InitRedis()
	initialize.InitRocketmq()
	cmd.Execute()
}
