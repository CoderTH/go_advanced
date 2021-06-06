package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

func GetUserName(id int) (string, error) {
	db, err := sql.Open("mysql",
		"root:root@tcp(localhost:3306)/test")
	if err != nil {
		return "", errors.Wrap(err, "mysql connect err")
	}
	defer db.Close()
	var UserName string
	err = db.QueryRow("select name from users where id = ?", id).Scan(&UserName)
	if err != nil {
		return "", errors.Wrap(err, "query error")
	}
	return UserName, nil
}

func main() {
	name, err := GetUserName(2)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("没有该人")
			return
		}
		fmt.Printf("original error %T,%v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack error:\n%+v\n", err)
		return
	}
	fmt.Println("we got the name:" + name)
}
