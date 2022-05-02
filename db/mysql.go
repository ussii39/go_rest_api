package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	// _ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() (*sql.DB, error) {
	connectionString := getConnectionString()
	db, err := sql.Open("mysql", connectionString)
	// db2, err2 := gorm.Open("mysql", connectionString)
	// fmt.Println(db2)
	// fmt.Println(err2)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	if env != "" {
		return env
	}
	return defaultValue
}

func getConnectionString() string {
	host := getParamString("MYSQL_DB_HOST", "db")
	port := getParamString("MYSQL_PORT", "3306")
	user := getParamString("MYSQL_USER", "root")
	pass := getParamString("MYSQL_PASSWORD", "")
	dbname := getParamString("MYSQL_DB", "todoList")
	if os.Getenv("APP_ENV") == "test" {
		dbname = strings.Join([]string{dbname, "_test"}, "")
	}
	protocol := getParamString("MYSQL_PROTOCOL", "tcp")
	dbargs := getParamString("MYSQL_DBARGS", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	fmt.Println(user, pass, protocol, host, port, dbname, dbargs)
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}
