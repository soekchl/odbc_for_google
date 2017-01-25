package main

import (
	"database/sql"
	"fmt"

	_ "code.google.com/p/odbc"
)

func main() {
	conn, err := sql.Open("odbc", "driver={sql server};server=192.168.1.1;port=1433;uid=sa;pwd=password;database=test")

	if err != nil {
		fmt.Println("Connecting Error")
		return
	}
	defer conn.Close()
	stmt, err := conn.Prepare("select top 5 id from users_role")
	if err != nil {
		fmt.Println("Query Error", err)
		return
	}
	defer stmt.Close()
	row, err := stmt.Query()
	if err != nil {
		fmt.Println("Query Error", err)
		return
	}
	defer row.Close()
	for row.Next() {
		var id int
		if err := row.Scan(&id); err == nil {
			fmt.Println(id)
		}
	}
	fmt.Printf("%s\n", "finish")
	return
}
