package datasource

import (
	"github.com/hipeday/azir-rdb-golang-api/api"
	"github.com/jmoiron/sqlx"
)

type MySQLClient struct {
	api.Client
	db *sqlx.DB
}

// Databases 实现 Client 接口的方法
func (c *MySQLClient) Databases() ([]string, error) {
	var databases []string
	err := c.db.Select(&databases, "SHOW DATABASES")
	if err != nil {
		return nil, err
	}
	return databases, nil
}
