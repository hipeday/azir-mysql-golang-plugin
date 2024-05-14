package plugin

import (
	"encoding/json"
	"github.com/hipeday/azir-plugin-golang/pkg/plugin"
	"github.com/hipeday/azir-plugin-golang/pkg/properties"
	"net"
	"path/filepath"
)

type MySQL interface {
	plugin.Callback
	Databases(args []string) (interface{}, error)
}

type MySQLPlugin struct {
	plugin.ListenPlugin
	MySQL
}

func (m *MySQLPlugin) Run(args []string) (interface{}, error) {
	return m.ListenPlugin.Run(args)
}

func (m *MySQLPlugin) CallbackRender(result interface{}, args []string) error {
	var err error
	_, err = m.ParseConfig("parse", args)
	if err != nil {
		return err
	}
	var (
		property   = m.GetConfig().(properties.DefaultProperty)
		logger     = m.GetLogger()
		socketHome = filepath.Join(property.Home, property.Name, "socks")
		socketPath = filepath.Join(socketHome, "plugin.sock")
	)

	conn, err := net.Dial("unix", socketPath)

	if err != nil {
		logger.Fatalf("Error dialing socket: %v", err)
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			logger.Fatalf("Error closing connection: %v", err)
		}
	}(conn)

	body, err := json.Marshal(result)

	if err != nil {
		return err
	}

	_, err = conn.Write(body)

	return err
}
