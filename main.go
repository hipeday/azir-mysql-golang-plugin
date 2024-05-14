package main

import (
	_ "github.com/hipeday/azir-mysql-golang-plugin/cmd/datasource"
	_ "github.com/hipeday/azir-mysql-golang-plugin/cmd/register"
	"github.com/hipeday/azir-plugin-golang/cmd/runner"
)

func main() {
	runner.Run()
}
