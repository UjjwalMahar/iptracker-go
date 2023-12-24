/*
Copyright © 2023 NAME HERE <ujjwal.mahar01@gmail.com>
*/
package main

import (
	"github.com/UjjwalMahar/iptracker/cmd"
	"github.com/UjjwalMahar/iptracker/initializer"
)

func init(){
	initializer.LoadEnv()
}
func main() {
	cmd.Execute()
}
