/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "cobra/cmd"

func main() {
	cmd.RedisClient()
	cmd.InitRocketmq()
	cmd.Execute()
}
