package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:ShodiqB20@tcp(localhost:3306)/db_golang_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db

	// migrate -database "mysql://root:ShodiqB20@tcp(localhost:3306)/db_golang_migration" -path db/migrations up
	// migrate -database "mysql://root:ShodiqB20@tcp(localhost:3306)/db_golang_migration" -path db/migrations down

	// migrate -database "mysql://root:ShodiqB20@tcp(localhost:3306)/db_golang_migration" -path db/migrations up 2
	// migrate -database "mysql://root:ShodiqB20@tcp(localhost:3306)/db_golang_migration" -path db/migrations down 1

	// migrate -database "mysql://root:ShodiqB20@tcp(localhost:3306)/db_golang_migration" -path db/migrations version
	// migrate -database "mysql://root:ShodiqB20@tcp(localhost:3306)/db_golang_migration" -path db/migrations force 20240628070505
}
