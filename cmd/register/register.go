package register

import (
	"github.com/hipeday/azir-mysql-golang-plugin/plugin"
	"github.com/hipeday/azir-plugin-golang/cmd/command"
	plugin2 "github.com/hipeday/azir-plugin-golang/pkg/plugin"
)

var (
	pluginIns plugin.MySQLPlugin
)

func init() {
	pluginIns = plugin.MySQLPlugin{}
	command.Registry.RegisterCommand("run", plugin2.CommandFunctions{Command: pluginIns.Run, Callback: nil})
	command.Registry.RegisterCommand("databases", plugin2.CommandFunctions{Command: pluginIns.Databases, Callback: pluginIns.CallbackRender})
}
