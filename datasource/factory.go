package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hipeday/azir-rdb-golang-api/api"
	"github.com/jmoiron/sqlx"
)

func init() {
	api.RegisterClientFactory(api.MySQL, MySQLClientFactory{})
}

type MySQLClientFactory struct{}

func (f MySQLClientFactory) CreateClient(info api.ConnectionInfo) (api.Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		info.User, info.Password, info.Host, info.Port, info.Database)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &MySQLClient{db: db}, nil
}
