package request

import "fmt"

type InitDB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbName" v:"required"`
	UserName string `json:"userName" v:"required"`
	Password string `json:"password"`
	LogMod   bool  `json:"logMod"`
}

func (i *InitDB) SqlDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.UserName, i.Password, i.Host, i.Port)
}

func (i *InitDB) GetCreateTableSql() string {
	return fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", i.DBName)
}
