package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var DB *sql.DB

func main() {
	name, err := getQueryUserInfo(1)
	if err != nil {
		fmt.Printf("ID = %v has no data, err is %+v\n", 100, err)
		return
	}
	fmt.Println(name)
}

func init() {
	DB, err := sql.Open("mysql", "root:queen-1992@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("ping", err)
	}
}

func getQueryUserInfo(id int) (string, error) {
	var (
		name string
		err  error
	)
	err = DB.QueryRow("select * from user where id = 3").Scan(name)
	err = errors.Wrapf(err, "getQueryUserInfo")
	return name, err
}
