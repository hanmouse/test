package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	driverName := "mysql"

	user := "root"
	passwd := "root.123"
	protocol := "tcp"
	host := "localhost"
	port := 3306
	dbname := "testdb"

	dataSource := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", user, passwd, protocol, host, port, dbname)
	fmt.Printf("dataSource: %#v\n", dataSource)

	// 일반적으로는 이 시잠에 DB에 접속하지 않고, 이후 query 문 발생 시 실제로 접속한다.
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var name string
	const tableName = "test1"
	id := 1

	var query string

	query = fmt.Sprintf("SELECT name FROM %s WHERE id = ?", tableName)
	fmt.Printf("query: %#v\n", query)

	// QueryRow 메써드는 하나의 row만 select 할 때 사용한다.
	err = db.QueryRow(query, id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %#v\n", name)

	query = fmt.Sprintf("SELECT id, name FROM %s WHERE id >= ?", tableName)
	fmt.Printf("query: %#v\n", query)

	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %#v, name: %#v\n", id, name)
	}

	query = fmt.Sprintf("INSERT INTO %s (id, name) VALUES (?, ?)", tableName)
	fmt.Printf("query: %#v\n", query)

	newID := 3
	newName := "Anonymous"

	result, err := db.Exec(query, newID, newName)
	if err != nil {
		log.Fatal(err)
	}

	numRowsAffected, err := result.RowsAffected()
	if numRowsAffected == 1 {
		fmt.Printf("")
	}
}
