package leetcode

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

func TestSqlDrivers(t *testing.T) {
	for _, s := range sql.Drivers() {
		println(s)
	}
}

func TestSql(t *testing.T) {
	const (
		host     = "postgres.suqf.top"
		user     = "postgres"
		password = "postgres"
		dbname   = "test"
	)

	datasource := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
	)

	db, _ := sql.Open("postgres", datasource)
	defer db.Close()

	rows, _ := db.Query("select version()")

	for rows.Next() {
		var version string
		_ = rows.Scan(&version)
		println("version:", version)
	}

	tx, _ := db.Begin()
	stmt, _ := db.Prepare("insert into factor(factor_id, identity_id, trade_day, value) values ($1, $2, $3, $4)")

	_, _ = stmt.Exec(0, 0, 0, 0.0)

	_ = tx.Commit()
}
