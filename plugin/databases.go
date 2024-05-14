package plugin

import (
	"github.com/hipeday/azir-rdb-golang-api/api"
)

func (m *MySQLPlugin) Databases(args []string) (interface{}, error) {
	var err error
	_, err = m.ParseConfig("databases", args)
	if err != nil {
		return nil, err
	}
	var (
		logger = m.GetLogger()
	)

	client, err := api.NewClient(api.MySQL, api.ConnectionInfo{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "000000..",
		Database: "",
	})

	if err != nil {
		logger.Fatalf("Error creating mysql client: %v", err)
	}

	return client.Databases()
}
