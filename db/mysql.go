package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

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
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

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
	print(os.Getenv("APP_ENV"))
	print(os.Getenv("ENV"))
	// local環境では127.0.0.1
	// host := getParamString("MYSQL_DB_HOST", "127.0.0.1")
	// 本番環境では127.0.0.1
	host := getParamString("MYSQL_DB_HOST", "127.0.0.1")
	port := getParamString("MYSQL_PORT", "3306")
	user := getParamString("MYSQL_USER", "root")
	// local環境では
	// pass := getParamString("MYSQL_PASSWORD", "")
	// 本番環境では
	pass := getParamString("MYSQL_PASSWORD", "password")
	dbname := getParamString("MYSQL_DB", "todoList")
	protocol := os.Getenv("MYSQL_PROTOCOL")

	if os.Getenv("ENV") == "local" {
		host = os.Getenv("MYSQL_DB_HOST")
		port = os.Getenv("MYSQL_PORT")
		user = os.Getenv("MYSQL_USER")
		pass = os.Getenv("MYSQL_PASSWORD")
		dbname = os.Getenv("MYSQL_DB")
	}

	if os.Getenv("APP_ENV") == "test" {
		dbname = strings.Join([]string{dbname, "_test"}, "")
	}
	// protocol := getParamString("MYSQL_PROTOCOL", "tcp")
	dbargs := getParamString("MYSQL_DBARGS", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	fmt.Println(user, pass, protocol, host, port, dbname, dbargs, "bbbb")
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}
