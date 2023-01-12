package test

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func OpenDB(user, password, host, port, dbName string) *sql.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		user,
		password,
		host,
		port,
		dbName,
	)

	db, _ := sql.Open("mysql", dataSource)
	return db
}

func CloseDB(db *sql.DB) {
	_ = db.Close()
}

func TruncateTables(db *sql.DB, tables []string) {
	_, _ = db.Exec("SET FOREIGN_KEY_CHECKS=0;")

	for _, v := range tables {
		_, _ = db.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", v))
	}

	_, _ = db.Exec("SET FOREIGN_KEY_CHECKS=1;")
}

func LoadFixtures(db *sql.DB, queries []string) {
	for _, query := range queries {
		query = strings.TrimSuffix(query, ",\n")
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
}
