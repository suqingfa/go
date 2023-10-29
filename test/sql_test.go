package test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"testing"
)

func TestSqlDrivers(t *testing.T) {
	for _, s := range sql.Drivers() {
		println(s)
	}
}

func TestSql(t *testing.T) {
	db, _ := sql.Open("sqlite3", "file::memory:?cache=shared")
	defer db.Close()

	rows, _ := db.Query("select sqlite_version()")

	for rows.Next() {
		var version string
		_ = rows.Scan(&version)
		log.Println("sql version:", version)
	}

	_, _ = db.Exec("create table user(id int primary key, name varchar(255), password varchar(255))")

	tx, _ := db.Begin()
	stmt, _ := db.Prepare("insert into user(id, name, password) values ($1, $2, $3)")

	_, _ = stmt.Exec(0, "user", "password")
	_, _ = stmt.Exec(1, "admin", "password")

	query, _ := db.Query("select id, name, password from user")
	for query.Next() {
		columns, _ := rows.Columns()
		log.Println("entity:", columns)
	}

	_ = tx.Commit()
}
