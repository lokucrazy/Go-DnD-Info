package services

import (
	"database/sql"
	"log"
)

type Service interface {
	Get(string) ([]byte, error)
}

func closeRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		log.Println(err.Error())
	}
}
