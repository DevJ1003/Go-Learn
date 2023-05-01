package dbutil

import (
	"fmt"
)

// Info ...
type Info struct {
	ID   int
	Name string
}

// DbDriver ...
const DbDriver = "mysql"

// Username ...
const User = "root"

// Password ...
const Password = "password"

// DbName ...
const DbName = "go_db1"

// TableName ...
const TableName = "person"

// DataSourceName ...
// DataSourceName := "username:password@tcp/go_db1?charset=utf8"
var DataSourceName = fmt.Sprintf("%s:%s@tcp/%s?charset=utf8", User, Password, DbName)
