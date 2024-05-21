package test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqlDrivers(t *testing.T) {
	for _, s := range sql.Drivers() {
		t.Log(s)
	}
}

func TestSql(t *testing.T) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	assert.Nil(t, err)
	defer db.Close()

	rows, err := db.Query("select sqlite_version()")
	assert.Nil(t, err)

	for rows.Next() {
		var version string
		err = rows.Scan(&version)
		assert.Nil(t, err)
		t.Log("sql version:", version)
	}

	_, err = db.Exec("create table user(id int primary key, name varchar(255), password varchar(255))")
	assert.Nil(t, err)

	tx, err := db.Begin()
	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	assert.Nil(t, err)
	stmt, err := db.Prepare("insert into user(id, name, password) values ($1, $2, $3)")
	assert.Nil(t, err)

	_, err = stmt.Exec(0, "user", "password")
	assert.Nil(t, err)
	_, err = stmt.Exec(1, "admin", "password")
	assert.Nil(t, err)

	query, err := db.Query("select id, name, password from user")
	assert.Nil(t, err)

	columns, _ := query.Columns()
	t.Log("query", "columns", columns)

	for query.Next() {
		var id int
		var name string
		var password string
		err = query.Scan(&id, &name, &password)
		assert.Nil(t, err)
		t.Log("query result", "id", id, "name", name, "password", password)
	}
}
